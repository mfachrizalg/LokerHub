package dtos

type RegisterRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
	Role     string `json:"role" validate:"required,oneof=Candidate Recruiter"`
}
