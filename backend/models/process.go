package models

import (
	"github.com/google/uuid"
)

type Process struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	JobID       uuid.UUID `gorm:"type:uuid;not null"`
	RecruiterID uuid.UUID `gorm:"type:uuid;not null"`
	Stage       string    `gorm:"size:50;not null"`
	Detail      string    `gorm:"size:100;not null"`

	Job       Job       `gorm:"foreignKey:JobID"`
	Recruiter Recruiter `gorm:"foreignKey:RecruiterID"`
}
