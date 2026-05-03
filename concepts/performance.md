# Performance & Advanced Topics

> **Phase 6 · Week 11–12**  
> Covers: reflection, `pprof`, benchmarks, escape analysis, `sync.Pool`, build & deploy, interview internals

---

## What It Is

_Coming in Phase 6. Stub created — fill after completing Phase 5._

Topics to cover:
- **Reflection** — `reflect` package: `reflect.TypeOf`, `reflect.ValueOf`, struct tag parsing
  - Rule of thumb: if you're using reflection, question whether generics or interfaces solve it first
- **`pprof`** — CPU profiling, heap profiling, goroutine dumps, `go tool pprof`
- **Benchmarks** — `testing.B`, `-benchmem`, identifying allocations
- **Escape analysis** — `go build -gcflags="-m"` — does this value live on stack or heap?
- **`sync.Pool`** — reuse temporary objects to reduce GC pressure
- **Cross-compilation** — `GOOS=linux GOARCH=amd64 go build`
- **Docker multi-stage builds** — scratch image for static Go binaries
- **`-ldflags`** — inject version info at build time: `-ldflags="-X main.version=v1.2.3"`
- **GC internals** — tri-color mark & sweep, GOGC, GOMEMLIMIT
- **Slice internals** — backing array, header (ptr+len+cap), append doubling
- **Map internals** — hash buckets, overflow, load factor
- **Interface internals** — (itab, data) pair, nil interface vs interface holding nil pointer
- **Goroutine scheduling** — work stealing, preemption, GOMAXPROCS

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

## Related Files

- Interview prep: [`interview/questions.md`](../interview/questions.md)
- Previous: [`concepts/project-structure.md`](project-structure.md)
