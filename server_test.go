package main

import (
	"log"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
	"github.com/pp3times/assessment/models"
	"github.com/pp3times/assessment/service"
	"github.com/pp3times/assessment/storage"
	"github.com/stretchr/testify/assert"
)

func TestGetExpenseRoute(t *testing.T) {

	// Define Fiber app.
	err := godotenv.Load(".env")

	if err != nil {

		log.Fatal(err)

	}

	config := &storage.Config{

		Host: os.Getenv("DB_HOST"),

		Port: os.Getenv("DB_PORT"),

		Password: os.Getenv("DB_PASS"),

		User: os.Getenv("DB_USER"),

		SSLMode: os.Getenv("DB_SSLMODE"),

		DBName: os.Getenv("DB_NAME"),
	}

	db, err := storage.NewConnection(config)

	if err != nil {

		log.Fatal("could not load database")

	}

	err = models.MigrateBooks(db)

	if err != nil {

		log.Fatal("could not migrate db")

	}

	r := &service.Repository{

		DB: db,
	}

	app := fiber.New()

	r.SetupRoutes(app)

	req := httptest.NewRequest("GET", "/expenses", nil)
	resp, _ := app.Test(req, 100)
	assert.Equalf(t, 200, resp.StatusCode, "Expected status code %d, got %d", 200, resp.StatusCode)
}
