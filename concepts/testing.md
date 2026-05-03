# Testing

> **Phase 4 · Week 8**  
> Covers: `testing` package, table-driven tests, `t.Run` subtests, benchmarks, `httptest`, test fixtures

---

## What It Is

_Coming in Phase 4. Stub created — fill after completing Phase 3._

Topics to cover:
- `func TestXxx(t *testing.T)` — test function naming convention
- `t.Error` vs `t.Fatal` — continue vs stop test on failure
- Table-driven tests — the idiomatic Go test pattern
- `t.Run(name, func)` — named subtests, parallel with `t.Parallel()`
- `testing.B` — benchmark functions: `b.N` loop, `b.ResetTimer`
- `httptest.NewRecorder()`, `httptest.NewServer()` — testing HTTP handlers
- Test helpers: `t.Helper()`, `testify/assert` (popular third-party)
- Golden files — for testing output that's complex to inline
- Running: `go test ./...`, `-v`, `-run TestName`, `-bench=.`, `-cover`

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

See [`exercises/06-testing/`](../exercises/06-testing/README.md)

---

## Related Files

- Cheatsheet: [`cheatsheets/testing.md`](../cheatsheets/testing.md)
