package controllers

import (
	"backend/dtos"
	"backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type CompanyController struct {
	companyService *services.CompanyService
	validator      *validator.Validate
}

func NewCompanyController(companyService *services.CompanyService) *CompanyController {
	return &CompanyController{
		companyService: companyService,
		validator:      validator.New(),
	}
}

func (c *CompanyController) RegisterCompany(ctx *fiber.Ctx) error {
	var request dtos.RegisterCompanyRequest

	// Parse request body
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate request data
	if err := c.validator.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	// Process registration
	response, err := c.companyService.RegisterCompany(&request)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Registration failed",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}

func (c *CompanyController) GetAllCompany(ctx *fiber.Ctx) error {
	response, err := c.companyService.GetAllCompany()
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to fetch all company",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
}
