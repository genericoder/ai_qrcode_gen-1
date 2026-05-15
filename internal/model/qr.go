package model

type QRRequest struct {
	Content string `json:"content"`
}

type QRResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message,omitempty"`
	Data    string `json:"data,omitempty"`
}

type QRData struct {
	Content string `json:"content"`
	Image   string `json:"image"` // Base64 encoded PNG
}
