package routes

import (
	"backend/controllers"
	"backend/middleware"
	"backend/repositories"
	"backend/services"
	"github.com/gofiber/fiber/v2"
)

func SetupCompanyRoutes(app *fiber.App) {
	companyRepo := repositories.NewCompanyRepository()
	companyService := services.NewCompanyService(companyRepo)
	companyController := controllers.NewCompanyController(companyService)

	companyRoute := app.Group("/api/companies")
	companyRoute.Use(middleware.JWTAuth())
	companyRoute.Post("/register", middleware.RoleAuth("Recruiter"), companyController.RegisterCompany)
	companyRoute.Get("/", companyController.GetAllCompany)
}
