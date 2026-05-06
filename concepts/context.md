# Context

> **Phase 3–4 · Week 6/8**  
> Covers: `context.Background()`, `WithCancel`, `WithTimeout`, `WithValue`, propagation rules

---

## What It Is

`context.Context` is Go's mechanism for carrying deadlines, cancellation signals, and request-scoped values across API boundaries and goroutines.

Every long-running or I/O-bound function in production Go accepts `ctx context.Context` as its first parameter. This is one of Go's strongest conventions.

### Creating Contexts

```go
// Root contexts — only use at the top level (main, init, test setup)
ctx := context.Background()   // the "empty" context — never cancelled
ctx := context.TODO()         // placeholder when you haven't decided which context to use
```

### Derived Contexts

All derived contexts form a **tree**. Cancelling a parent cancels all its children.

```go
// WithCancel — manual cancellation
ctx, cancel := context.WithCancel(parent)
defer cancel()   // ALWAYS defer cancel to release resources

// WithTimeout — auto-cancel after duration
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()

// WithDeadline — auto-cancel at a specific time
ctx, cancel := context.WithDeadline(parent, time.Now().Add(5*time.Second))
defer cancel()

// WithValue — attach request-scoped values (use sparingly)
ctx := context.WithValue(parent, "requestID", "abc-123")
```

### Checking for Cancellation

```go
select {
case <-ctx.Done():
    return ctx.Err()   // context.Canceled or context.DeadlineExceeded
case result := <-doWork():
    return result
}
```

Or for simple checks:
```go
if ctx.Err() != nil {
    return ctx.Err()
}
```

### Reading Values

```go
reqID := ctx.Value("requestID")   // returns any — type-assert as needed
if id, ok := reqID.(string); ok {
    log.Println("request:", id)
}
```

---

## Java/Python Comparison

| Concept | Go | Java | Python |
|---|---|---|---|
| Cancellation propagation | `context.Context` | `Future.cancel()`, interrupt flags | `asyncio.Task.cancel()` |
| Timeouts | `context.WithTimeout` | `Future.get(timeout, unit)` | `asyncio.wait_for(coro, timeout)` |
| Request-scoped data | `context.WithValue` | `ThreadLocal` / MDC | `contextvars.ContextVar` (3.7+) |
| Convention | First param: `ctx context.Context` | No universal convention | No universal convention |

**Key difference**: Go's context is pervasive — it's threaded through every function in the call chain. Java and Python tend to use thread-local storage or async-local equivalents, which are more implicit.

---

## How It Works

### The Context Tree

```
Background()
├── WithCancel()           ← cancel() called → cancels this and all children
│   ├── WithTimeout(3s)    ← auto-cancels after 3s or when parent cancels
│   │   └── WithValue()    ← carries data, inherits parent's deadline
│   └── WithCancel()       ← independent cancel
└── WithTimeout(10s)       ← separate subtree
```

- `ctx.Done()` returns a channel that closes when the context is cancelled.
- `ctx.Err()` returns `nil` while active, then `context.Canceled` or `context.DeadlineExceeded`.
- `ctx.Deadline()` returns the deadline (if set) and whether one exists.
- `ctx.Value(key)` walks up the tree to find the nearest matching key.

### Propagation Rules

1. Cancelling a parent cancels all children (recursively).
2. Cancelling a child does **not** affect the parent or siblings.
3. A derived context's deadline cannot exceed its parent's — the shorter one wins.
4. Values are inherited: a child can see its ancestors' values, but not its siblings'.

### The `cancel()` Function

**Always call it.** Even if the context will be cancelled by a timeout, calling `cancel()` releases resources immediately instead of waiting for the timer.

```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
defer cancel()   // releases timer resources even if work finishes in 1s
```

---

## Key Gotchas

### 1. Not calling cancel()
```go
ctx, cancel := context.WithTimeout(parent, 5*time.Second)
// forgot defer cancel()
// → timer goroutine leaks until timeout expires
```
**Always** `defer cancel()` on the line immediately after creating the context.

### 2. Using WithValue as a grab bag
```go
ctx = context.WithValue(ctx, "db", database)
ctx = context.WithValue(ctx, "logger", log)
ctx = context.WithValue(ctx, "config", cfg)
// This is an anti-pattern — use proper dependency injection
```
`WithValue` is for **request-scoped** data that transits process boundaries: request IDs, trace IDs, auth tokens. Not for dependency injection.

### 3. String keys collide
```go
// Package A and Package B both use key "id" → collision
ctx = context.WithValue(ctx, "id", userID)
ctx = context.WithValue(ctx, "id", traceID)   // overwrites!
```
Fix: use unexported typed keys:
```go
type contextKey string
const requestIDKey contextKey = "requestID"
ctx = context.WithValue(ctx, requestIDKey, "abc-123")
```

### 4. Not checking ctx.Done() in loops
```go
for _, item := range bigSlice {
    process(item)   // if context was cancelled, this keeps running uselessly
}
```
Fix: check `ctx.Err()` periodically:
```go
for _, item := range bigSlice {
    if ctx.Err() != nil {
        return ctx.Err()
    }
    process(item)
}
```

---

## Idiomatic Go

1. **`ctx context.Context` is always the first parameter** — never put it in a struct.
   ```go
   func FetchUser(ctx context.Context, id int) (*User, error)
   ```
2. **`defer cancel()` immediately** — on the very next line after `WithCancel`/`WithTimeout`.
3. **Pass context down, never up** — context flows from caller to callee, through the call chain.
4. **Don't store contexts in structs** — pass them as function parameters. Storing in a struct makes the lifetime unclear.
5. **Use `context.Background()` in `main()` and tests** — `context.TODO()` is for code that should eventually get a real context.
6. **Use typed keys for `WithValue`** — prevents collisions between packages.
7. **Keep `WithValue` minimal** — if you're putting more than a request ID and trace ID, reconsider.

---

## Exercises

See [`exercises/04-concurrency/`](../exercises/04-concurrency/README.md) — Part 6

---

## Related Files

- Previous: [`concepts/concurrency-patterns.md`](concurrency-patterns.md)
- Next: [`concepts/stdlib-io.md`](stdlib-io.md)
