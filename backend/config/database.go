package config

import (
	"backend/models"
	"fmt"
	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
)

var DB *gorm.DB

func ConnectDB() {
	e := godotenv.Load()
	if e != nil {
		log.Fatal("Error loading .env file")
	}

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_SSL"),
		os.Getenv("DB_TIMEZONE"),
	)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	errorMigration := db.AutoMigrate(
		&models.Candidate{},
		&models.Company{},
		&models.Job{},
		&models.Process{},
		&models.Recruiter{},
		&models.User{},
	)
	if errorMigration != nil {
		log.Fatal("Failed to migrate the database: ", errorMigration)
	}

	DB = db
	log.Debug("Connected to the database")
}
