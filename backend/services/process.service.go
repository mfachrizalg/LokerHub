package services

import (
	"backend/dtos"
	"backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProcessService struct {
	processRepo *repositories.ProcessRepository
	jobRepo     *repositories.JobRepository // Inject JobRepository for authorization
}

// NewProcessService creates a new ProcessService
func NewProcessService(processRepo *repositories.ProcessRepository, jobRepo *repositories.JobRepository) *ProcessService {
	return &ProcessService{
		processRepo: processRepo,
		jobRepo:     jobRepo,
	}
}

// GetAllProcessesByJob retrieves all processes for a specific job, ensuring authorization
func (s *ProcessService) GetAllProcessesByJob(jobID uuid.UUID, ctx *fiber.Ctx) ([]dtos.ProcessResponse, error) {
	recruiterID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		log.Error("Error validating recruiter ID from context")
		return nil, errors.New("invalid user ID format")
	}

	// 1. Verify the job exists and the recruiter owns it
	job, err := s.jobRepo.FindById(jobID)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			log.Warnf("Attempt to access processes for non-existent job ID: %s", jobID)
			return nil, errors.New("job not found")
		}
		log.Errorf("Error finding job %s for process retrieval: %v", jobID, err)
		return nil, errors.New("failed to retrieve job details")
	}

	if job.RecruiterID != recruiterID {
		log.Warnf("Recruiter %s attempted to access processes for job %s owned by %s", recruiterID, jobID, job.RecruiterID)
		return nil, errors.New("unauthorized to view processes for this job")
	}

	// 2. Retrieve processes from the repository
	processes, err := s.processRepo.FindAllByJobID(jobID)
	if err != nil {
		log.Errorf("Error retrieving processes for job %s: %v", jobID, err)
		return nil, errors.New("failed to retrieve processes")
	}

	// 3. Map models to DTOs
	processResponses := make([]dtos.ProcessResponse, len(processes))
	for i, p := range processes {
		processResponses[i] = dtos.ProcessResponse{
			ID:          p.ID,
			JobID:       p.JobID,
			CandidateID: p.CandidateID,
			Status:      p.Status,
			AppliedAt:   p.AppliedAt,
			UpdatedAt:   p.UpdatedAt,
			Detail:      p.Detail,
			// Map candidate details here if preloaded/fetched
		}
	}

	return processResponses, nil
}

// Add other process service methods here (e.g., UpdateProcessStatus)
