package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repositories"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupRecruiterRoutes(app *fiber.App) {
	recruiterRepo := repositories.NewRecruiterRepository()
	recruiterService := services.NewRecruiterService(recruiterRepo)
	recruiterController := controllers.NewRecruitController(recruiterService)

	recruiterRoute := app.Group("/api/recruiters")
	recruiterRoute.Use(middleware.JWTAuth())
	recruiterRoute.Post("/register", middleware.RoleAuth("Recruiter"), recruiterController.RegisterRecruit)
}
