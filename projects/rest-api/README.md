# Project — REST API

> **Phase 5 · Week 9–10**  
> Build a production-ready REST API using Go's standard library (or `chi`).

---

## Goal

Build a JSON REST API from scratch demonstrating:
- Clean routing with Go 1.22 `ServeMux` or `chi`
- Middleware chain (logging, auth, request ID)
- Request validation
- JSON encode/decode
- `database/sql` or in-memory store
- Graceful shutdown
- Structured logging with `log/slog`

---

## Suggested Project: `bookshelf` — a simple book inventory API

```
POST   /books           → add a book
GET    /books           → list all books
GET    /books/{id}      → get book by ID
PUT    /books/{id}      → update a book
DELETE /books/{id}      → delete a book
GET    /health          → health check endpoint
```

---

## Minimum Requirements

- [ ] At least 5 endpoints (CRUD + health)
- [ ] JSON request/response with proper struct tags
- [ ] Middleware: request logging, request ID header
- [ ] Proper HTTP status codes (200, 201, 400, 404, 500)
- [ ] Graceful shutdown (`http.Server.Shutdown` with context timeout)
- [ ] In-memory store protected by `sync.RWMutex` (or use SQLite)
- [ ] `httptest`-based tests for at least 3 handlers
- [ ] Passes `go vet ./...` and `go test ./...`

---

## Getting Started

```bash
cd projects/rest-api
go mod init github.com/rohitdhiman/golang-learning/projects/rest-api

# Optional — chi router (if not using Go 1.22 ServeMux)
go get github.com/go-chi/chi/v5@latest

# Suggested structure:
# cmd/
#   server/
#     main.go     — start server, graceful shutdown
# internal/
#   handler/      — HTTP handlers
#   store/        — data access layer
#   middleware/   — logging, auth middleware
# main.go
```

---

_Placeholder — start when you reach Phase 5._
