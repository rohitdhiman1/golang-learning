# Functions

> **Phase 1 · Week 1**  
> Covers: multiple returns, named returns, variadic functions, first-class functions, closures, `init()`

---

## What It Is

Functions in Go are first-class values — they can be assigned to variables, passed as arguments, and returned from other functions. Go extends the familiar function model with multiple return values (the primary error-handling mechanism), named returns, and closures. There are no method overloading or default parameters.

---

## Java/Python Comparison

| Concept | Java | Python | Go |
|---|---|---|---|
| Multiple return | Wrap in object/tuple | `return a, b` | `return a, b` (natively supported) |
| Named returns | No | No | Yes (rare, but exists) |
| Variadic args | `String... args` | `*args` | `...T` |
| First-class funcs | Lambdas (Java 8+) | Yes | Yes |
| Closures | Lambda + effectively final | Yes | Yes |
| Default params | No (overloads instead) | Yes | No (use functional options pattern) |
| Overloading | Yes | No | No |
| Exceptions | `throws` / `try-catch` | `raise` / `try-except` | Return `error` as second value |
| Anonymous func | Lambda `() -> {}` | `lambda` | `func() {}` |

---

## How It Works

### Basic Functions

```go
// Signature: func name(params) returnType
func add(a int, b int) int {
    return a + b
}

// If consecutive params share a type, you can group them
func add(a, b int) int {
    return a + b
}

// No return value
func greet(name string) {
    fmt.Println("Hello,", name)
}

// Call
result := add(3, 4)   // 7
```

### Multiple Return Values

This is Go's primary error-handling mechanism — instead of exceptions.

```go
// Return multiple values
func divide(a, b float64) (float64, error) {
    if b == 0 {
        return 0, errors.New("division by zero")
    }
    return a / b, nil
}

// Must handle both return values (or discard with _)
result, err := divide(10, 3)
if err != nil {
    log.Fatal(err)
}
fmt.Println(result)   // 3.333...

// Discard a value with _
result, _ = divide(10, 3)

// Swap two values elegantly
a, b := 1, 2
a, b = b, a   // no temp variable
```

**From Java/Python:** You're trading `try/catch` blocks for explicit error checks after every call. It feels verbose at first, but it makes error paths visible and deliberate.

### Named Return Values

```go
// Name the return values — they're pre-declared as zero values
func minMax(arr []int) (min, max int) {
    min, max = arr[0], arr[0]
    for _, v := range arr[1:] {
        if v < min { min = v }
        if v > max { max = v }
    }
    return   // "naked return" — returns min and max
}

// ⚠️ Use sparingly — only when it genuinely aids readability
// Naked returns in long functions hide what's being returned
```

### Variadic Functions

```go
// ...T accepts zero or more arguments of type T
func sum(nums ...int) int {
    total := 0
    for _, n := range nums {
        total += n
    }
    return total
}

sum(1, 2, 3)         // 6
sum()                 // 0
nums := []int{1,2,3}
sum(nums...)          // spread a slice — equivalent to sum(1, 2, 3)

// fmt.Println is variadic: func Println(a ...any) (n int, err error)
fmt.Println("a", "b", "c")
```

### First-Class Functions

```go
// Assign a function to a variable
double := func(x int) int {
    return x * 2
}
fmt.Println(double(5))   // 10

// Function type
type MathFunc func(int, int) int

func apply(a, b int, fn MathFunc) int {
    return fn(a, b)
}

add := func(a, b int) int { return a + b }
fmt.Println(apply(3, 4, add))   // 7

// Pass a function literal inline
fmt.Println(apply(3, 4, func(a, b int) int { return a * b }))  // 12

// Return a function from a function
func multiplier(factor int) func(int) int {
    return func(x int) int {
        return x * factor
    }
}

triple := multiplier(3)
fmt.Println(triple(5))   // 15
```

### Closures

A closure is a function that "closes over" variables from its surrounding scope. Those variables persist as long as the closure exists.

