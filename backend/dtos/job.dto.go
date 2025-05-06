package dtos

import "github.com/google/uuid"

// CreateJobRequest represents the request body for creating a job
type CreateJobRequest struct {
	CompanyID     uuid.UUID `json:"company_id" validate:"required"`
	Name          string    `json:"name" validate:"required"`
	Criteria      string    `json:"criteria" validate:"required"`
	Qualification string    `json:"qualification" validate:"required"`
}

// JobResponse represents the response body for job operations
type JobResponse struct {
	ID            uuid.UUID `json:"id"`
	CompanyID     uuid.UUID `json:"company_id"`
	RecruiterID   uuid.UUID `json:"recruiter_id"`
	Name          string    `json:"name"`
	Criteria      string    `json:"criteria"`
	Qualification string    `json:"qualification"`
	Status        string    `json:"status"`
}

type UpdateJobRequest struct {
	Name          *string `json:"name,omitempty"` // Use pointers to distinguish between empty and not provided
	Criteria      *string `json:"criteria,omitempty"`
	Qualification *string `json:"qualification,omitempty"`
	Status        *string `json:"status,omitempty"` // e.g., "Active", "Inactive"
}
