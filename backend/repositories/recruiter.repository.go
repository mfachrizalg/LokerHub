package repositories

import (
	"backend/config"
	"backend/models"
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
	return r.DB.Save(&recruiter).Error
}

func (r *RecruiterRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}
