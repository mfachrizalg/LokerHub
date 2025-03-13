package models

import (
	"github.com/google/uuid"
)

type Role string

const (
	CandidateRole Role = "Candidate"
	RecruiterRole Role = "Recruiter"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email    string    `gorm:"size:100;unique;not null"`
	Password string    `gorm:"size:255;not null"`
	Role     Role      `gorm:"type:roles;not null"`

	Candidates []Candidate `gorm:"foreignKey:UserID"`
	Recruiters []Recruiter `gorm:"foreignKey:UserID"`
}
