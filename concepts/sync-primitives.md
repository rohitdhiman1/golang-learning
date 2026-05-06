# Sync Primitives

> **Phase 3 · Week 5**  
> Covers: `sync.Mutex`, `sync.RWMutex`, `sync.WaitGroup`, `sync.Once`, `sync.Map`

---

## What It Is

The `sync` package provides low-level synchronisation primitives for cases where channels are overkill — typically protecting shared state accessed by multiple goroutines.

### sync.Mutex — Mutual Exclusion Lock

```go
var mu sync.Mutex
var count int

func increment() {
    mu.Lock()
    count++
    mu.Unlock()
}
```

- Only one goroutine can hold the lock at a time.
- **Always unlock** — use `defer mu.Unlock()` right after `Lock()` to prevent forgetting.
- Zero value is an unlocked mutex — no need to initialise.

### sync.RWMutex — Read/Write Lock

```go
var rw sync.RWMutex
var data map[string]string

func read(key string) string {
    rw.RLock()
    defer rw.RUnlock()
    return data[key]
}

func write(key, val string) {
    rw.Lock()
    defer rw.Unlock()
    data[key] = val
}
```

- Multiple goroutines can hold `RLock()` simultaneously (concurrent reads).
- `Lock()` is exclusive — blocks until all readers release and no other writer holds it.
- Use when reads vastly outnumber writes.

### sync.WaitGroup — Wait for Goroutines

```go
var wg sync.WaitGroup

for i := 0; i < 5; i++ {
    wg.Add(1)
    go func(n int) {
        defer wg.Done()
        doWork(n)
    }(i)
}

wg.Wait()   // blocks until counter reaches 0
```

- `Add(n)` increments the counter, `Done()` decrements it (equivalent to `Add(-1)`).
- **Call `Add()` before the `go` statement**, not inside the goroutine — otherwise `Wait()` might return early.
- Zero value is ready to use.

### sync.Once — Run Exactly Once

```go
var once sync.Once
var db *sql.DB

func getDB() *sql.DB {
    once.Do(func() {
        db = connectToDatabase()   // runs only on first call
    })
    return db
}
```

- Thread-safe lazy initialisation.
- Even if 100 goroutines call `getDB()` simultaneously, `connectToDatabase()` runs exactly once.
- If the function passed to `Do` panics, `Once` still considers it "done" — it won't retry.

### sync.Map — Concurrent Map

```go
var m sync.Map

m.Store("key", "value")
v, ok := m.Load("key")
m.Delete("key")
m.Range(func(k, v any) bool {
    fmt.Println(k, v)
    return true   // continue iteration
})
```

- Type-unsafe (`any` keys and values) — no generics support yet.
- **Don't reach for this by default.** A regular `map` + `sync.Mutex` or `sync.RWMutex` is usually clearer and equally fast.
- Optimised for two patterns: (1) keys written once, read many times; (2) disjoint sets of keys across goroutines.

---

## Java/Python Comparison

| Concept | Go | Java | Python |
|---|---|---|---|
| Mutual exclusion | `sync.Mutex` | `synchronized` / `ReentrantLock` | `threading.Lock` |
| Read/write lock | `sync.RWMutex` | `ReadWriteLock` | `threading.RLock` (reentrant, not R/W) |
| Wait for threads | `sync.WaitGroup` | `CountDownLatch` / `Thread.join()` | `threading.Thread.join()` |
| Run once | `sync.Once` | `static {}` block / double-checked locking | `threading.Event` / module-level init |
| Concurrent map | `sync.Map` | `ConcurrentHashMap` | No direct equivalent (dict + Lock) |
| Atomic ops | `sync/atomic` | `AtomicInteger`, `AtomicReference` | No equivalent (GIL helps, but not for multiprocessing) |

**Key difference**: Go mutexes are **not reentrant** — if a goroutine that already holds a lock tries to lock it again, it deadlocks. Java's `synchronized` and `ReentrantLock` allow re-entry.

---

## How It Works

### Go Memory Model

The Go memory model defines when a read in one goroutine is guaranteed to see a write from another. Without synchronisation, there are **no guarantees** — the compiler and CPU can reorder operations.

**Happens-before** guarantees (simplified):
- `mu.Unlock()` happens-before a subsequent `mu.Lock()` on the same mutex.
- `wg.Done()` happens-before `wg.Wait()` returns.
- `ch <- v` happens-before `v = <-ch` completes (for unbuffered channels).
- A send on a buffered channel happens-before the corresponding receive completes.
- `close(ch)` happens-before a receive that returns the zero value due to the close.

### When to Use Channels vs Sync Primitives

| Use Channels When | Use sync When |
|---|---|
| Passing ownership of data | Protecting access to shared state |
| Distributing work | Incrementing a counter |
| Signalling events (done, cancel) | Caching / lazy init (`sync.Once`) |
| You need select/timeout | Simple mutual exclusion |

Rule of thumb: **channels for coordination, mutexes for state protection.**

---

## Key Gotchas

### 1. Forgetting to unlock
```go
mu.Lock()
if err != nil {
    return err   // BUG: mutex never unlocked
}
mu.Unlock()
```
Fix: `defer mu.Unlock()` immediately after `Lock()`.

### 2. Copying a Mutex
```go
type SafeCounter struct {
    mu sync.Mutex
    n  int
}

c1 := SafeCounter{}
c2 := c1   // BUG: copies the mutex — both c1 and c2 share broken state
```
`go vet` catches this. Pass sync types by pointer, never by value.

### 3. WaitGroup.Add() inside the goroutine
```go
for i := 0; i < 5; i++ {
    go func() {
        wg.Add(1)   // BUG: race — Wait() might return before Add()
        defer wg.Done()
    }()
}
wg.Wait()
```
Fix: call `wg.Add(1)` **before** the `go` statement.

### 4. Lock ordering deadlocks
If goroutine A locks `mu1` then `mu2`, and goroutine B locks `mu2` then `mu1` — deadlock. Always acquire locks in a consistent order.

---

## Idiomatic Go

1. **Protect the data, not the code** — associate the mutex with the data it guards. Embed it in the same struct.
   ```go
   type SafeMap struct {
       mu   sync.Mutex
       data map[string]int
   }
   ```
2. **`defer mu.Unlock()` immediately** — prevents unlock-forgetting bugs.
3. **Prefer channels for coordination** — use sync primitives only for shared state that doesn't naturally flow through channels.
4. **Use `go vet`** — it detects copied locks, incorrect WaitGroup usage, and more.
5. **Keep critical sections short** — hold locks for the minimum time necessary.

---

## Exercises

See [`exercises/04-concurrency/`](../exercises/04-concurrency/README.md) — Part 4

---

## Related Files

- Previous: [`concepts/goroutines-channels.md`](goroutines-channels.md)
- Next: [`concepts/concurrency-patterns.md`](concurrency-patterns.md)
