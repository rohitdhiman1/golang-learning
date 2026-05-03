# Error Handling

> **Phase 2 · Week 4**  
> Covers: `error` interface, `errors.New`, `fmt.Errorf %w`, `errors.Is`/`errors.As`, sentinel errors, custom error types

---

## What It Is

In Go, **errors are values**, not exceptions. Functions return errors as a second return value, and callers check them explicitly. There is no `try/catch` — the `if err != nil` pattern IS Go's error handling.

The `error` type is just an interface:

```go
type error interface {
    Error() string
}
```

---

## Java/Python Comparison

| Concept | Java | Python | Go |
|---|---|---|---|
| Signal error | `throw new Exception("msg")` | `raise Exception("msg")` | `return fmt.Errorf("msg")` |
| Handle error | `try { ... } catch (E e) { ... }` | `try: ... except E: ...` | `if err != nil { ... }` |
| Error hierarchy | Class hierarchy (`IOException extends Exception`) | Class hierarchy (`OSError(Exception)`) | Wrapping chains (`fmt.Errorf("op: %w", err)`) |
| Check type | `catch (IOException e)` | `except IOError` | `errors.As(err, &target)` |
| Check identity | `catch (e)` + `e == sentinel` | `except SpecificError` | `errors.Is(err, sentinel)` |
| Unchecked errors | Runtime exceptions pass silently | Unhandled exceptions crash | Compiler warns if you ignore the error (with linters) |

**Key insight:** Go's error handling is more verbose but more explicit. You can't accidentally ignore an exception. Every error path is visible in the code.

---

## How It Works

### Creating errors

```go
// Simple error
err := errors.New("something went wrong")

// Formatted error
err := fmt.Errorf("failed to open %s: file not found", filename)
```

### Wrapping errors (the `%w` verb)

Wrapping adds context while preserving the original error:

```go
data, err := os.ReadFile(path)
if err != nil {
    return fmt.Errorf("loadConfig: %w", err)
}
```

Now callers can inspect the chain:
```go
// errors.Is checks anywhere in the chain
if errors.Is(err, os.ErrNotExist) {
    // file doesn't exist
}
```

### Sentinel errors

Package-level error values for known, expected conditions:

```go
var (
    ErrNotFound   = errors.New("not found")
    ErrForbidden  = errors.New("forbidden")
    ErrValidation = errors.New("validation failed")
)

func FindUser(id int) (*User, error) {
    if id <= 0 {
        return nil, ErrValidation
    }
    // ...
    return nil, ErrNotFound
}

// Caller:
if errors.Is(err, ErrNotFound) {
    // handle not found
}
```

### Custom error types

When you need to carry extra data with the error:

```go
type ValidationError struct {
    Field   string
    Message string
}

func (e *ValidationError) Error() string {
    return fmt.Sprintf("validation: %s — %s", e.Field, e.Message)
}

func Validate(u User) error {
    if u.Name == "" {
        return &ValidationError{Field: "Name", Message: "required"}
    }
    return nil
}

// Caller — extract the concrete error:
var ve *ValidationError
if errors.As(err, &ve) {
    fmt.Printf("field %s: %s\n", ve.Field, ve.Message)
}
```

### `errors.Is` vs `errors.As`

| Function | Purpose | Checks |
|---|---|---|
| `errors.Is(err, target)` | Is this specific error value in the chain? | Identity (sentinel errors) |
| `errors.As(err, &target)` | Is there an error of this type in the chain? | Type (custom error types) |

Both unwrap the error chain automatically.

### panic/recover (rare)

`panic` is for truly unrecoverable situations (programmer bugs, not user errors):

```go
// panic — crashes the program (unless recovered)
panic("impossible: negative array length")

// recover — catch a panic (only in deferred functions)
func safeCall(f func()) (err error) {
    defer func() {
        if r := recover(); r != nil {
            err = fmt.Errorf("panic: %v", r)
        }
    }()
    f()
    return nil
}
```

**Rule:** Don't use panic for normal error handling. Reserve it for invariant violations.

---

## Key Gotchas

1. **Don't use `_` to ignore errors** without a good reason:
   ```go
   data, _ := os.ReadFile("config.json") // BAD — silent failure
   ```

2. **`%w` vs `%v` in `fmt.Errorf`:**
   - `%w` wraps — preserves the chain, `errors.Is`/`errors.As` can find it
   - `%v` formats — creates a new error string, original error is lost

3. **error is an interface** — a nil `*MyError` is not a nil `error`:
   ```go
   func doThing() error {
       var err *MyError = nil
       return err // NOT nil! Same nil-interface trap as interfaces.md
   }
   ```
   Fix: `return nil` explicitly.

4. **Don't wrap with redundant context:**
   ```go
   // BAD: "open config.json: open config.json: no such file"
   return fmt.Errorf("open %s: %w", path, err)
   // when err is already "open config.json: no such file"
   ```

5. **Sentinel errors are public API** — once exported, you can't change them without breaking callers.

---

## Idiomatic Go

- **Handle errors immediately** — don't defer error checking. The `if err != nil { return }` pattern is idiomatic, not verbose.
- **Add context when wrapping** — describe what YOU were doing, not what the callee was doing:
  ```go
  return fmt.Errorf("save user %d: %w", id, err) // good — describes YOUR operation
  ```
- **Use sentinel errors for expected conditions** (`ErrNotFound`, `ErrTimeout`).
- **Use custom error types when callers need structured data** (field, HTTP status, etc.).
- **Don't log and return** — either log the error (and handle it) or return it. Don't do both.
- **Errors should flow up** — lower layers wrap, upper layers (main, HTTP handler) decide what to do.
- **panic is not for error handling** — use it for impossible states and programmer bugs.

---

## Exercises

See [`exercises/03-interfaces/`](../exercises/03-interfaces/README.md) — Part 5: Error Handling

---

## Related Files

- Previous: [`concepts/generics.md`](generics.md)
- Next: [`concepts/goroutines-channels.md`](goroutines-channels.md)
