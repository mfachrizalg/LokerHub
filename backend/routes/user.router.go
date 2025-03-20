package routes

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupUserRoutes(app *fiber.App) {
	userRepo := repositories.NewUserRepository()
	candidateRepo := repositories.NewCandidateRepository()
	recruiterRepo := repositories.NewRecruiterRepository()
	userService := services.NewUserService(userRepo, candidateRepo, recruiterRepo)
	userController := controllers.NewUserController(userService)

	// User routes
	userRoutes := app.Group("/api/users")
	userRoutes.Post("/register", userController.Register)
}
