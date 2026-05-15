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
			return c.Status(code).JSON(fiber.Map{
				"error": err.Error(),
			})
		},
	})

	app.Use(recover.New())
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		workDir, _ := os.Getwd()
		frontendPath := filepath.Join(workDir, "frontend", "dist")
		return c.SendFile(frontendPath + "/index.html")
	})

	app.Static("/", filepath.Join("frontend", "dist"))

	app.Post("/api/generate", qrHandler.GenerateQR)
	app.Get("/health", qrHandler.Health)

	log.Println("Server starting on :9999")
	if err := app.Listen(":9999"); err != nil {
		log.Fatal(err)
	}
}
