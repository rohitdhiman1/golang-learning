# Error Handling

> **Phase 2 · Week 4**  
> Covers: `error` interface, `errors.New`, `fmt.Errorf %w`, `errors.Is`/`errors.As`, sentinel errors, custom error types

---

## What It Is

_Coming in Phase 2. Stub created — fill after completing Phase 1._

Topics to cover:
- `error` is just an interface: `type error interface { Error() string }`
- Creating errors: `errors.New`, `fmt.Errorf`
- Wrapping errors for context: `fmt.Errorf("op: %w", err)`
- Unwrapping: `errors.Is(err, target)`, `errors.As(err, &target)`
- Sentinel errors: `var ErrNotFound = errors.New("not found")`
- Custom error types: implement the `error` interface
- Why no exceptions — errors as values = explicit, auditable, composable
- Error handling patterns: early return, wrapping chains, panic/recover (rare)

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

See [`exercises/03-interfaces/`](../exercises/03-interfaces/README.md)

---

## Related Files

- Previous: [`concepts/generics.md`](generics.md)
- Next: [`concepts/goroutines-channels.md`](goroutines-channels.md)
