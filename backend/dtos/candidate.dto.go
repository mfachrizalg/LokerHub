package dtos

type RegisterCandidateRequest struct {
	Name        string `json:"name" validate:"required"`
	Description string `json:"description" validate:"required"`
	Handphone   string `json:"handphone" validate:"required,e164"`
	Photo       string `json:"photo" validate:"required"`
	Education   string `json:"education" validate:"required"`
	Field       string `json:"field" validate:"required"`
	Location    string `json:"location" validate:"required"`
	CV          string `json:"cv" validate:"required"`
}
