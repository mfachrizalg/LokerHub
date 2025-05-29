package services

import (
	"backend/dtos"
	"backend/models"
	"backend/repositories"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"gorm.io/gorm"
	"os"

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

// GetAllJobs retrieves all jobs
func (s *JobService) GetAllJobs() (*[]dtos.JobResponse, error) {
	// Call repository method without pagination parameters
	jobs, err := s.jobRepo.FindAll()
	if err != nil {
		log.Error("Error retrieving jobs: ", err)
		return nil, errors.New("failed to retrieve jobs")
	}

	// Convert model to DTO
	jobResponses := make([]dtos.JobResponse, len(jobs))
	for i, job := range jobs {
		companyName, location, logo, err := s.jobRepo.GetCompanyNameAndLocationByID(job.CompanyID)
		if err != nil {
			log.Error("Error retrieving company name: ", err)
			return nil, errors.New("failed to retrieve company name")
		}
		// Convert each job to its corresponding DTO
		jobResponses[i] = dtos.JobResponse{
			ID:          job.ID,
			CompanyID:   job.CompanyID,
			Name:        job.Name,
			CompanyName: companyName,
			Location:    location,
			CreatedAt:   job.CreatedAt,
			CompanyLogo: logo,
		}
	}

	return &jobResponses, nil
}

// GetJobByID retrieves a job by its ID
func (s *JobService) GetJobByID(jobID uuid.UUID) (*dtos.JobDetailResponse, error) {
	job, err := s.jobRepo.FindById(jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Job with ID %s not found", jobID)
			return nil, errors.New("job not found")
		}
		log.Error("Error retrieving job: ", err)
		return nil, errors.New("failed to retrieve job")
	}

	companyName, location, logo, err := s.jobRepo.GetCompanyNameAndLocationByID(job.CompanyID)
	if err != nil {
		log.Error("Error retrieving company name: ", err)
		return nil, errors.New("failed to retrieve company name")
	}

	jobResponse := dtos.JobDetailResponse{
		ID:             job.ID,
		Name:           job.Name,
		Type:           job.Type,
		Position:       job.Position,
		Salary:         job.Salary,
		Field:          job.Field,
		Description:    job.Description,
		Responsibility: job.Responsibility,
		Qualification:  job.Qualification,
		CompanyLogo:    logo,
		CompanyName:    companyName,
		Location:       location,
	}

	return &jobResponse, nil
}

// GetJobsByRecruiterID retrieves all jobs posted by a specific recruiter
func (s *JobService) GetJobsByRecruiterID(ctx *fiber.Ctx) (*[]dtos.JobResponse, error) {
	cookie := ctx.Cookies("LokerHubCookie")
	if cookie == "" {
		log.Error("No authentication cookie found")
		return nil, errors.New("authentication required")
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		// Return the secret key used for signing
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		log.Error("Error parsing token: ", err)
		return nil, errors.New("invalid authentication token")
	}

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Error("Invalid token claims")
		return nil, errors.New("invalid authentication token")
	}

	userIdStr, ok := claims["id"].(string)
	if !ok {
		log.Error("User ID not found in token claims")
		return nil, errors.New("invalid authentication token")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Error("Error parsing user ID: ", err)
		return nil, errors.New("invalid user information")
	}

	recruiterID, err := s.jobRepo.GetRecruiterIDByUserID(userId)
	if err != nil {
		log.Error("Error validating recruiter ID")
		return nil, errors.New("invalid recruiter ID format")
	}

	jobs, err := s.jobRepo.FindAllByRecruiterID(recruiterID)
	if err != nil {
		log.Error("Error retrieving jobs by recruiter ID: ", err)
		return nil, errors.New("failed to retrieve jobs")
	}

	// Convert model to DTO
	jobResponses := make([]dtos.JobResponse, len(jobs))
	for i, job := range jobs {
		companyName, location, logo, err := s.jobRepo.GetCompanyNameAndLocationByID(job.CompanyID)
		if err != nil {
			log.Error("Error retrieving company name: ", err)
			return nil, errors.New("failed to retrieve company name")
		}
		// Convert each job to its corresponding DTO
		jobResponses[i] = dtos.JobResponse{
			ID:          job.ID,
			CompanyID:   job.CompanyID,
			Name:        job.Name,
			CompanyName: companyName,
			Location:    location,
			CreatedAt:   job.CreatedAt,
			CompanyLogo: logo,
		}
	}
	return &jobResponses, nil
}

