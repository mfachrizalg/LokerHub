package controllers

import (
	"backend/dtos"
	"backend/services"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

type UserController struct {
	userService *services.UserService
	validate    *validator.Validate
}

func NewUserController(userService *services.UserService) *UserController {
	return &UserController{
		userService: userService,
		validate:    validator.New(),
	}
}

func (c *UserController) Register(ctx *fiber.Ctx) error {
	var request dtos.RegisterRequest

	// Parse request body
	if err := ctx.BodyParser(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Invalid request body",
			"error":   err.Error(),
		})
	}

	// Validate request data
	if err := c.validate.Struct(&request); err != nil {
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Validation failed",
			"error":   err.Error(),
		})
	}

	// Process registration
	response, err := c.userService.Register(&request)
	if err != nil {
		if err.Error() == "email already registered" {
			return ctx.Status(fiber.StatusConflict).JSON(fiber.Map{
				"message": "Registration failed",
				"error":   err.Error(),
			})
		}
		return ctx.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"message": "Registration failed",
			"error":   err.Error(),
		})
	}

	return ctx.Status(fiber.StatusCreated).JSON(response)
}
