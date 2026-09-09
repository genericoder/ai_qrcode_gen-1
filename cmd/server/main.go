package main

import (
	"log/slog"
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
	level := slog.LevelInfo
	if os.Getenv("LOG_LEVEL") == "debug" {
		level = slog.LevelDebug
	}
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: level})))

	qrService := service.NewQRService()
	qrHandler := handler.NewQRHandler(qrService)

	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				code = e.Code
				slog.Warn("http error", "status", code, "path", c.Path(), "error", err)
			} else {
				slog.Error("unhandled error", "status", code, "path", c.Path(), "error", err)
			}
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
	slog.Info("CORS configured", "origins", allowedOrigins)
	app.Use(cors.New(cors.Config{
		AllowOrigins: allowedOrigins,
		AllowMethods: "GET,POST,OPTIONS",
		AllowHeaders: "Content-Type",
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		workDir, err := os.Getwd()
		if err != nil {
			slog.Error("failed to get working directory", "error", err)
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
	slog.Info("server starting", "port", port)
	if err := app.Listen(":" + port); err != nil {
		slog.Error("server failed to start", "error", err)
		os.Exit(1)
	}
}
