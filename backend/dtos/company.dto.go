package dtos

import "github.com/google/uuid"

type RegisterCompanyRequest struct {
	Name     string `json:"name" validate:"required"`
	Location string `json:"location" validate:"required"`
	Industry string `json:"industry" validate:"required"`
	Logo     string `json:"logo"`
}

type GetAllCompanyResponse []struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}
