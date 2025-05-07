package controllers

import (
	"backend/dtos"
	"backend/services"
	"errors"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JobController struct {
	service *services.JobService
}

func NewJobController(service *services.JobService) *JobController {
	return &JobController{
		service: service,
	}
}

// GetAllJobs retrieves all jobs
func (c *JobController) GetAllJobs(ctx *fiber.Ctx) error {
	jobs, err := c.service.GetAllJobs()
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve jobs",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(jobs)
}

// UpdateJob updates an existing job
func (c *JobController) UpdateJob(ctx *fiber.Ctx) error {
	jobIDStr := ctx.Params("jobId")
	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Job ID format",
		})
	}

	var req dtos.UpdateJobRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	response, err := c.service.UpdateJob(jobID, &req, ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Job not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update job",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// CreateJob creates a new job
func (c *JobController) CreateJob(ctx *fiber.Ctx) error {
	var req dtos.CreateJobRequest
	if err := ctx.BodyParser(&req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	response, err := c.service.CreateJob(&req, ctx)
	if err != nil {
		// Handle specific errors from the service layer
		if err.Error() == "job already exists" {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": err.Error()})
		}
		if err.Error() == "invalid recruiter ID format" {
			return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{"message": err.Error()})
		}
		// Generic error for other issues
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to create job",
			"error":   err.Error(), // Avoid exposing too much detail in production
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

// ApplyJob handles the request for a candidate to apply for a job
func (c *JobController) ApplyJob(ctx *fiber.Ctx) error {
	jobIDStr := ctx.Params("jobId")
	jobID, err := uuid.Parse(jobIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid Job ID format",
		})
	}

	response, err := c.service.ApplyJob(jobID, ctx)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Job not found"})
		}
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to apply for job",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
