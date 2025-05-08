package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repositories"
	"backend/services"

	"github.com/gofiber/fiber/v2"
)

func SetupJobRoutes(app *fiber.App) {
	jobRepo := repositories.NewJobRepository()
	processRepo := repositories.NewProcessRepository()
	jobService := services.NewJobService(jobRepo, processRepo)
	jobController := controllers.NewJobController(jobService)

	jobRoutes := app.Group("/api/jobs")
	jobRoutes.Get("/", jobController.GetAllJobs)
	jobRoutes.Use(middleware.JWTAuth())
	jobRoutes.Post("/", middleware.RoleAuth("Recruiter"), jobController.CreateJob)
	jobRoutes.Patch("/:id", middleware.RoleAuth("Recruiter"), jobController.UpdateJob)
	jobRoutes.Post("/apply/:id", middleware.RoleAuth("Candidate"), jobController.ApplyJob)
}
