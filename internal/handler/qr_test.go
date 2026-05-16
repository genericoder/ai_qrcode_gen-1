package handler

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"

	"qr-generator/internal/model"
	"qr-generator/internal/service"
)

func TestNewQRHandler(t *testing.T) {
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)
	if qrHandler == nil {
		t.Error("NewQRHandler should not return nil")
	}
}


func TestGenerateQR_Success(t *testing.T) {
	app := fiber.New()
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)

	app.Post("/api/generate", qrHandler.GenerateQR)

	reqBody := model.QRRequest{Content: "https://example.com"}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/generate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}

	var respBody model.QRResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if !respBody.Success {
		t.Error("Expected success to be true")
	}
	if respBody.Data == nil {
		t.Fatal("Expected data to be non-nil")
	}
	if respBody.Data.Content != "https://example.com" {
		t.Errorf("Expected data.content to be 'https://example.com', got %s", respBody.Data.Content)
	}
	if respBody.Data.Image == "" {
		t.Error("Expected data.image to be non-empty")
	}
}

func TestGenerateQR_InvalidJSON(t *testing.T) {
	app := fiber.New()
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)

	app.Post("/api/generate", qrHandler.GenerateQR)

	req := httptest.NewRequest(http.MethodPost, "/api/generate", bytes.NewReader([]byte("invalid json")))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
	}

	var respBody model.QRResponse
	json.NewDecoder(resp.Body).Decode(&respBody)

	if respBody.Success {
		t.Error("Expected success to be false")
	}
}

func TestGenerateQR_EmptyContent(t *testing.T) {
	app := fiber.New()
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)

	app.Post("/api/generate", qrHandler.GenerateQR)

	reqBody := model.QRRequest{Content: ""}
	body, err := json.Marshal(reqBody)
	if err != nil {
		t.Fatal(err)
	}

	req := httptest.NewRequest(http.MethodPost, "/api/generate", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusBadRequest {
		t.Errorf("Expected status %d, got %d", fiber.StatusBadRequest, resp.StatusCode)
	}

	var respBody model.QRResponse
	if err := json.NewDecoder(resp.Body).Decode(&respBody); err != nil {
		t.Fatal(err)
	}

	if respBody.Success {
		t.Error("Expected success to be false")
	}
}

func TestGenerateQR_MethodNotAllowed(t *testing.T) {
	app := fiber.New()
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)

	app.Post("/api/generate", qrHandler.GenerateQR)

	req := httptest.NewRequest(http.MethodGet, "/api/generate", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	// Fiber returns 405 Method Not Allowed for POST-only routes
	if resp.StatusCode != fiber.StatusMethodNotAllowed {
		t.Errorf("Expected status %d, got %d", fiber.StatusMethodNotAllowed, resp.StatusCode)
	}
}

func TestHealth(t *testing.T) {
	app := fiber.New()
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)

	app.Get("/health", qrHandler.Health)

	req := httptest.NewRequest(http.MethodGet, "/health", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}
}
