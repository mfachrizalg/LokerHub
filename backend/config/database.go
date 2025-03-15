package config

import (
	"backend/models"
	"fmt"
	"os"
	"time"

	"github.com/gofiber/fiber/v2/log"
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

// DB is a global variable that holds the database connection
var DB *gorm.DB

// ConnectDB initializes the database connection and returns it
func ConnectDB() *gorm.DB {
	// Load environment variables
	if err := godotenv.Load(); err != nil {
		log.Warn("Error loading .env file, using environment variables")
	}

	// Validate required environment variables
	requiredEnvVars := []string{"DB_HOST", "DB_USERNAME", "DB_PASSWORD", "DB_NAME", "DB_PORT"}
	for _, envVar := range requiredEnvVars {
		if os.Getenv(envVar) == "" {
			log.Fatalf("Required environment variable %s is not set", envVar)
		}
	}

	// Set default values for optional environment variables
	sslMode := os.Getenv("DB_SSL")
	if sslMode == "" {
		sslMode = "disable"
	}

	timeZone := os.Getenv("DB_TIMEZONE")
	if timeZone == "" {
		timeZone = "UTC"
	}

	// Build connection string
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		os.Getenv("DB_HOST"),
		os.Getenv("DB_USERNAME"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_NAME"),
		os.Getenv("DB_PORT"),
		sslMode,
		timeZone,
	)

	// Connect to the database
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "",
			SingularTable: false,
		},
	})
	if err != nil {
		log.Fatal("Failed to connect to the database: ", err)
	}

	// Configure connection pool
	sqlDB, err := DB.DB()
	if err != nil {
		log.Fatal("Failed to get database connection pool: ", err)
	}

	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(time.Hour)

	// Run migrations
	if err := migrateDB(DB); err != nil {
		log.Fatal("Database migration failed: ", err)
	}

	log.Info("Successfully connected to the database")
	return DB
}

func migrateDB(db *gorm.DB) error {
	m := []interface{}{
		&models.Candidate{},
		&models.Company{},
		&models.Job{},
		&models.Process{},
		&models.Recruiter{},
		&models.User{},
	}

	return db.AutoMigrate(m...)
}

func CloseDB() {
	if DB != nil {
		sqlDB, err := DB.DB()
		if err != nil {
			log.Error("Failed to get database connection: ", err)
			return
		}
		if err := sqlDB.Close(); err != nil {
			log.Error("Error closing database connection: ", err)
		} else {
			log.Info("Database connection closed")
		}
	}
}
