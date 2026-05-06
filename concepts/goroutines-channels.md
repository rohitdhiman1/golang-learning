# Goroutines & Channels

> **Phase 3 · Week 5**  
> Covers: lightweight threads, `go` keyword, GMP scheduling model, unbuffered vs buffered channels, directional channels, range over channel, select

---

## What It Is

Go's concurrency model is built on two primitives: **goroutines** (lightweight threads) and **channels** (typed conduits for communication between goroutines).

### Goroutines

A goroutine is a function executing concurrently with other goroutines in the same address space. They are **not** OS threads — they're multiplexed onto a small number of OS threads by the Go runtime.

```go
go doWork()           // launch a goroutine — returns immediately

go func() {           // anonymous goroutine
    fmt.Println("hi")
}()
```

Key facts:
- **Cheap**: ~2 KB initial stack (grows/shrinks dynamically). You can run millions.
- **Cooperative scheduling**: goroutines yield at channel ops, system calls, `runtime.Gosched()`, function calls.
- **No return value**: `go f()` discards the return — use channels or shared state to get results back.
- **Lifetime**: a goroutine runs until its function returns. When `main()` returns, all goroutines are killed — no cleanup.

### Channels

Channels are Go's primary synchronisation and communication mechanism. They are typed, first-class values.

```go
ch := make(chan int)       // unbuffered channel of ints
ch := make(chan int, 5)    // buffered channel, capacity 5

ch <- 42                   // send
v := <-ch                  // receive
v, ok := <-ch              // receive with closed-check (ok == false if closed)
close(ch)                  // signal no more sends
```

**Unbuffered** (default): send blocks until a receiver is ready, receive blocks until a sender is ready. This gives you a synchronisation point — both goroutines rendezvous.

**Buffered**: send blocks only when the buffer is full. Receive blocks only when the buffer is empty. Useful for decoupling producer/consumer speeds.

### Directional Channels

Function signatures can restrict a channel to send-only or receive-only:

```go
func producer(out chan<- int) { out <- 1 }    // can only send
func consumer(in <-chan int)  { v := <-in }   // can only receive
```

The compiler enforces these — you get a compile error if you try to receive on a `chan<-` or send on a `<-chan`. Bidirectional `chan T` is implicitly convertible to either direction.

### Range Over a Channel

```go
for v := range ch {
    fmt.Println(v)   // loops until ch is closed
}
```

`range` blocks waiting for the next value and exits when the channel is closed. If nobody closes the channel, this is a goroutine leak.

### Select

`select` lets you wait on multiple channel operations simultaneously — like `switch` but for channels:

```go
select {
case v := <-ch1:
    fmt.Println("from ch1:", v)
case ch2 <- 42:
    fmt.Println("sent to ch2")
case <-time.After(1 * time.Second):
    fmt.Println("timeout")
default:
    fmt.Println("nothing ready — non-blocking")
}
```

- If multiple cases are ready, one is chosen **at random** (prevents starvation).
- `default` makes it non-blocking.
- `select{}` (empty select) blocks forever — useful in long-running services.

---

## Java/Python Comparison

| Concept | Go | Java | Python |
|---|---|---|---|
| Lightweight thread | `go f()` | `Thread` / virtual threads (Project Loom) | `threading.Thread` / `asyncio.create_task` |
| Communication | Channels (`chan T`) | `BlockingQueue`, `CompletableFuture` | `queue.Queue`, `asyncio.Queue` |
| Thread cost | ~2 KB, millions possible | ~512 KB–1 MB per thread (pre-Loom) | GIL limits true parallelism (threads) |
| Synchronisation | Channel ops are sync points | `synchronized`, `Lock`, `Semaphore` | `Lock`, `Event`, `Condition` |
| Select/multiplex | `select` statement | `CompletableFuture.anyOf()` | `asyncio.wait(..., FIRST_COMPLETED)` |
| CSP model | Built-in | Not native (need libraries) | Not native |

**Key difference**: Go follows the **CSP** (Communicating Sequential Processes) model — "Don't communicate by sharing memory; share memory by communicating." Java and Python default to shared-memory + locks.

---

## How It Works — The GMP Model

Go's runtime scheduler uses three entities:

- **G** (Goroutine): the goroutine itself — its stack, instruction pointer, and state.
- **M** (Machine): an OS thread. The runtime creates these as needed (up to `GOMAXPROCS` running at once).
- **P** (Processor): a logical processor with a local run queue. `GOMAXPROCS` (default = number of CPU cores) determines how many Ps exist.

```
     P0            P1
  [G G G G]    [G G G G]     ← local run queues
     |             |
     M0            M1         ← OS threads
     |             |
   CPU core 0   CPU core 1
```

Scheduling flow:
1. Each P has a local queue of Gs (goroutines ready to run).
2. An M picks a P, dequeues a G, and runs it.
3. When a G blocks (channel op, syscall, I/O), the M may park or the P steals from another P's queue (**work stealing**).
4. When a G is created with `go`, it's placed on the current P's local queue.
5. There's also a global run queue used when local queues are empty.

**Why this matters**: goroutines are cheap because context-switching between them happens in user space (no kernel transition), and the scheduler is aware of Go-specific blocking (channels, select).

---

## Key Gotchas

### 1. Main exits = all goroutines die
```go
func main() {
    go fmt.Println("hello")   // might never print
}
// main() returns → process exits → goroutine killed
```
Fix: use `sync.WaitGroup`, channels, or `select{}`.

### 2. Closure variable capture in loops
```go
for i := 0; i < 5; i++ {
    go func() {
        fmt.Println(i)   // BUG: all goroutines see i=5
    }()
}
```
Fix: pass `i` as a parameter, or use `i := i` to shadow:
```go
for i := 0; i < 5; i++ {
    go func(n int) {
        fmt.Println(n)   // correct: each goroutine gets its own copy
    }(i)
}
```

### 3. Sending on a closed channel panics
```go
close(ch)
ch <- 1   // panic: send on closed channel
```
Only the **sender** should close a channel. Never close from the receiver side.

### 4. Goroutine leaks
A goroutine that blocks forever (waiting on a channel nobody sends to / receives from) is a leak. It will never be garbage collected.

### 5. No guaranteed execution order
Goroutines are scheduled cooperatively — you cannot assume any ordering without explicit synchronisation.

---

## Idiomatic Go

1. **"Do not communicate by sharing memory; share memory by communicating."** — prefer channels over mutexes when the problem is about passing data/ownership.
2. **Start goroutines with clear ownership** — whoever starts a goroutine should know when and how it will stop.
3. **Close channels from the sender side** — the sender knows when it's done. Receivers just range or check `ok`.
4. **Use directional channel types in function signatures** — `chan<-` / `<-chan` communicates intent and lets the compiler help.
5. **Keep goroutine lifetimes short and obvious** — long-lived goroutines should have a cancellation mechanism (context, done channel).
6. **Never start a goroutine without knowing when it will stop.**

---

## Exercises

See [`exercises/04-concurrency/`](../exercises/04-concurrency/README.md) — Parts 1–3

---

## Related Files

- Previous: [`concepts/error-handling.md`](error-handling.md)
- Next: [`concepts/sync-primitives.md`](sync-primitives.md)
- Cheatsheet: [`cheatsheets/concurrency.md`](../cheatsheets/concurrency.md)
