package model

type QRRequest struct {
	Content string `json:"content"`
}

type QRData struct {
	Content string `json:"content"`
	Image   string `json:"image"`
}

type QRResponse struct {
	Success bool    `json:"success"`
	Message string  `json:"message,omitempty"`
	Data    *QRData `json:"data,omitempty"`
}
