package repositories

import (
	"backend/config"
	"backend/models"
	"gorm.io/gorm"
)

type CandidateRepository struct {
	DB *gorm.DB
}

func NewCandidateRepository() *CandidateRepository {
	return &CandidateRepository{
		DB: config.DB,
	}
}

func (r *CandidateRepository) Create(candidate *models.Candidate) error {
	return r.DB.Create(&candidate).Error
}

func (r *CandidateRepository) Update(candidate *models.Candidate) error {
	return r.DB.Save(&candidate).Error
}

func (r *CandidateRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}
