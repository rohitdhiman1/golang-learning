# Project Structure & Production Patterns

> **Phase 5 · Week 9–10**  
> Covers: `cmd/`/`internal/`/`pkg/` layout, dependency management, configuration, logging, database, CLI, graceful shutdown

---

## What It Is

_Coming in Phase 5. Stub created — fill after completing Phase 4._

Topics to cover:
- Standard Go project layout: `cmd/`, `internal/`, `pkg/`
- When to split packages — avoid premature abstraction
- Circular import prevention — `internal/` restricts imports
- `go.mod`, `go.sum`, `go get`, `go mod tidy`, private modules
- Configuration: `os.Getenv`, `flag` package, config structs, 12-factor app
- `log/slog` (Go 1.21+) — structured logging, levels, JSON handler
- `database/sql` — connection pooling, prepared statements, `sql.Null*` types
- `cobra` — CLI subcommands, flags, persistent flags, PersistentPreRunE hooks
- REST API: `chi` or Go 1.22 `ServeMux`, middleware chains, graceful shutdown

---

## Java/Python Comparison

_Fill when reading._

---

## How It Works

_Fill when reading._

---

## Key Gotchas

_Log gotchas here as you encounter them._

---

## Idiomatic Go

_Fill when reading._

---

## Project

See [`projects/cli-tool/`](../projects/cli-tool/) and [`projects/rest-api/`](../projects/rest-api/)

---

## Related Files

- Next: [`concepts/performance.md`](performance.md)
