# Sync Primitives

> **Phase 3 · Week 5**  
> Covers: `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `sync.Once`, `sync.Map`

---

## What It Is

_Coming in Phase 3. Stub created — fill after completing Phase 2._

Topics to cover:
- `sync.Mutex` — mutual exclusion lock for shared state
- `sync.RWMutex` — read/write lock (multiple concurrent readers OR one writer)
- `sync.WaitGroup` — wait for a set of goroutines to finish
- `sync.Once` — run a function exactly once (lazy init, singletons)
- `sync.Map` — concurrent map (prefer regular map + Mutex unless profiling shows contention)
- Go memory model basics: happens-before guarantees
- When to use channels vs sync primitives ("share memory by communicating")

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

- Previous: [`concepts/goroutines-channels.md`](goroutines-channels.md)
- Next: [`concepts/concurrency-patterns.md`](concurrency-patterns.md)
