package controllers

import (
	"backend/dtos"
	"backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
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
	name := ctx.FormValue("name")
	location := ctx.FormValue("location")
	industry := ctx.FormValue("industry")

	if name == "" || location == "" || industry == "" {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Missing required fields",
		})
	}

	file, err := ctx.FormFile("logo")
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Error retrieving file",
			"error":   err.Error(),
		})
	}

	logoURL, err := c.companyService.UploadCompanyLogo(file)
	if err != nil {
		return ctx.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": "Failed to upload image",
			"error":   err.Error(),
		})
	}

	req := dtos.RegisterCompanyRequest{
		Name:     name,
		Location: location,
		Industry: industry,
		Logo:     logoURL,
	}

	if err := c.validator.Struct(req); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	response, err := c.companyService.RegisterCompany(&req)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to register company",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(response)
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

func (c *CompanyController) GetDetailCompany(ctx *fiber.Ctx) error {
	companyIDStr := ctx.Params("companyID")
	companyID, err := uuid.Parse(companyIDStr)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid company ID",
			"error":   err.Error(),
		})
	}

	company, err := c.companyService.GetCompanyDetail(&companyID)
	if err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Failed to fetch company detail",
			"error":   err.Error(),
		})
	}

	if company == nil {
		return ctx.Status(fiber.StatusNotFound).JSON(fiber.Map{
			"message": "Company not found",
		})
	}

	return ctx.Status(fiber.StatusOK).JSON(company)
}
