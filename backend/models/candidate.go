package models

import (
	"github.com/google/uuid"
)

type Candidate struct {
	ID          uuid.UUID `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	Name        string    `gorm:"size:50;not null"`
	Education   string    `gorm:"size:50;not null"`
	Handphone   string    `gorm:"size:15;unique;not null"`
	Domicile    string    `gorm:"size:50"`
	SocialMedia string    `gorm:"type:text"`

	User User `gorm:"foreignKey:UserID"`
}
