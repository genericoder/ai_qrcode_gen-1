package service

import (
	"errors"
)

// QRService handles QR code generation logic
type QRService struct{}

// NewQRService creates a new QR service
func NewQRService() *QRService {
	return &QRService{}
}

// ValidateContent validates the content for QR code generation
func (s *QRService) ValidateContent(content string) error {
	if content == "" {
		return errors.New("content cannot be empty")
	}
	if len(content) > 4096 {
		return errors.New("content too long (max 4096 characters)")
	}
	return nil
}
