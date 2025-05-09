package controllers

import (
	"backend/dtos"
	"backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type CandidateController struct {
	candidateService *services.CandidateService
	validate         *validator.Validate
}

func NewCandidateController(candidateService *services.CandidateService) *CandidateController {
	return &CandidateController{
		candidateService: candidateService,
		validate:         validator.New(),
	}
}

func (c *CandidateController) RegisterCandidate(ctx *fiber.Ctx) error {
	// Get form values
	name := ctx.FormValue("name")
	description := ctx.FormValue("description")
	handphone := ctx.FormValue("handphone")
	education := ctx.FormValue("education")
	field := ctx.FormValue("field")
	location := ctx.FormValue("location")

	// Get CV from form
	cv, err := ctx.FormFile("cv")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error retrieving file",
			"error":   err.Error(),
		})
	}

	// Get photo from form
	photo, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error retrieving file",
			"error":   err.Error(),
		})
	}
	// Validate required fields
	if name == "" || description == "" || handphone == "" || education == "" || field == "" || location == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields",
		})
	}
	// Handle photo upload
	photoURL, err := c.candidateService.UploadCandidatePhoto(photo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload image",
			"error":   err.Error(),
		})
	}

	// Handle CV upload
	cvURL, err := c.candidateService.UploadCandidateCV(cv)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload CV",
			"error":   err.Error(),
		})
	}

	// Create request object
	req := &dtos.RegisterCandidateRequest{
		Name:        name,
		Description: description,
		Handphone:   handphone,
		Photo:       photoURL,
		Education:   education,
		Field:       field,
		Location:    location,
		CV:          cvURL,
	}
	// Validate request
	if err := c.validate.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}
	// Call service to register candidate
	response, err := c.candidateService.RegisterCandidate(req, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to register candidate",
			"error":   err.Error(),
		})
	}
	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *CandidateController) GetCandidateDetail(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID format",
		})
	}

	candidate, err := c.candidateService.GetCandidateDetail(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get candidate details",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(candidate)
}
