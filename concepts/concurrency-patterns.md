# Concurrency Patterns

> **Phase 3 · Week 6**  
> Covers: fan-in, fan-out, worker pools, pipeline, context-based cancellation, errgroup, common pitfalls

---

## What It Is

_Coming in Phase 3. Stub created — fill after completing Phase 2._

Topics to cover:
- **Pipeline** — chain stages with channels: `generate → transform → sink`
- **Fan-out** — spread work across multiple goroutines from one input channel
- **Fan-in** — merge multiple channels into one (merge pattern)
- **Worker pool** — fixed N goroutines draining a job channel
- **Context cancellation** — `context.WithCancel` to stop goroutines cleanly
- **errgroup** (`golang.org/x/sync/errgroup`) — run goroutines, collect first error
- **Goroutine leaks** — channels that are never closed, goroutines that never exit
- **Race detector** — `go test -race ./...`

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

See [`exercises/04-concurrency/`](../exercises/04-concurrency/README.md)

---

## Related Files

- Previous: [`concepts/sync-primitives.md`](sync-primitives.md)
- Next: [`concepts/context.md`](context.md)
- Cheatsheet: [`cheatsheets/concurrency.md`](../cheatsheets/concurrency.md)
