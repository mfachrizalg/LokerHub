package dtos

import "github.com/google/uuid"

type RegisterRecruiterRequest struct {
	Name      string    `json:"name" validate:"required"`
	Handphone string    `json:"handphone" validate:"required,e164"`
	CompanyID uuid.UUID `json:"company_id" validate:"required"`
}
