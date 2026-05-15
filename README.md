# QR Code Generator

A simple Go-based QR code generator web application using the Fiber framework.

## Features

- Generate QR codes from any text or URL
- Clean web interface
- Health check endpoint

## Requirements

- Go 1.25 or later

## Getting Started

### Using Makefile

```bash
# Build the server
make build

# Run the server
make run

# Run tests
make test
```

### Manual Commands

```bash
# Build
go build -o bin/server ./cmd/server

# Run
./bin/server

# Test
go test -v ./...
```

The server runs on `http://localhost:9999`

## Endpoints

| Method | Path | Description |
|--------|------|-------------|
| GET | `/` | Web interface |
| POST | `/api/generate` | Generate QR code |
| GET | `/health` | Health check |

## Project Structure

```
.
├── cmd/server/       # Application entry point
├── internal/         # Internal packages
│   ├── handler/      # HTTP handlers
│   ├── service/      # Business logic
│   └── model/         # Data models
├── Makefile          # Build automation
└── README.md         # This file
```
