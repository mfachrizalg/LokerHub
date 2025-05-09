package controllers

import (
	"backend/dtos"
	"backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
)

type RecruitController struct {
	recruitService *services.RecruiterService
	validate       *validator.Validate
}

func NewRecruitController(recruitService *services.RecruiterService) *RecruitController {
	return &RecruitController{
		recruitService: recruitService,
		validate:       validator.New(),
	}
}

func (c *RecruitController) RegisterRecruit(ctx *fiber.Ctx) error {
	// Get form values
	name := ctx.FormValue("name")
	position := ctx.FormValue("position")
	handphone := ctx.FormValue("handphone")
	companyIDStr := ctx.FormValue("company_id")

	// Validate required fields
	if name == "" || position == "" || handphone == "" || companyIDStr == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields",
		})
	}

	// Parse UUID
	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid company ID format",
			"error":   err.Error(),
		})
	}

	// Get file from form
	file, err := ctx.FormFile("photo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error retrieving file",
			"error":   err.Error(),
		})
	}

	// Handle file upload (assuming you'll implement this in a service)
	// This will need to be implemented in your service
	photoURL, err := c.recruitService.UploadRecruiterPhoto(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload image",
			"error":   err.Error(),
		})
	}

	// Create request object
	request := dtos.RegisterRecruiterRequest{
		Name:      name,
		Position:  position,
		Handphone: handphone,
		PhotoURL:  photoURL,
		CompanyID: companyID,
	}

	// Validate request struct
	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.recruitService.RegisterRecruiter(&request, ctx)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Register failed",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}

func (c *RecruitController) GetRecruiterDetail(ctx *fiber.Ctx) error {
	userID, ok := ctx.Locals("userID").(uuid.UUID)
	if !ok {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid user ID format",
		})
	}

	recruiter, err := c.recruitService.GetRecruiterDetail(userID)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to get recruiter detail",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(recruiter)
}
