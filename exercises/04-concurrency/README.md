# Exercises — 04: Concurrency

> **Phase 3 · Week 5–6**  
> Read [`concepts/goroutines-channels.md`](../../concepts/goroutines-channels.md), [`concepts/sync-primitives.md`](../../concepts/sync-primitives.md), [`concepts/concurrency-patterns.md`](../../concepts/concurrency-patterns.md), and [`concepts/context.md`](../../concepts/context.md) before starting.

---

## Run

```bash
cd exercises/04-concurrency
go run .
go run -race .    # run with race detector
go vet ./...
```

---

## Part 1 — Goroutines

_File: `part1_goroutines.go`_

| # | Exercise | Concepts |
|---|---|---|
| 1 | Basic goroutine | `go func(){}()`, goroutine launch |
| 2 | WaitGroup | `sync.WaitGroup`, `Add`, `Done`, `Wait` |
| 3 | Closure capture gotcha | Loop variable capture — wrong way vs right way |
| 4 | Concurrent printing | Non-deterministic ordering |
| 5 | Named function goroutine | Passing `*sync.WaitGroup` to a named function |

---

## Part 2 — Unbuffered Channels

_File: `part2_unbuffered_channels.go`_

| # | Exercise | Concepts |
|---|---|---|
| 1 | Basic send & receive | `make(chan T)`, `<-`, blocking semantics |
| 2 | Channel as synchronisation | Done signal pattern |
| 3 | Producer → consumer | `range` over channel, `close` |
| 4 | Directional channels | `chan<- T`, `<-chan T` in function signatures |
| 5 | Multiple senders | Many goroutines sending to one channel |

---

## Part 3 — Buffered Channels & Select

_File: `part3_buffered_channels_select.go`_

| # | Exercise | Concepts |
|---|---|---|
| 1 | Buffered channel | `make(chan T, cap)`, `len`, `cap` |
| 2 | Select statement | Wait on multiple channels |
| 3 | Non-blocking select | `select` with `default` |
| 4 | Timeout | `time.After` in select |
| 5 | Done channel pattern | Cancellable infinite generator |

---

## Part 4 — Sync Primitives

_File: `part4_sync_primitives.go`_

| # | Exercise | Concepts |
|---|---|---|
| 1 | Mutex — safe counter | `sync.Mutex`, `Lock`, `Unlock` |
| 2 | Mutex in struct | Embedded mutex pattern |
| 3 | RWMutex | `sync.RWMutex`, concurrent reads, exclusive writes |
| 4 | sync.Once | Lazy initialisation, run-exactly-once |
| 5 | Race detector | `go run -race .`, identifying data races |

---

## Part 5 — Concurrency Patterns

_File: `part5_concurrency_patterns.go`_

| # | Exercise | Concepts |
|---|---|---|
| 1 | Pipeline | `generate → double → consume` |
| 2 | Fan-out | Multiple workers reading from one channel |
| 3 | Fan-in (merge) | Combine multiple channels into one |
| 4 | Worker pool | Fixed N goroutines, job + results channels |
| 5 | Cancellable pipeline | Done channel to stop infinite generator |

---

## Part 6 — Context

_File: `part6_context.go`_

| # | Exercise | Concepts |
|---|---|---|
| 1 | WithCancel | Manual cancellation, `ctx.Done()`, `ctx.Err()` |
| 2 | WithTimeout | Automatic timeout, deadline exceeded |
| 3 | WithValue | Typed keys, request-scoped data |
| 4 | Propagation | Parent cancellation cascades to children |
| 5 | Graceful shutdown | Simulated server with context-based shutdown |
