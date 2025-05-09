package services

import (
	"backend/dtos"
	"backend/helpers"
	"backend/models"
	"backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/google/uuid"
	"mime/multipart"
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

func (s *CandidateService) GetCandidateDetail(userID uuid.UUID) (*models.Candidate, error) {
	tx := s.repo.BeginTransaction()
	defer tx.Rollback()

	candidate, err := s.repo.GetByUserID(&userID)
	if err != nil {
		log.Error("Error getting candidate: ", err)
		return nil, errors.New("failed to get candidate details")
	}

	if err := tx.Commit().Error; err != nil {
		log.Error("Error committing transaction: ", err)
		return nil, errors.New("failed to process request")
	}

	return candidate, nil
}

func (s *CandidateService) UploadCandidatePhoto(file *multipart.FileHeader) (string, error) {
	return helpers.UploadPhoto(file, "1uGimZhfrohl_UefdkAEYH4j49Jmhn_hX")
}

func (s *CandidateService) UploadCandidateCV(file *multipart.FileHeader) (string, error) {
	return helpers.UploadPhoto(file, "1uGimZhfrohl_UefdkAEYH4j49Jmhn_hX")
}
