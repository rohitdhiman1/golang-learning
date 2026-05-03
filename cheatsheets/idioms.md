# Cheatsheet — Go Idioms & Proverbs

> Fill as you progress through all phases. The "Go way" of thinking.

---

## Go Proverbs (Rob Pike)

> "Don't communicate by sharing memory; share memory by communicating."  
> "Concurrency is not parallelism."  
> "The bigger the interface, the weaker the abstraction."  
> "Make the zero value useful."  
> "A little copying is better than a little dependency."  
> "Clear is better than clever."  
> "Errors are values."  
> "Don't just check errors, handle them gracefully."  
> "Design the architecture, name the components, document the details."

---

## Naming Conventions

| Rule | Example | Anti-pattern |
|---|---|---|
| `MixedCaps` not `snake_case` | `userAge` | `user_age` |
| Short vars in small scopes | `i`, `v`, `k`, `err`, `ok` | `index`, `value`, `keyName` |
| Acronyms all-caps | `URL`, `HTTP`, `ID`, `ServeHTTP` | `Url`, `Http`, `Id`, `ServeHttp` |
| Single-method interfaces: `-er` | `Reader`, `Stringer`, `Handler` | `IReadable`, `StringInterface` |
| Package names: short, lowercase | `http`, `json`, `io` | `httpUtils`, `json_helper` |
| Error variable always `err` | `err` | `e`, `error`, `myErr` |
| Constructor: `NewXxx` | `NewServer() *Server` | `CreateServer()`, `MakeServer()` |
| Don't repeat package name | `http.Handler` | `http.HTTPHandler` |
| Receiver: short, consistent | `(s *Server)` | `(this *Server)`, `(self *Server)` |

---

## Common Idioms

```go
// Comma-ok pattern (map lookup, type assertion, channel receive)
v, ok := m[key]
v, ok := x.(Type)
v, ok := <-ch

// Early return (guard clauses — avoid deep nesting)
if err != nil {
    return fmt.Errorf("op: %w", err)
}

// Defer for cleanup
defer f.Close()
defer mu.Unlock()

// Constructor returns pointer when methods use pointer receivers
func NewServer(addr string) *Server {
    return &Server{Addr: addr}
}

// Accept interfaces, return structs
func Process(r io.Reader) (*Result, error) { ... }

// Table-driven tests
tests := []struct {
    name string
    in   int
    want int
}{
    {"positive", 5, 25},
    {"zero", 0, 0},
}
for _, tt := range tests {
    t.Run(tt.name, func(t *testing.T) {
        got := Square(tt.in)
        if got != tt.want {
            t.Errorf("Square(%d) = %d, want %d", tt.in, got, tt.want)
        }
    })
}
```

---

## Struct & Method Idioms (Phase 2)

```go
// Pointer receiver — use when method mutates OR struct is large
func (s *Server) Close() error { ... }

// Value receiver — use for small, read-only types
func (p Point) Distance(q Point) float64 { ... }

// Be consistent: if one method is pointer receiver, make all pointer receiver

// Embedding for composition (not inheritance)
type Server struct {
    http.Handler  // embedded — promotes ServeHTTP
    addr string
}

// Zero value should be useful
var mu sync.Mutex  // ready to use, no NewMutex() needed
var buf bytes.Buffer  // ready to use
```

---

## Interface Idioms (Phase 2)

```go
// Small interfaces (1-3 methods)
type Reader interface { Read([]byte) (int, error) }

// Define interfaces at the consumer, not the implementor

// The nil interface trap — always return nil explicitly
func doThing() error {
    var err *MyError = nil
    if err == nil {
        return nil  // NOT: return err (would be non-nil interface)
    }
    return err
}

// Type assertion (comma-ok for safety)
if s, ok := v.(string); ok { ... }

// Type switch for multiple types
switch x := v.(type) {
case int:    ...
case string: ...
}
```

---

## Error Handling Idioms (Phase 2)

```go
// Wrap with context (what YOU were doing)
return fmt.Errorf("save user %d: %w", id, err)

// Sentinel errors for expected conditions
var ErrNotFound = errors.New("not found")

// Check sentinels with errors.Is (unwraps chain)
if errors.Is(err, ErrNotFound) { ... }

// Extract custom errors with errors.As
var ve *ValidationError
if errors.As(err, &ve) { ... }

// Don't log AND return — do one or the other

// %w wraps (preserves chain), %v formats (loses chain)
```

---

## Generics Idioms (Phase 2)

```go
// Concrete types first — add generics when 3+ callers differ only in type

// Good generic use: data structures, utility functions
func Filter[T any](s []T, pred func(T) bool) []T { ... }

// Use stdlib packages (Go 1.21+): slices.Sort, slices.Contains, maps.Keys

// Keep constraints narrow: any < comparable < cmp.Ordered < custom

// Zero value of generic type
var zero T
```

---

_Updated through Phase 2._
