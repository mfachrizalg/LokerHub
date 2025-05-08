package main

import (
	"backend/config"
	"backend/routes"
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
)

func setupRoutes(app *fiber.App) {
	routes.SetupUserRoutes(app)
	routes.SetupAuthRoutes(app)
	routes.SetupRecruiterRoutes(app)
	routes.SetupCandidateRoutes(app)
	routes.SetupCompanyRoutes(app)
	routes.SetupJobRoutes(app)
	routes.SetupProcessRoutes(app)
}

func main() {
	config.ConnectDB()
	defer config.CloseDB()

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	setupRoutes(app)

	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	log.Printf("Server started on port %s", port)
	app.Get("/", func(ctx *fiber.Ctx) error {
		return ctx.SendString("LokerHub API!")
	})

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("Server shutdown failed: ", err)
	}
}
