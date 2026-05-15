package handler

import (
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

func (h *QRHandler) Index(c *fiber.Ctx) error {
	return c.SendString("QR Generator API is running. Use the frontend app.")
}

func (h *QRHandler) GenerateQR(c *fiber.Ctx) error {
	var req model.QRRequest
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.QRResponse{
			Success: false,
			Message: "Invalid request body",
		})
	}

	if err := h.qrService.ValidateContent(req.Content); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(model.QRResponse{
			Success: false,
			Message: err.Error(),
		})
	}

	image, err := h.qrService.GenerateQR(req.Content)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(model.QRResponse{
			Success: false,
			Message: "Failed to generate QR code",
		})
	}

	return c.JSON(fiber.Map{
		"success": true,
		"data": fiber.Map{
			"content": req.Content,
			"image":   image,
		},
	})
}

func (h *QRHandler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}
