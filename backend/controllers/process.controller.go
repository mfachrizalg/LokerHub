package controllers

import (
	"backend/services"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type ProcessController struct {
	service *services.ProcessService
}

func NewProcessController(service *services.ProcessService) *ProcessController {
	return &ProcessController{
		service: service,
	}
}

// GetAllProcessesByJob handles the request to get all processes for a job
func (c *ProcessController) GetAllProcessesByJob(ctx *fiber.Ctx) error {
	jobIDStr := ctx.Params("jobId") // Match parameter name in the route
	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Job ID format",
		})
	}

	// Call the service method, passing the context for authorization
	processes, err := c.service.GetAllProcessesByJob(jobID, ctx)
	if err != nil {
		// Handle specific errors from the service
		errMsg := err.Error()
		if errors.Is(err, gorm.ErrRecordNotFound) || errMsg == "job not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Job not found"})
		}
		if errMsg == "unauthorized to view processes for this job" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
		}
		// Generic internal error
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve processes",
			// "error": err.Error(), // Avoid exposing details in production
		})
	}

	// Return the list of process DTOs
	return ctx.Status(fiber.StatusOK).JSON(processes)
}

// Add other process controller methods here
