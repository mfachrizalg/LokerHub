package main

import (
	"backend/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

func main() {
	config.ConnectDB()
	app := fiber.New()
	err := app.Listen(":3000")
	if err != nil {
		log.Fatal("Error starting server: ", err)
	}
}
