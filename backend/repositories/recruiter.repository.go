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

func (r *RecruiterRepository) Create(recruiter *models.Recruiter) error {
	return r.DB.Create(&recruiter).Error
}

func (r *RecruiterRepository) Update(recruiter *models.Recruiter) error {
	return r.DB.Updates(&recruiter).Error
}

func (r *RecruiterRepository) GetRecruiterDetail(userID *uuid.UUID) (*models.Recruiter, error) {
	var recruiter *models.Recruiter
	err := r.DB.Where("user_id = ?", userID).Find(&recruiter).Error
	return recruiter, err
}

func (r *RecruiterRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}
