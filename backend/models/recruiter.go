package models

import (
	"github.com/google/uuid"
)

type Recruiter struct {
	ID        uuid.UUID  `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	CompanyID *uuid.UUID `gorm:"type:uuid"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null"`
	Name      string     `gorm:"size:50"`
	Handphone string     `gorm:"size:15;unique"`
	Position  string     `gorm:"size:50"`
	PhotoURL  string     `gorm:"size:255"`

	Company Company `gorm:"foreignKey:CompanyID"`
	User    User    `gorm:"foreignKey:UserID"`
	Jobs    []Job   `gorm:"foreignKey:RecruiterID"`
}
