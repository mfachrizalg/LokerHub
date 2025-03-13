package services

import (
	"backend/dtos"
	"backend/models"
	"backend/repositories"
	"errors"
	"github.com/gofiber/fiber/v2/log"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"os"
	"time"
)

type AuthService struct {
	repo *repositories.UserRepository
}

func NewAuthService(repo *repositories.UserRepository) *AuthService {
	return &AuthService{
		repo: repo,
	}
}

func (s *AuthService) Login(req *dtos.LoginRequest) (*dtos.LoginResponse, error) {
	var user models.User

	// Find user by email
	existingUser, err := s.repo.FindByEmail(req.Email)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Database error: ", err)
		return nil, errors.New("failed to process registration")
	}

	if existingUser == nil {
		return nil, errors.New("email not found")
	}

	// Verify password
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("invalid email or password")
	}

	// Generate JWT token
	token, err := generateJWTToken(user)
	if err != nil {
		return nil, errors.New("failed to generate token")
	}

	return &dtos.LoginResponse{
		Token:   token,
		Message: "Login successful",
	}, nil
}

func (s *AuthService) Logout() {

}

// generateJWTToken creates a new JWT token for the authenticated user
func generateJWTToken(user models.User) (string, error) {
	// Set expiration time
	expTime := time.Now().Add(1 * time.Hour)

	// Create claims
	claims := jwt.MapClaims{
		"id":    user.ID,
		"email": user.Email,
		"role":  user.Role,
		"exp":   expTime.Unix(),
	}

	// Create token with claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate signed token
	secret := os.Getenv("JWT_SECRET")
	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}
