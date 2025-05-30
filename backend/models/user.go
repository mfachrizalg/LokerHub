package models

import (
	"github.com/google/uuid"
)

type User struct {
	ID       uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Email    string    `gorm:"size:100;unique;not null"`
	Password string    `gorm:"size:255;not null"`
	Role     string    `gorm:"type:roles;not null"`

	Candidates []Candidate `gorm:"foreignKey:UserID"`
	Recruiters []Recruiter `gorm:"foreignKey:UserID"`
}
