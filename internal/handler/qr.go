package handler

import (
	"github.com/gofiber/fiber/v2"

	"qr-generator/internal/model"
	"qr-generator/internal/service"
)

// QRHandler handles QR code related HTTP requests
type QRHandler struct {
	qrService *service.QRService
}

// NewQRHandler creates a new QR handler
func NewQRHandler(qrService *service.QRService) *QRHandler {
	return &QRHandler{qrService: qrService}
}

// Index serves the main HTML page
func (h *QRHandler) Index(c *fiber.Ctx) error {
	c.Set("Content-Type", "text/html")
	return c.SendString(indexHTML)
}

// GenerateQR handles QR code generation requests
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

	// Return the content - QR generation is done client-side via JavaScript
	return c.JSON(model.QRResponse{
		Success: true,
		Data:    req.Content,
	})
}

// Health handles health check requests
func (h *QRHandler) Health(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{"status": "ok"})
}

const indexHTML = `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>QR Code Generator</title>
    <script src="https://cdnjs.cloudflare.com/ajax/libs/qrcodejs/1.0.0/qrcode.min.js"></script>
    <style>
        * {
            box-sizing: border-box;
            margin: 0;
            padding: 0;
        }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            display: flex;
            justify-content: center;
            align-items: center;
            padding: 20px;
        }
        .container {
            background: white;
            border-radius: 16px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            padding: 40px;
            width: 100%;
            max-width: 500px;
        }
        h1 {
            color: #333;
            margin-bottom: 30px;
            text-align: center;
            font-size: 28px;
        }
        .input-group {
            margin-bottom: 20px;
        }
        label {
            display: block;
            margin-bottom: 8px;
            color: #555;
            font-weight: 500;
        }
        input, textarea {
            width: 100%;
            padding: 14px;
            border: 2px solid #e0e0e0;
            border-radius: 8px;
            font-size: 16px;
            transition: border-color 0.3s;
        }
        input:focus, textarea:focus {
            outline: none;
            border-color: #667eea;
        }
        textarea {
            resize: vertical;
            min-height: 100px;
        }
        button {
            width: 100%;
            padding: 16px;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            border: none;
            border-radius: 8px;
            font-size: 18px;
            font-weight: 600;
            cursor: pointer;
            transition: transform 0.2s, box-shadow 0.2s;
        }
        button:hover {
            transform: translateY(-2px);
            box-shadow: 0 8px 20px rgba(102, 126, 234, 0.4);
        }
        button:active {
            transform: translateY(0);
        }
        #result {
            margin-top: 30px;
            text-align: center;
        }
        #qrcode {
            display: inline-block;
            margin-top: 20px;
        }
        #qrcode img {
            border-radius: 8px;
            box-shadow: 0 4px 12px rgba(0,0,0,0.1);
        }
        .error {
            color: #e74c3c;
            margin-top: 10px;
            text-align: center;
        }
        .success-message {
            color: #27ae60;
            margin-top: 10px;
            text-align: center;
        }
    </style>
</head>
<body>
    <div class="container">
        <h1>QR Code Generator</h1>
        <div class="input-group">
            <label for="content">Enter content (URL, text, etc.)</label>
            <textarea id="content" placeholder="https://example.com or Hello World"></textarea>
        </div>
        <button onclick="generateQR()">Generate QR Code</button>
        <div id="result">
            <div id="qrcode"></div>
            <p id="message"></p>
        </div>
    </div>
    <script>
        function generateQR() {
            var content = document.getElementById('content').value.trim();
            var qrContainer = document.getElementById('qrcode');
            var messageEl = document.getElementById('message');
            
            messageEl.textContent = '';
            messageEl.className = '';
            qrContainer.innerHTML = '';
            
            if (!content) {
                messageEl.textContent = 'Please enter some content';
                messageEl.className = 'error';
                return;
            }
            
            try {
                if (typeof QRCode === 'undefined') {
                    messageEl.textContent = 'Error: QRCode library not loaded';
                    messageEl.className = 'error';
                    return;
                }
                
                new QRCode(qrContainer, {
                    text: content,
                    width: 300,
                    height: 300,
                    colorDark : "#000000",
                    colorLight : "#ffffff",
                    correctLevel : QRCode.CorrectLevel.H
                });
                messageEl.textContent = 'QR Code generated successfully!';
                messageEl.className = 'success-message';
            } catch (err) {
                messageEl.textContent = 'Error generating QR code: ' + err.message;
                messageEl.className = 'error';
            }
        }
        
        // Generate QR on Enter key
        document.getElementById('content').addEventListener('keypress', function(e) {
            if (e.key === 'Enter' && !e.shiftKey) {
                e.preventDefault();
                generateQR();
            }
        });
    </script>
</body>
</html>`
