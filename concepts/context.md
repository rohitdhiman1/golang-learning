# Context

> **Phase 3–4 · Week 6/8**  
> Covers: `context.Background()`, `WithCancel`, `WithTimeout`, `WithValue`, propagation rules

---

## What It Is

_Coming in Phase 3/4. Stub created — fill after completing goroutines and channels._

Topics to cover:
- `context.Context` — carries deadlines, cancellations, and request-scoped values
- `context.Background()` — root context (use in `main`, tests, top-level server handlers)
- `context.TODO()` — placeholder when you haven't decided yet
- `context.WithCancel(parent)` → `(ctx, cancel)` — cancel subtree
- `context.WithTimeout(parent, duration)` — auto-cancel after duration
- `context.WithDeadline(parent, time)` — cancel at a specific instant
- `context.WithValue(parent, key, val)` — attach request-scoped values (use sparingly)
- Always call `cancel()` — defer it immediately after `WithCancel`/`WithTimeout`
- Propagation: child contexts are cancelled when parent is cancelled
- The `ctx.Done()` channel — close when context is cancelled
- `ctx.Err()` — `context.Canceled` or `context.DeadlineExceeded`

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

See [`exercises/04-concurrency/`](../exercises/04-concurrency/README.md) — Part 6

---

## Related Files

- Previous: [`concepts/concurrency-patterns.md`](concurrency-patterns.md)
- Next: [`concepts/stdlib-io.md`](stdlib-io.md)
