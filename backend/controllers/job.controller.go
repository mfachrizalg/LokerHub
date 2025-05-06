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

// GetAllJobs retrieves all jobs (without pagination)
func (c *JobController) GetAllJobs(ctx *fiber.Ctx) error {
	// Get jobs from service (assuming service method is updated to not use pagination)
	jobs, err := c.service.GetAllJobs() // Call service without page/limit
	if err != nil {
		// Use a more generic error message for internal server errors
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to retrieve jobs",
			// "error": err.Error(), // Avoid exposing internal errors in production
		})
	}

	// Convert to response format
	jobResponses := make([]dtos.JobResponse, len(jobs))
	for i, job := range jobs {
		jobResponses[i] = dtos.JobResponse{
			ID:            job.ID,
			CompanyID:     job.CompanyID,
			RecruiterID:   job.RecruiterID,
			Name:          job.Name,
			Criteria:      job.Criteria,
			Qualification: job.Qualification,
			Status:        job.Status,
		}
	}

	// Return the list of jobs directly as a JSON array
	return ctx.Status(fiber.StatusOK).JSON(jobResponses)
}

// UpdateJob updates an existing job
func (c *JobController) UpdateJob(ctx *fiber.Ctx) error {
	jobIDStr := ctx.Params("id") // Assuming the route parameter is named "id"
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

	// Call the service to update the job
	response, err := c.service.UpdateJob(jobID, &req, ctx)
	if err != nil {
		// Handle specific errors from the service layer
		if errors.Is(err, gorm.ErrRecordNotFound) || err.Error() == "job not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": "Job not found"})
		}
		if err.Error() == "unauthorized to update this job" {
			return ctx.Status(fiber.StatusForbidden).JSON(fiber.Map{"message": err.Error()})
		}
		// Generic error for other issues
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to update job",
			"error":   err.Error(), // Avoid exposing too much detail in production
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

// CreateJob creates a new job
func (c *JobController) CreateJob(ctx *fiber.Ctx) error {
	var req *dtos.CreateJobRequest

	// Parse and validate request
	if err := ctx.BodyParser(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
		})
	}

	// Validate required fields (could use validator package here)
	if req.Name == "" || req.Criteria == "" || req.Qualification == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields",
		})
	}

	// Create job
	response, err := c.service.CreateJob(req, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": err.Error(),
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
		// Check for specific user-facing errors
		if err.Error() == "job not found" {
			return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{"message": err.Error()})
		}
		if err.Error() == "already applied for this job" {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{"message": err.Error()})
		}
		// Generic internal error for others
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to apply for job", // Avoid exposing internal details
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
