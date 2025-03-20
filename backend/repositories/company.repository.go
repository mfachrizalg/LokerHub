package repositories

import (
	"backend/config"
	"backend/models"
	"gorm.io/gorm"
)

type CompanyRepository struct {
	DB *gorm.DB
}

func NewCompanyRepository() *CompanyRepository {
	return &CompanyRepository{
		DB: config.DB,
	}
}

func (r *CompanyRepository) Create(company *models.Company) error {
	return r.DB.Create(&company).Error
}

func (r *CompanyRepository) FindAll(company *[]models.Company) error {
	return r.DB.Select("id, name").Find(&company).Error
}

func (r *CompanyRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}
