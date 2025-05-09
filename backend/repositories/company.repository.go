package repositories

import (
	"backend/config"
	"backend/models"
	"github.com/google/uuid"
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

func (r *CompanyRepository) FindByID(companyID *uuid.UUID) (*models.Company, error) {
	var company *models.Company
	err := r.DB.Where("id = ?", companyID).Find(&company).Error
	return company, err
}
