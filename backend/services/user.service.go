package services

import (
	"backend/dtos"
	"backend/models"
	"backend/repositories"
	"errors"

	"github.com/gofiber/fiber/v2/log"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type UserService struct {
	userRepo      *repositories.UserRepository
	candidateRepo *repositories.CandidateRepository
	recruiterRepo *repositories.RecruiterRepository
}

func NewUserService(
	userRepo *repositories.UserRepository,
	candidateRepo *repositories.CandidateRepository,
	recruiterRepo *repositories.RecruiterRepository,
) *UserService {
	return &UserService{
		userRepo:      userRepo,
		candidateRepo: candidateRepo,
		recruiterRepo: recruiterRepo,
	}
}

func (s *UserService) Register(req *dtos.RegisterRequest) (*dtos.MessageResponse, error) {
	// Check if email already exists
	existingUser, err := s.userRepo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Database error: ", err)
		return nil, errors.New("failed to process registration")
	}

	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		log.Error("Error hashing password: ", err)
		return nil, errors.New("failed to process registration")
	}

	// Create user transaction
	tx := s.userRepo.BeginTransaction()
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	// Create new user
	user := models.User{
		Email:    req.Email,
		Password: string(hashedPassword),
		Role:     req.Role,
	}

	if err := s.userRepo.Create(&user); err != nil {
		tx.Rollback()
		log.Error("Error creating user: ", err)
		return nil, errors.New("failed to create user")
	}

	// Based on role, create either Candidate or Recruiter
	if req.Role == "Candidate" {
		candidate := models.Candidate{
			UserID: user.ID,
		}
		if err := s.candidateRepo.Create(&candidate).Error; err != nil {
			tx.Rollback()
			log.Error("Error creating candidate: ", err)
			return nil, errors.New("failed to create candidate profile")
		}
	} else if req.Role == "Recruiter" {
		recruiter := models.Recruiter{
			UserID: user.ID,
		}
		if err := s.recruiterRepo.Create(&recruiter).Error; err != nil {
			tx.Rollback()
			log.Error("Error creating recruiter: ", err)
			return nil, errors.New("failed to create recruiter profile")
		}
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("registration failed")
	}

	return &dtos.MessageResponse{
		Message: "User registered successfully!",
	}, nil
}
