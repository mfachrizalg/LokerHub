package dtos

import (
	"backend/models"
	"github.com/google/uuid"
	"time"
)

// ProcessResponse represents the response for a single process record
type ProcessResponse struct {
	ID          uuid.UUID            `json:"id"`
	JobID       uuid.UUID            `json:"job_id"`
	CandidateID uuid.UUID            `json:"candidate_id"`
	Status      models.ProcessStatus `json:"status"`
	AppliedAt   time.Time            `json:"applied_at"`
	UpdatedAt   time.Time            `json:"updated_at"`
	Detail      string               `json:"detail"`
	// Consider adding candidate details if needed, e.g., CandidateName string `json:"candidate_name"`
}

// Add other process-related DTOs here if needed (e.g., UpdateProcessRequest)
