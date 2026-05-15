package service

import (
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
