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

func TestIndex(t *testing.T) {
	app := fiber.New()
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)

	app.Get("/", qrHandler.Index)

	req := httptest.NewRequest(http.MethodGet, "/", nil)
	resp, err := app.Test(req)
	if err != nil {
		t.Fatal(err)
	}

	if resp.StatusCode != fiber.StatusOK {
		t.Errorf("Expected status %d, got %d", fiber.StatusOK, resp.StatusCode)
	}

	contentType := resp.Header.Get("Content-Type")
	if contentType != "text/html" {
		t.Errorf("Expected content-type text/html, got %s", contentType)
	}
}

func TestGenerateQR_Success(t *testing.T) {
	app := fiber.New()
	qrService := service.NewQRService()
	qrHandler := NewQRHandler(qrService)

	app.Post("/api/generate", qrHandler.GenerateQR)

	reqBody := model.QRRequest{Content: "https://example.com"}
	body, _ := json.Marshal(reqBody)

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
	json.NewDecoder(resp.Body).Decode(&respBody)

	if !respBody.Success {
		t.Error("Expected success to be true")
	}
	if respBody.Data != "https://example.com" {
		t.Errorf("Expected data to be 'https://example.com', got %s", respBody.Data)
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
	body, _ := json.Marshal(reqBody)

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
	json.NewDecoder(resp.Body).Decode(&respBody)

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
