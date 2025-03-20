package dtos

type RegisterRecruiterRequest struct {
	Name      string `json:"name" validate:"required"`
	Handphone string `json:"handphone" validate:"required,e164"`
}
