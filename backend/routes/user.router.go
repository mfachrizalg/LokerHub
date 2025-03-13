package routes

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	// Initialize dependencies
	userRepo := repositories.NewUserRepository()
	userService := services.NewUserService(userRepo)
	userController := controllers.NewUserController(userService)

	// User routes
	userRoutes := app.Group("/api/users")
	userRoutes.Post("/register", userController.Register)
}
