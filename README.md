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

## CI/CD

This project uses **GitHub Actions** for continuous integration. The workflow is defined in `.github/workflows/ci.yml`.

### Workflow Jobs

| Job | Description | Trigger |
|-----|-------------|---------|
| **build** | Compiles the Go application | Runs on every push/PR |
| **test** | Runs unit tests | Runs after build succeeds |
| **vet** | Runs `go vet` for static analysis | Runs after build succeeds |
| **format** | Checks code formatting with `gofmt` | Runs after build succeeds |

### Running CI Jobs Manually

Developers can trigger workflow runs manually from GitHub:

1. Go to the **Actions** tab in the repository
2. Select the **CI** workflow
3. Click **Run workflow** → **Run workflow**

This allows running the full CI pipeline or checking individual job status directly from GitHub.

## Project Structure

```
.
├── .github/workflows/  # GitHub Actions CI/CD
├── cmd/server/         # Application entry point
├── internal/           # Internal packages
│   ├── handler/        # HTTP handlers
│   ├── service/        # Business logic
│   └── model/          # Data models
├── Makefile            # Build automation
└── README.md           # This file
```
