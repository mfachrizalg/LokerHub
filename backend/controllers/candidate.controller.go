package controllers

import (
	"backend/dtos"
	"backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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

	// Validate required fields
	if name == "" || description == "" || handphone == "" || education == "" || field == "" || location == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields",
		})
	}

	photo, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error retrieving photo file",
			"error":   err.Error(),
		})
	}

	cv, err := ctx.FormFile("cv")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error retrieving CV file",
			"error":   err.Error(),
		})
	}

	photoURL, err := c.candidateService.UploadCandidatePhoto(photo)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload photo",
			"error":   err.Error(),
		})
	}

	cvURL, err := c.candidateService.UploadCandidateCV(cv)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload CV",
			"error":   err.Error(),
		})
	}

	request := dtos.RegisterCandidateRequest{
		Name:        name,
		Description: description,
		Handphone:   handphone,
		Photo:       photoURL,
		Education:   education,
		Field:       field,
		Location:    location,
		CV:          cvURL,
	}

	// Validate request struct
	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.candidateService.RegisterCandidate(&request, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Register failed",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
