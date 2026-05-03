# Goroutines & Channels

> **Phase 3 · Week 5**  
> Covers: lightweight threads, `go` keyword, GMP scheduling model, unbuffered vs buffered channels, directional channels, range over channel, select

---

## What It Is

_Coming in Phase 3. Stub created — fill after completing Phase 2._

Topics to cover:
- `go fn()` — starts a goroutine (not a thread, not a coroutine)
- GMP model: Goroutines, OS Threads (M), Logical Processors (P)
- Channels: `ch := make(chan int)` — synchronised communication
- Unbuffered channels: sender blocks until receiver is ready (and vice versa)
- Buffered channels: `make(chan int, 5)` — sends don't block until buffer full
- Directional types: `chan<- T` (send only), `<-chan T` (receive only)
- `range` over a channel — reads until channel is closed
- `select` — multiplex channel operations (like `switch` but for channels)
- `close(ch)` — signal that no more values will be sent

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

- Previous: [`concepts/error-handling.md`](error-handling.md)
- Next: [`concepts/sync-primitives.md`](sync-primitives.md)
- Cheatsheet: [`cheatsheets/concurrency.md`](../cheatsheets/concurrency.md)
