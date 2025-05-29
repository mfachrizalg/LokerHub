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

	// Get token from cookie instead of context locals
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

	// Extract user ID from claims
	userIdStr, ok := claims["id"].(string)
	if !ok {
		log.Error("User ID not found in token")
		return nil, errors.New("invalid user information")
	}

	userId, err := uuid.Parse(userIdStr)
	if err != nil {
		log.Error("Error parsing user ID: ", err)
		return nil, errors.New("invalid user ID format")
	}

	id, err := s.repo.FindByUserId(userId)
	if err != nil {
		log.Error("Error finding user: ", err)
		return nil, errors.New("invalid user ID")
	}

	recruiter := models.Recruiter{
		ID:        id,
		CompanyID: &req.CompanyID,
		UserID:    userId,
		Name:      req.Name,
		Position:  req.Position,
		PhotoURL:  req.PhotoURL,
		Handphone: req.Handphone,
	}

	if err := s.repo.UpdateWithTx(tx, &recruiter); err != nil {
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

func (s *RecruiterService) UploadRecruiterPhoto(file *multipart.FileHeader) (string, error) {
	// Using the folder ID for recruiter photos
	return helpers.UploadPhoto(file, "1uGimZhfrohl_UefdkAEYH4j49Jmhn_hX")
}
