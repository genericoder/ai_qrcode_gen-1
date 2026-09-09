package service

import (
	"encoding/base64"
	"errors"
	"log/slog"

	"github.com/skip2/go-qrcode"
)

type QRService struct{}

func NewQRService() *QRService {
	return &QRService{}
}

func (s *QRService) ValidateContent(content string) error {
	if content == "" {
		return errors.New("content cannot be empty")
	}
	if len(content) > 4096 {
		return errors.New("content too long (max 4096 characters)")
	}
	return nil
}

func (s *QRService) GenerateQR(content string) (string, error) {
	slog.Debug("encoding QR code", "content_length", len(content))
	png, err := qrcode.Encode(content, qrcode.Medium, 256)
	if err != nil {
		return "", err
	}
	slog.Debug("QR code encoded successfully", "content_length", len(content), "png_bytes", len(png))
	return base64.StdEncoding.EncodeToString(png), nil
}
