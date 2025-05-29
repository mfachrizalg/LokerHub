package repositories

import (
	"backend/config"
	"backend/models"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{
		DB: config.DB,
	}
}

func (r *UserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	err := r.DB.Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateWithTx(tx *gorm.DB, user *models.User) error {
	return tx.Create(user).Error
}

func (r *UserRepository) BeginTransaction() *gorm.DB {
	return r.DB.Begin()
}
