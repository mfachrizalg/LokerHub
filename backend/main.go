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

func main() {
	config.ConnectDB()
	defer config.CloseDB()

	// Create Fiber app with custom config
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Register middlewares
	app.Use(logger.New())
	app.Use(recover.New())
	app.Use(cors.New())

	// Setup routes
	setupRoutes(app)

	// Default port or from environment
	port := os.Getenv("PORT")
	if port == "" {
		port = "3000"
	}

	// Start server with graceful shutdown
	go func() {
		if err := app.Listen(":" + port); err != nil {
			log.Fatal("Error starting server: ", err)
		}
	}()

	log.Printf("Server started on port %s", port)

	// Graceful shutdown
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, syscall.SIGTERM)
	<-quit

	log.Println("Shutting down server...")
	if err := app.Shutdown(); err != nil {
		log.Fatal("Server shutdown failed: ", err)
	}
}

// setupRoutes configures all API routes
func setupRoutes(app *fiber.App) {
	// Set up user routes for registration
	routes.SetupUserRoutes(app)

	// Set up auth routes for login
	routes.SetupAuthRoutes(app)
}
