# Repository Guidelines

## Project Structure

```
/cmd           # Entry points
/internal     # Internal packages (handler, service, model)
/tests         # Test files (co-located with internal packages)
```

- Keep source files organized by feature or module.
- Place tests alongside the code they cover in `_test` packages.

## Build, Test, and Development Commands

| Command | Description |
|---------|-------------|
| `go build -o bin/server ./cmd/server` | Build the server |
| `go test -v ./...` | Run all tests |
| `go run ./cmd/server` | Run the server |

## Coding Style & Naming Conventions

- **Indentation**: Go standard (Tabs)
- **Naming**: Use `camelCase` for variables/functions, `PascalCase` for types
- **Files**: Use kebab-case is NOT used - use snake_case or PascalCase (e.g., `qr_handler.go`)
- **Tools**: Run `go fmt` and `go vet` before committing

## Testing Guidelines

- Use Go's built-in testing package
- Test files use `_test.go` suffix
- Run tests with `go test -v ./...`

## Getting Started

1. Build: `go build -o bin/server ./cmd/server`
2. Run: `./bin/server` (starts on port 9999)
3. Test: `go test -v ./...`

## Endpoints

- `GET /` - Serve HTML frontend
- `POST /api/generate` - Generate QR code (accepts JSON `{"content": "..."}`)
- `GET /health` - Health check

For questions, open an issue or reach out to maintainers.
