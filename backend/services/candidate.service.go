package services

import (
	"backend/dtos"
	"backend/helpers"
	"backend/models"
	"backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"mime/multipart"
	"os"
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

	cookie := ctx.Cookies("LokerHubCookie")
	if cookie == "" {
		log.Error("No authentication cookie found")
		return nil, errors.New("authentication required")
	}

	// Parse and validate the JWT token
	token, err := jwt.Parse(cookie, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid token signing method")
		}
		// Return the secret key used for signing
		secret := os.Getenv("JWT_SECRET")
		return []byte(secret), nil
	})

	if err != nil {
		log.Error("Error parsing token: ", err)
		return nil, errors.New("invalid authentication token")
	}

	// Extract claims from token
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		log.Error("Invalid token claims")
		return nil, errors.New("invalid authentication token")
	}

	userIdStr, ok := claims["id"].(string)
	if !ok {
		log.Error("User ID not found in token claims")
		return nil, errors.New("invalid authentication token")
	}
	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Error("Error parsing user ID: ", err)
		return nil, errors.New("invalid user information")
	}

	id, err := s.repo.FindByUserId(userId)
	if err != nil {
		log.Error("Error finding candidate by user ID: ", err)
		return nil, errors.New("failed to find candidate")
	}

	candidate := models.Candidate{
		ID:          id,
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

	if err := s.repo.UpdateWithTx(tx, &candidate); err != nil {
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

func (s *CandidateService) UploadCandidatePhoto(file *multipart.FileHeader) (string, error) {
	// Using the folder ID for recruiter photo
	return helpers.UploadPhoto(file, "1uGimZhfrohl_UefdkAEYH4j49Jmhn_hX")
}

func (s *CandidateService) UploadCandidateCV(file *multipart.FileHeader) (string, error) {
	return helpers.UploadPhoto(file, "1uGimZhfrohl_UefdkAEYH4j49Jmhn_hX")
}
