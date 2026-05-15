package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"qr-generator/internal/handler"
	"qr-generator/internal/service"
)

func main() {
	// Initialize services
	qrService := service.NewQRService()

	// Initialize handlers
	qrHandler := handler.NewQRHandler(qrService)

	// Create Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	// Middleware
	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New())

	// Routes
	app.Get("/", qrHandler.Index)
	app.Post("/api/generate", qrHandler.GenerateQR)
	app.Get("/health", qrHandler.Health)

	// Start server
	log.Println("Server starting on :9999")
	if err := app.Listen(":9999"); err != nil {
		log.Fatal(err)
	}
}
