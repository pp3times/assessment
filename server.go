package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/gofiber/fiber/v2"
	"github.com/pp3times/assessment/models"
	"github.com/pp3times/assessment/service"
	"github.com/pp3times/assessment/storage"
)

func main() {
	fmt.Println("Please use server.go for main file")
	fmt.Println("start at port:", os.Getenv("PORT"))
	// err := godotenv.Load(".env")

	// if err != nil {

	// 	log.Fatal(err)

	// }

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

	// Gracefully shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	go func() {
		_ = <-c
		fmt.Println("Gracefully shutting down...")
		_ = app.Shutdown()
	}()

	if err := app.Listen(os.Getenv("PORT")); err != nil {
		log.Panic(err)
	}
}
