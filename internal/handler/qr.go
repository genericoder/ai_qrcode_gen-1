package handler

import (
	"log/slog"

	"github.com/gofiber/fiber/v2"

	"qr-generator/internal/model"
	"qr-generator/internal/service"
)

type QRHandler struct {
	qrService *service.QRService
}

func NewQRHandler(qrService *service.QRService) *QRHandler {
	return &QRHandler{qrService: qrService}
}

func (h *QRHandler) GenerateQR(c *fiber.Ctx) error {
	var req model.QRRequest
	if err := c.BodyParser(&req); err != nil {
		slog.Warn("failed to parse request body", "error", err, "ip", c.IP())
		return c.Status(fiber.StatusBadRequest).JSON(model.QRResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	if err := h.qrService.ValidateContent(req.Content); err != nil {
		slog.Warn("content validation failed", "error", err, "content_length", len(req.Content), "ip", c.IP())
		return c.Status(fiber.StatusBadRequest).JSON(model.QRResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	image, err := h.qrService.GenerateQR(req.Content)
	if err != nil {
		slog.Error("QR generation failed", "error", err, "content_length", len(req.Content), "ip", c.IP())
		return c.Status(fiber.StatusInternalServerError).JSON(model.QRResponse{
			Success: false,
			Message: "Failed to generate QR code",
		})
	}

	slog.Debug("QR code generated", "content_length", len(req.Content), "ip", c.IP())
	return c.JSON(model.QRResponse{
		Success: true,
		Data: &model.QRData{
			Content: req.Content,
			Image:   image,
		},
	})
}

func (h *QRHandler) Health(c *fiber.Ctx) error {
	slog.Debug("health check")
	return c.JSON(map[string]string{"status": "ok"})
}
