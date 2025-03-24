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

type CandidateService struct {
	repo *repositories.CandidateRepository
}

func NewCandidateService(repo *repositories.CandidateRepository) *CandidateService {
	return &CandidateService{
		repo: repo,
	}
}

func (s *CandidateService) RegisterCandidate(req *dtos.RegisterCandidateRequest, ctx *fiber.Ctx) (*dtos.MessageResponse, error) {
	tx := s.repo.BeginTransaction()
	defer tx.Rollback()

	userId, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		log.Error("Error validating user id")
		return nil, errors.New("invalid user id format")
	}
	candidate := models.Candidate{
		UserID:      userId,
		Name:        req.Name,
		Description: req.Description,
		Handphone:   req.Handphone,
		Photo:       req.Photo,
		Education:   req.Education,
		Field:       req.Field,
		Location:    req.Location,
		CV:          req.CV,
	}

	if err := s.repo.Update(&candidate); err != nil {
		tx.Rollback()
		log.Error("Error creating candidate: ", err)
		return nil, errors.New("failed to process registration")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to process registration")
	}

	return &dtos.MessageResponse{
		Message: "Candidate registered successfully",
	}, nil
}
