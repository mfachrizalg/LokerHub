package models

import (
	"github.com/google/uuid"
	"time"
)

type ProcessStatus string

const (
	Applied  ProcessStatus = "Applied"
	Reviewed ProcessStatus = "Reviewed"
	Rejected ProcessStatus = "Rejected"
	Hired    ProcessStatus = "Hired"
)

type Process struct {
	ID          uuid.UUID     `gorm:"type:uuid;default:gen_random_uuid();primaryKey"`
	JobID       uuid.UUID     `gorm:"type:uuid;not null"`
	CandidateID uuid.UUID     `gorm:"type:uuid;not null"`
	Status      ProcessStatus `gorm:"size:25;not null;default:'Applied'"`
	AppliedAt   time.Time     `gorm:"autoCreateTime"`
	UpdatedAt   time.Time     `gorm:"autoUpdateTime"`
	Detail      string        `gorm:"size:100"`

	Job       Job       `gorm:"foreignKey:JobID"`
	Candidate Candidate `gorm:"foreignKey:CandidateID"`
}
