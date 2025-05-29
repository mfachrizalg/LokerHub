package repositories

import (
	"backend/config"
	"backend/models"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type RecruiterRepository struct {
	DB *gorm.DB
}

func NewRecruiterRepository() *RecruiterRepository {
	return &RecruiterRepository{
		DB: config.DB,
	}
}

func (r *RecruiterRepository) CreateWithTx(tx *gorm.DB, recruiter *models.Recruiter) error {
	return tx.Create(recruiter).Error
}

func (r *RecruiterRepository) UpdateWithTx(tx *gorm.DB, recruiter *models.Recruiter) error {
	return tx.Save(recruiter).Error
}

func (r *RecruiterRepository) FindByUserId(id uuid.UUID) (uuid.UUID, error) {
	var recruiter *models.Recruiter

	err := r.DB.Where("user_id = ?", id).First(&recruiter).Error
	if err != nil {
		return uuid.Nil, err
	}
	return recruiter.ID, nil
}

func (r *RecruiterRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}
