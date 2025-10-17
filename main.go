package main

import (
	"log"
	"workshop3/internal/config"
	"workshop3/internal/database"
	"workshop3/internal/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Initialize database
	db, err := database.Init(cfg.DBPath)
	if err != nil {
		log.Fatalf("Failed to initialize database: %v", err)
	}

	// Create Fiber app
	app := fiber.New()

	// Health check endpoint
	app.Get("/health", func(c *fiber.Ctx) error {
		return c.JSON(fiber.Map{"status": "ok"})
	})

	// Register routes
	routes.RegisterUserRoutes(app, db)

	// Start server
	log.Printf("Starting server on %s", cfg.ServerPort)
	if err := app.Listen(cfg.ServerPort); err != nil {
		log.Fatalf("Server error: %v", err)
	}
}
