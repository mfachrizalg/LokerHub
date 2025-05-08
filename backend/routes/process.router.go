package routes

import (
	"backend/controllers"
	"backend/repositories"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupProcessRoutes(app *fiber.App) {
	jobRepo := repositories.NewJobRepository()
	processRepo := repositories.NewProcessRepository()
	processService := services.NewProcessService(processRepo, jobRepo)
	processController := controllers.NewProcessController(processService)

	processRoutes := app.Group("/process")
	processRoutes.Get("/job/:jobId", processController.GetAllProcessesByJob) // Get all processes for a job
}
