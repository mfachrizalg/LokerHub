package models

import (
	"github.com/google/uuid"
	"time"
)

type Job struct {
	ID             uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CompanyID      uuid.UUID `gorm:"type:uuid;not null"`
	RecruiterID    uuid.UUID `gorm:"type:uuid;not null"`
	Name           string    `gorm:"size:100;not null"`
	Type           string    `gorm:"size:100;not null"`
	Position       string    `gorm:"size:100;not null"`
	Salary         int       `gorm:"not null"`
	Field          string    `gorm:"size:100;not null"`
	Description    string    `gorm:"type:text;not null"`
	Responsibility string    `gorm:"type:text;not null"`
	Qualification  string    `gorm:"type:text;not null"`
	ClosedAt       time.Time `gorm:"not null"`
	CreatedAt      time.Time `gorm:"autoCreateTime"`
	UpdatedAt      time.Time `gorm:"autoUpdateTime"`

	Company   Company   `gorm:"foreignKey:CompanyID"`
	Recruiter Recruiter `gorm:"foreignKey:RecruiterID"`
	Processes []Process `gorm:"foreignKey:JobID"`
}
