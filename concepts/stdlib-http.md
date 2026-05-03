# Standard Library — net/http

> **Phase 4 · Week 8**  
> Covers: `http.ListenAndServe`, `http.Handler`/`HandlerFunc`, `ServeMux`, middleware, `http.Client`, timeouts

---

## What It Is

_Coming in Phase 4. Stub created — fill after completing Phase 3._

Topics to cover:
- `http.Handler` interface: `ServeHTTP(ResponseWriter, *Request)`
- `http.HandlerFunc` — convert a function to a handler
- `http.ServeMux` (Go 1.22 enhanced routing with method and wildcard support)
- Middleware pattern: `func(http.Handler) http.Handler`
- `http.Client` — making outbound requests, timeouts, custom transport
- `http.Request` — reading URL params, query string, body, headers
- `http.ResponseWriter` — writing status, headers, body
- Graceful shutdown: `http.Server.Shutdown(ctx)`
- `net/http/httptest` — test servers without network

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

## Exercises

See [`exercises/05-http/`](../exercises/05-http/README.md) — Parts 5–6

---

## Related Files

- Previous: [`concepts/stdlib-io.md`](stdlib-io.md)
- Next: [`concepts/testing.md`](testing.md)
- Cheatsheet: [`cheatsheets/stdlib.md`](../cheatsheets/stdlib.md)
- Project: [`projects/rest-api/`](../projects/rest-api/)
