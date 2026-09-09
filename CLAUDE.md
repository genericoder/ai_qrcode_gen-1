# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Project Overview

QR Code Generator — a full-stack web app with a Go/Fiber backend API and a React/Vite frontend SPA. The backend generates QR codes as base64-encoded PNGs; the frontend renders them from the API response.

## Commands

### Backend (Go)

```bash
make build   # compile to bin/server
make run     # build + run
make dev     # go run (no binary, for quick iteration)
make test    # go test -v ./...
make vet     # go vet ./...
make lint    # gofmt -l . (list unformatted files)
make fmt     # go fmt ./...
make clean   # remove bin/server
```

Run a single test:
```bash
go test -v ./internal/service -run TestValidateContent
```

### Frontend (from `frontend/`)

```bash
npm install
npm run dev      # Vite dev server with HMR
npm run build    # compile to frontend/dist/
npm run lint     # ESLint
npm run preview  # preview production build
```

## Architecture

### Request Flow

```text
User input → POST /api/generate → handler → service (validate + encode) → base64 PNG → frontend <img>
```

### Backend Layers (`internal/`)

- **handler/qr.go** — HTTP request parsing, response serialization, delegates to service
- **service/qr.go** — validation (non-empty, ≤4096 chars) and QR generation via `go-qrcode` (Medium error correction, 256×256)
- **model/qr.go** — `QRRequest` / `QRResponse` structs
- **cmd/server/main.go** — Fiber setup: CORS (origin allowlist via `ALLOWED_ORIGINS` env var), logger, error recovery, static file serving from `frontend/dist/`

### API

| Method | Path | Description |
|--------|------|-------------|
| POST | `/api/generate` | `{"content":"..."}` → `{"success":true,"data":{"content":"...","image":"<base64>"}}` |
| GET | `/health` | `{"status":"ok"}` |
| GET | `/` | serves `frontend/dist/index.html` |

### Frontend

Single component (`src/App.jsx`) manages all state with hooks. API base URL is read from the `VITE_API_URL` environment variable (defaults to `http://localhost:9999`).

### Static file serving

The backend serves `frontend/dist/` at the root. Build the frontend first (`npm run build`) before starting the backend if you want a unified server. In dev, run both servers separately (Vite on :5173, backend on :9999); the default `ALLOWED_ORIGINS` allowlist covers both.

## Testing

Tests cover handler (HTTP status codes, error cases) and service (validation boundaries, base64 output). CI runs build, test, vet, and format in parallel (GitHub Actions, Go 1.25).
