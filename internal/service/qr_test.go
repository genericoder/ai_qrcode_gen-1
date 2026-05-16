package service

import (
	"encoding/base64"
	"strings"
	"testing"
)

func TestNewQRService(t *testing.T) {
	qrService := NewQRService()
	if qrService == nil {
		t.Error("NewQRService should not return nil")
	}
}

func TestValidateContent(t *testing.T) {
	qrService := NewQRService()

	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{
			name:    "valid content",
			content: "https://example.com",
			wantErr: false,
		},
		{
			name:    "valid text content",
			content: "Hello World",
			wantErr: false,
		},
		{
			name:    "empty content",
			content: "",
			wantErr: true,
		},
		{
			name:    "content too long",
			content: string(make([]byte, 4097)),
			wantErr: true,
		},
		{
			name:    "max allowed length",
			content: string(make([]byte, 4096)),
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := qrService.ValidateContent(tt.content)
			if (err != nil) != tt.wantErr {
				t.Errorf("ValidateContent() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestGenerateQR(t *testing.T) {
	qrService := NewQRService()

	tests := []struct {
		name    string
		content string
		wantErr bool
	}{
		{name: "url", content: "https://example.com", wantErr: false},
		{name: "plain text", content: "Hello World", wantErr: false},
		{name: "single character", content: "a", wantErr: false},
		{name: "long content", content: strings.Repeat("a", 1000), wantErr: false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := qrService.GenerateQR(tt.content)
			if (err != nil) != tt.wantErr {
				t.Fatalf("GenerateQR() error = %v, wantErr %v", err, tt.wantErr)
			}
			if tt.wantErr {
				return
			}

			if result == "" {
				t.Fatal("GenerateQR() returned empty string")
			}

			decoded, err := base64.StdEncoding.DecodeString(result)
			if err != nil {
				t.Fatalf("GenerateQR() output is not valid base64: %v", err)
			}

			// PNG magic bytes: 0x89 0x50 0x4E 0x47
			if len(decoded) < 4 || string(decoded[:4]) != "\x89PNG" {
				t.Error("GenerateQR() output does not decode to a PNG")
			}
		})
	}
}

func TestGenerateQR_OutputIsDeterministic(t *testing.T) {
	qrService := NewQRService()
	content := "https://example.com"

	first, err := qrService.GenerateQR(content)
	if err != nil {
		t.Fatal(err)
	}
	second, err := qrService.GenerateQR(content)
	if err != nil {
		t.Fatal(err)
	}

	if first != second {
		t.Error("GenerateQR() produced different output for the same input")
	}
}
