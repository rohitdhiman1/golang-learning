# Cheatsheet — Testing

> Fill as you complete Phase 4.

---

## Quick Reference

_Coming in Phase 4._

Topics to capture here:
- Test file naming: `_test.go`
- Test function signature: `func TestXxx(t *testing.T)`
- `t.Error` vs `t.Fatal` vs `t.Log`
- Table-driven test template
- `t.Run` subtest pattern
- `t.Parallel()` — parallel subtests
- Benchmark function: `func BenchmarkXxx(b *testing.B)`
- `b.ResetTimer()`, `b.ReportAllocs()`
- `httptest.NewRecorder()` template
- `httptest.NewServer()` template
- Running commands:
  ```bash
  go test ./...
  go test -v ./...
  go test -run TestName ./...
  go test -bench=. ./...
  go test -benchmem ./...
  go test -cover ./...
  go test -race ./...
  ```

---

_Placeholder — fill at the end of Phase 4._
