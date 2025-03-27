package services

import (
	"backend/dtos"
	"backend/models"
	"backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
)

type RecruiterService struct {
	repo *repositories.RecruiterRepository
}

func NewRecruiterService(repo *repositories.RecruiterRepository) *RecruiterService {
	return &RecruiterService{
		repo: repo,
	}
}

func (s *RecruiterService) RegisterRecruiter(req *dtos.RegisterRecruiterRequest, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	tx := s.repo.BeginTransaction()
	defer tx.Rollback()
	userId, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		log.Error("Error validating user id")
		return nil, errors.New("invalid user id format")
	}
	recruiter := models.Recruiter{
		CompanyID: req.CompanyID,
		UserID:    userId,
		Name:      req.Name,
		Handphone: req.Handphone,
	}

	if err := s.repo.Update(&recruiter); err != nil {
		tx.Rollback()
		log.Error("Error creating recruiter: ", err)
		return nil, errors.New("failed to process registration")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to process registration")
	}

	return &dtos.MessageResponse{
		Message: "Recruiter registered successfully",
	}, nil
}
