package models

import (
	"github.com/google/uuid"
)

type Company struct {
	ID        uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	Name      string    `gorm:"size:50;not null"`
	Location  string    `gorm:"size:255;not null"`
	Handphone string    `gorm:"size:15;unique;not null"`
	Logo      string    `gorm:"type:text;not null"`

	Recruiters []Recruiter `gorm:"foreignKey:CompanyID"`
	Jobs       []Job       `gorm:"foreignKey:CompanyID"`
}
