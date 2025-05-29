package dtos

import (
	"github.com/google/uuid"
	"time"
)

type JobResponse struct {
	ID          uuid.UUID `json:"id"`
	CompanyID   uuid.UUID `json:"company_id"`
	CompanyLogo string    `json:"company_logo"`
	Name        string    `json:"name"`
	CompanyName string    `json:"company_name"`
	Location    string    `json:"location"`
	CreatedAt   time.Time `json:"created_at"`
}

type JobDetailResponse struct {
	ID             uuid.UUID `json:"id"`
	Name           string    `json:"name"`
	Type           string    `json:"type"`
	Position       string    `json:"position"`
	Salary         int       `json:"salary"`
	Field          string    `json:"field"`
	Description    string    `json:"description"`
	Responsibility string    `json:"responsibility"`
	Qualification  string    `json:"qualification"`
	Location       string    `json:"location"`
	CompanyLogo    string    `json:"company_logo"`
	CompanyName    string    `json:"company_name"`
}

// CreateJobRequest represents the request body for creating a job
type CreateJobRequest struct {
	Name           string    `json:"name" validate:"required"`
	Type           string    `json:"type" validate:"required"`
	Position       string    `json:"position" validate:"required"`
	Salary         int       `json:"salary" validate:"required"`
	Field          string    `json:"field" validate:"required"`
	Description    string    `json:"description" validate:"required"`
	Responsibility string    `json:"responsibility" validate:"required"`
	ClosedAt       time.Time `json:"closed_at" validate:"required"`
	Qualification  string    `json:"qualification" validate:"required"`
}

// UpdateJobRequest represents the request body for updating a job
type UpdateJobRequest struct {
	Name           *string    `json:"name,omitempty"`
	Type           *string    `json:"type,omitempty"`
	Position       *string    `json:"position,omitempty"`
	Salary         *int       `json:"salary,omitempty"`
	Field          *string    `json:"field,omitempty"`
	Description    *string    `json:"description,omitempty"`
	Responsibility *string    `json:"responsibility,omitempty"`
	ClosedAt       *time.Time `json:"closed_at,omitempty"`
	Qualification  *string    `json:"qualification,omitempty"`
}
