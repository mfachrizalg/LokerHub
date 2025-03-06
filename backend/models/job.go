package models

import (
	"github.com/google/uuid"
)

type Job struct {
	ID            uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CompanyID     uuid.UUID `gorm:"type:uuid;not null"`
	RecruiterID   uuid.UUID `gorm:"type:uuid;not null"`
	Name          string    `gorm:"size:100;not null"`
	Criteria      string    `gorm:"type:text;not null"`
	Qualification string    `gorm:"type:text;not null"`
	Status        string    `gorm:"size:25;not null"`

	Company   Company   `gorm:"foreignKey:CompanyID"`
	Recruiter Recruiter `gorm:"foreignKey:RecruiterID"`
	Processes []Process `gorm:"foreignKey:JobID"`
}
