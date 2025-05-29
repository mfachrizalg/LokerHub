package models

import (
	"github.com/google/uuid"
)

type Candidate struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"size:50"`
	Description string    `gorm:"type:text"`
	Handphone   string    `gorm:"size:15;unique"`
	Photo       string    `gorm:"size:255"`
	Education   string    `gorm:"size:50"`
	Field       string    `gorm:"size:50"`
	Location    string    `gorm:"size:50"`
	CV          string    `gorm:"size:255"`

	User      User      `gorm:"foreignKey:UserID"`
	Processes []Process `gorm:"foreignKey:CandidateID"`
}
