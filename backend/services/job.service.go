package services

import (
	"backend/dtos"
	"backend/models"
	"backend/repositories"
	"errors"
	"fmt"
	"gorm.io/gorm"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type JobService struct {
	jobRepo     *repositories.JobRepository
	processRepo *repositories.ProcessRepository
}

func NewJobService(jobRepo *repositories.JobRepository, processRepo *repositories.ProcessRepository) *JobService {
	return &JobService{
		jobRepo:     jobRepo,
		processRepo: processRepo,
	}
}

// GetAllJobs retrieves all jobs (without pagination)
func (s *JobService) GetAllJobs() ([]models.Job, error) {
	// Call repository method without pagination parameters
	jobs, err := s.jobRepo.FindAll()
	if err != nil {
		log.Error("Error retrieving jobs: ", err)
		return nil, errors.New("failed to retrieve jobs")
	}

	// Return only the jobs slice and error
	return jobs, nil
}

// CreateJob creates a new job posting
func (s *JobService) CreateJob(req *dtos.CreateJobRequest, ctx *fiber.Ctx) (*dtos.JobResponse, error) {
	tx := s.jobRepo.BeginTransaction()
	defer tx.Rollback()

	recruiterID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		log.Error("Error validating recruiter ID")
		return nil, errors.New("invalid recruiter ID format")
	}

	job := models.Job{
		CompanyID:     req.CompanyID,
		RecruiterID:   recruiterID,
		Name:          req.Name,
		Criteria:      req.Criteria,
		Qualification: req.Qualification,
		Status:        "Active", // Default status for new jobs
	}

	if err := s.jobRepo.Create(&job); err != nil {
		log.Error("Error creating job: ", err)
		return nil, errors.New("failed to create job posting")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to create job posting")
	}

	return &dtos.JobResponse{
		ID:            job.ID,
		CompanyID:     job.CompanyID,
		RecruiterID:   job.RecruiterID,
		Name:          job.Name,
		Criteria:      job.Criteria,
		Qualification: job.Qualification,
		Status:        job.Status,
	}, nil
}

// ApplyJob allows a candidate to apply for a job
func (s *JobService) ApplyJob(jobID uuid.UUID, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	candidateID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		log.Error("Error validating candidate ID from context")
		return nil, errors.New("invalid user ID format")
	}

	// Check if job exists (optional but good practice)
	_, err := s.jobRepo.FindById(jobID)
	if err != nil {
		log.Error("Error finding job: ", err)
		return nil, errors.New("job not found")
	}

	// Check if candidate already applied
	exists, err := s.processRepo.CheckIfExists(jobID, candidateID)
	if err != nil {
		log.Error("Error checking existing application: ", err)
		return nil, errors.New("failed to process application")
	}
	if exists {
		return nil, errors.New("already applied for this job")
	}

	// Use processRepo's transaction
	tx := s.processRepo.BeginTransaction()
	defer tx.Rollback()

	process := models.Process{
		JobID:       jobID,
		CandidateID: candidateID,
		Status:      models.Applied, // Default status
	}

	// Use the transaction for the Create operation
	if err := s.processRepo.Create(tx, &process); err != nil {
		log.Error("Error creating process record: ", err)
		// Rollback is handled by defer
		return nil, errors.New("failed to apply for job")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		// Rollback is handled by defer
		return nil, errors.New("failed to apply for job")
	}

	return &dtos.MessageResponse{
		Message: fmt.Sprintf("Successfully applied for job %s", jobID),
	}, nil
}

func (s *JobService) UpdateJob(jobID uuid.UUID, req *dtos.UpdateJobRequest, ctx *fiber.Ctx) (*dtos.JobResponse, error) {
	tx := s.jobRepo.BeginTransaction() // Use jobRepo's transaction
	defer tx.Rollback()

	recruiterID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		log.Error("Error validating recruiter ID from context")
		return nil, errors.New("invalid user ID format")
	}

	// Find the existing job
	job, err := s.jobRepo.FindById(jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Job with ID %s not found for update", jobID)
			return nil, errors.New("job not found")
		}
		log.Error("Error finding job for update: ", err)
		return nil, errors.New("failed to retrieve job for update")
	}

	// Authorization check: Ensure the recruiter owns the job
	if job.RecruiterID != recruiterID {
		log.Warnf("Recruiter %s attempted to update job %s owned by %s", recruiterID, jobID, job.RecruiterID)
		return nil, errors.New("unauthorized to update this job")
	}

	// Update fields only if they are provided in the request
	updated := false
	if req.Name != nil && *req.Name != "" {
		job.Name = *req.Name
		updated = true
	}
	if req.Criteria != nil && *req.Criteria != "" {
		job.Criteria = *req.Criteria
		updated = true
	}
	if req.Qualification != nil && *req.Qualification != "" {
		job.Qualification = *req.Qualification
		updated = true
	}
	if req.Status != nil && *req.Status != "" {
		// Add validation for allowed statuses if needed
		job.Status = *req.Status
		updated = true
	}

	if !updated {
		// Optional: return an error or message if no fields were updated
		log.Info("No fields provided for update for job ID: ", jobID)
		// Return current job state without hitting DB again
		return &dtos.JobResponse{
			ID:            job.ID,
			CompanyID:     job.CompanyID,
			RecruiterID:   job.RecruiterID,
			Name:          job.Name,
			Criteria:      job.Criteria,
			Qualification: job.Qualification,
			Status:        job.Status,
		}, nil
		// Or return errors.New("no update data provided")
	}

	// Pass the transaction to Update
	if err := s.jobRepo.Update(tx, job); err != nil {
		log.Error("Error updating job: ", err)
		// Rollback is handled by defer
		return nil, errors.New("failed to update job posting")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction for job update: ", err)
		// Rollback is handled by defer
		return nil, errors.New("failed to update job posting")
	}

	return &dtos.JobResponse{
		ID:            job.ID,
		CompanyID:     job.CompanyID,
		RecruiterID:   job.RecruiterID,
		Name:          job.Name,
		Criteria:      job.Criteria,
		Qualification: job.Qualification,
		Status:        job.Status,
	}, nil
}