// CreateJob creates a new job posting
func (s *JobService) CreateJob(req *dtos.CreateJobRequest, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	tx := s.jobRepo.BeginTransaction()
	defer tx.Rollback()

	cookie := ctx.Cookies("LokerHubCookie")
	if cookie == "" {
		log.Error("No authentication cookie found")
		return nil, errors.New("authentication required")
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		// Return the secret key used for signing
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		log.Error("Error parsing token: ", err)
		return nil, errors.New("invalid authentication token")
	}

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Error("Invalid token claims")
		return nil, errors.New("invalid authentication token")
	}

	userIdStr, ok := claims["id"].(string)
	if !ok {
		log.Error("User ID not found in token claims")
		return nil, errors.New("invalid authentication token")
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Error("Error parsing user ID: ", err)
		return nil, errors.New("invalid user information")
	}

	recruiterID, err := s.jobRepo.GetRecruiterIDByUserID(userId)
	if err != nil {
		log.Error("Error validating recruiter ID")
		return nil, errors.New("invalid recruiter ID format")
	}

	companyID, err := s.jobRepo.GetCompanyIDByRecruiterID(recruiterID)
	if err != nil {
		log.Error("Error retrieving company ID: ", err)
		return nil, errors.New("failed to retrieve company ID")
	}

	job := models.Job{
		RecruiterID:    recruiterID,
		CompanyID:      companyID,
		Name:           req.Name,
		Type:           req.Type,
		Position:       req.Position,
		Salary:         req.Salary,
		Field:          req.Field,
		Description:    req.Description,
		Responsibility: req.Responsibility,
		ClosedAt:       req.ClosedAt,
		Qualification:  req.Qualification,
	}

	if err := s.jobRepo.Create(&job); err != nil {
		log.Error("Error creating job: ", err)
		return nil, errors.New("failed to create job posting")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to create job posting")
	}

	return &dtos.MessageResponse{
		Message: "Successfully created job",
	}, nil
}

// ApplyJob allows a candidate to apply for a job
func (s *JobService) ApplyJob(jobID uuid.UUID, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	cookie := ctx.Cookies("LokerHubCookie")
	if cookie == "" {
		log.Error("No authentication cookie found")
		return nil, errors.New("authentication required")
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		// Return the secret key used for signing
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		log.Error("Error parsing token: ", err)
		return nil, errors.New("invalid authentication token")
	}

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Error("Invalid token claims")
		return nil, errors.New("invalid authentication token")
	}

	userIdStr, ok := claims["id"].(string)
	if !ok {
		log.Error("User ID not found in token claims")
		return nil, errors.New("invalid authentication token")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Error("Error parsing user ID: ", err)
		return nil, errors.New("invalid user information")
	}

	candidateID, err := s.jobRepo.GetCandidateIDByUserID(userId)
	if err != nil {
		log.Error("Error retrieving candidate ID: ", err)
		return nil, errors.New("invalid candidate ID format")
	}

	// Check if job exists (optional but good practice)
	_, err = s.jobRepo.FindById(jobID)
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

// UpdateJob updates an existing job posting
func (s *JobService) UpdateJob(jobID uuid.UUID, req *dtos.UpdateJobRequest, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	tx := s.jobRepo.BeginTransaction() // Use jobRepo's transaction
	defer tx.Rollback()

	cookie := ctx.Cookies("LokerHubCookie")
	if cookie == "" {
		log.Error("No authentication cookie found")
		return nil, errors.New("authentication required")
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		// Return the secret key used for signing
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		log.Error("Error parsing token: ", err)
		return nil, errors.New("invalid authentication token")
	}

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Error("Invalid token claims")
		return nil, errors.New("invalid authentication token")
	}

	userIdStr, ok := claims["id"].(string)
	if !ok {
		log.Error("User ID not found in token claims")
		return nil, errors.New("invalid authentication token")
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Error("Error parsing user ID: ", err)
		return nil, errors.New("invalid user information")
	}

	recruiterID, err := s.jobRepo.GetRecruiterIDByUserID(userId)
	if err != nil {
		log.Error("Error retrieving recruiter ID: ", err)
		return nil, errors.New("invalid recruiter ID format")
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

	// Update job fields
	// Update job fields that are present in the request
	if req.Name != nil {
		job.Name = *req.Name
	}
	if req.Type != nil {
		job.Type = *req.Type
	}
	if req.Position != nil {
		job.Position = *req.Position
	}
	if req.Salary != nil {
		job.Salary = *req.Salary
	}
	if req.Field != nil {
		job.Field = *req.Field
	}
	if req.Description != nil {
		job.Description = *req.Description
	}
	if req.Responsibility != nil {
		job.Responsibility = *req.Responsibility
	}
	if req.ClosedAt != nil {
		job.ClosedAt = *req.ClosedAt
	}
	if req.Qualification != nil {
		job.Qualification = *req.Qualification
	}
	if err := s.jobRepo.Update(tx, job); err != nil {
		log.Error("Error updating job: ", err)
		return nil, errors.New("failed to update job")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to update job")
	}

	return &dtos.MessageResponse{
		Message: "Successfully updated job",
	}, nil

}
