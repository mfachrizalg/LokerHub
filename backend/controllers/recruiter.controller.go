package controllers

import (
	"backend/dtos"
	"backend/services"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
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
	var request dtos.RegisterRecruiterRequest

	// Parse request body
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate request
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
