package model

// QRRequest represents the request body for QR code generation
type QRRequest struct {
	Content string `json:"content"`
}

// QRResponse represents the response for QR code generation
type QRResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}
