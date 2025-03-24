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
	var request dtos.RegisterCandidateRequest

	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	if err := c.validate.Struct(request); err != nil {
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

	return ctx.Status(fiber.StatusOK).JSON(response)
}
