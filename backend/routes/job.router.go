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
	jobRoutes.Get("/detail/:jobId", jobController.GetDetailJob)
	jobRoutes.Use(middleware.JWTAuth())
	jobRoutes.Post("/create", middleware.RoleAuth("Recruiter"), jobController.CreateJob)
	jobRoutes.Put("/:jobId", middleware.RoleAuth("Recruiter"), jobController.UpdateJob)
	jobRoutes.Get("/recruiter", middleware.RoleAuth("Recruiter"), jobController.GetJobsByRecruiterID)
	jobRoutes.Post("/apply/:jobId", middleware.RoleAuth("Candidate"), jobController.ApplyJob)
}
