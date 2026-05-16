package main

import (
	"log"
	"os"
	"path/filepath"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"qr-generator/internal/handler"
	"qr-generator/internal/service"
)

func main() {
	qrService := service.NewQRService()
	qrHandler := handler.NewQRHandler(qrService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
			}
			log.Printf("error: %v", err)
			return c.Status(code).JSON(fiber.Map{
				"error": "internal server error",
			})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())

	allowedOrigins := os.Getenv("ALLOWED_ORIGINS")
	if allowedOrigins == "" {
		allowedOrigins = "http://localhost:5173,http://localhost:9999"
	}
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Content-Type",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		workDir, err := os.Getwd()
		if err != nil {
			return err
		}
		return c.SendFile(filepath.Join(workDir, "frontend", "dist", "index.html"))
	})

	app.Static("/", filepath.Join("frontend", "dist"))

	app.Post("/api/generate", qrHandler.GenerateQR)
	app.Get("/health", qrHandler.Health)

	port := os.Getenv("PORT")
	if port == "" {
		port = "9999"
	}
	log.Printf("Server starting on :%s", port)
	if err := app.Listen(":" + port); err != nil {
		log.Fatal(err)
	}
}
