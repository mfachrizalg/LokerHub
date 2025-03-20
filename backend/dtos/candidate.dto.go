package dtos

type RegisterCandidateRequest struct {
	Name       string `json:"name" validate:"required"`
	Education  string `json:"education" validate:"required"`
	Experience string `json:"experience" validate:"required"`
	Handphone  string `json:"handphone" validate:"required,e164"`
	Field      string `json:"field" validate:"required"`
	Location   string `json:"location" validate:"required"`
}