```go
// counter returns a closure that maintains its own count
func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

c1 := counter()
c2 := counter()   // independent counter

fmt.Println(c1())  // 1
fmt.Println(c1())  // 2
fmt.Println(c2())  // 1  ← independent
fmt.Println(c1())  // 3

// Closures capture by reference — common gotcha in loops
funcs := make([]func(), 3)
for i := 0; i < 3; i++ {
    i := i   // shadow i — create a new variable per iteration
    funcs[i] = func() { fmt.Println(i) }
}
funcs[0]()   // 0
funcs[1]()   // 1
funcs[2]()   // 2

// Without the shadow (i := i), all three would print 3
```

**From Python:** Python closures also capture by reference. The loop variable gotcha exists in both languages.

### Defer (with functions)

```go
// defer runs AFTER the function returns, but BEFORE the caller resumes
// Deferred calls are stacked — LIFO order

func cleanup() {
    fmt.Println("clean up")
}

func main() {
    defer cleanup()
    fmt.Println("working...")
}
// Output:
// working...
// clean up

// Defer with arguments — arguments are evaluated immediately
x := 10
defer fmt.Println(x)   // prints 10, even if x changes later
x = 20
// Function body sees x=20, but deferred fmt.Println captured x=10
```

### `init()` Function

```go
// init() runs automatically before main(), after package-level vars are initialized
// A package can have multiple init() functions
// You cannot call init() manually

var config string

func init() {
    config = os.Getenv("APP_ENV")
    if config == "" {
        config = "development"
    }
}

func main() {
    fmt.Println("Config:", config)
}

// Ordering: package-level vars → init() → main()
// Import ordering determines which package's init() runs first
```

---

## Key Gotchas

### 1. You must use all return values (or discard with `_`)
```go
result, err := divide(10, 3)
// If you only write:
result := divide(10, 3)   // ERROR: multiple-value divide() used in single-value context
```

### 2. Closure loop variable capture
```go
// WRONG — all closures share the same loop variable
for i := 0; i < 3; i++ {
    go func() { fmt.Println(i) }()   // all may print 3
}

// CORRECT — pass as argument or shadow
for i := 0; i < 3; i++ {
    go func(n int) { fmt.Println(n) }(i)   // each goroutine gets its own copy
}
```

### 3. Named returns + naked return in long functions
```go
// Named returns are fine for short functions, but naked return in a 50-line
// function makes it hard to know what's being returned — be explicit
func longFunction() (result int, err error) {
    // ... lots of code
    return result, err   // explicit is better in long functions
}
```

### 4. No function overloading — use different names or variadic
```go
// INVALID: two functions with same name different params
func process(x int) {}
func process(x string) {}   // compile error

// Go way: different names or use interface{} / generics
func processInt(x int) {}
func processString(x string) {}
```

### 5. Defer runs in enclosing function, not in a block
```go
for _, f := range files {
    defer f.Close()   // all closes happen when the FUNCTION returns, not the loop iteration
}
// Fix: wrap in IIFE (immediately invoked function expression)
for _, f := range files {
    func() {
        defer f.Close()
        process(f)
    }()
}
```

---

## Idiomatic Go

```go
// ✅ Always check errors — don't use _ for err unless you know it can't fail
result, err := strconv.Atoi("123")
if err != nil {
    return fmt.Errorf("parse error: %w", err)
}

// ✅ Keep functions small and focused
// ✅ Use named returns only when they add clarity (rare)
// ✅ Prefer explicit returns over naked returns in functions > ~10 lines

// ✅ Functional options pattern — the Go substitute for default params
type ServerConfig struct {
    port    int
    timeout time.Duration
}

type Option func(*ServerConfig)

func WithPort(p int) Option {
    return func(c *ServerConfig) { c.port = p }
}

func NewServer(opts ...Option) *ServerConfig {
    cfg := &ServerConfig{port: 8080, timeout: 30 * time.Second}
    for _, opt := range opts {
        opt(cfg)
    }
    return cfg
}

// Usage: clean and extensible
s := NewServer(WithPort(9090))

// ✅ Error wrapping with context
if err := db.Query(); err != nil {
    return fmt.Errorf("fetchUser %d: %w", userID, err)
}
```

---

## Exercises

See [`exercises/01-basics/`](../exercises/01-basics/README.md) — Part 3: Functions

---

## Related Files

- Cheatsheet: [`cheatsheets/syntax.md`](../cheatsheets/syntax.md)
- Next concept: [`concepts/slices-maps.md`](slices-maps.md)
- Error handling deep dive: [`concepts/error-handling.md`](error-handling.md)
