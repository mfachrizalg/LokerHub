package routes

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupAuthRoutes(app *fiber.App) {
	// Initialize dependencies
	authRepo := repositories.NewUserRepository()
	authService := services.NewAuthService(authRepo)
	authController := controllers.NewAuthController(authService)

	// Auth routes
	authRoutes := app.Group("/api/auth")
	authRoutes.Post("/login", authController.Login)
}
