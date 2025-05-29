package repositories

import (
	"backend/config"
	"backend/models"
	"github.com/google/uuid"
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

func (r *CandidateRepository) CreateWithTx(tx *gorm.DB, candidate *models.Candidate) error {
	return tx.Create(candidate).Error
}

func (r *CandidateRepository) UpdateWithTx(tx *gorm.DB, candidate *models.Candidate) error {
	return tx.Save(candidate).Error
}

func (r *CandidateRepository) FindByUserId(userId uuid.UUID) (uuid.UUID, error) {
	var candidate models.Candidate

	err := r.DB.Where("user_id = ?", userId).First(&candidate).Error
	if err != nil {
		return uuid.Nil, err
	}
	return candidate.ID, nil
}

func (r *CandidateRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}
