package models

import (
	"github.com/google/uuid"
)

type Candidate struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"size:50;not null"`
	Description string    `gorm:"type:text; not null"`
	Handphone   string    `gorm:"size:15;unique;not null"`
	Photo       string    `gorm:"size:255"`
	Education   string    `gorm:"size:50;not null"`
	Field       string    `gorm:"size:50;not null"`
	Location    string    `gorm:"size:50;not null"`
	CV          string    `gorm:"size:255"`

	User      User      `gorm:"foreignKey:UserID"`
	Processes []Process `gorm:"foreignKey:CandidateID"`
}
