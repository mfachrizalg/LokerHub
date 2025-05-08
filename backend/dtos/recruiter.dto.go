package dtos

import "github.com/google/uuid"

type RegisterRecruiterRequest struct {
	Name      string    `json:"name" validate:"required"`
	Position  string    `json:"position" validate:"required"`
	Handphone string    `json:"handphone" validate:"required,e164"`
	PhotoURL  string    `json:"photo_url" validate:"required"`
	CompanyID uuid.UUID `json:"company_id" validate:"required"`
}
