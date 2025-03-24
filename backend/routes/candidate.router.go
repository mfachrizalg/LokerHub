package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repositories"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupCandidateRoutes(app *fiber.App) {
	candidateRepo := repositories.NewCandidateRepository()
	candidateService := services.NewCandidateService(candidateRepo)
	candidateController := controllers.NewCandidateController(candidateService)

	candidateRoutes := app.Group("/api/candidates")
	candidateRoutes.Use(middleware.JWTAuth())
	candidateRoutes.Post("/register", middleware.RoleAuth("Candidate"), candidateController.RegisterCandidate)
}
