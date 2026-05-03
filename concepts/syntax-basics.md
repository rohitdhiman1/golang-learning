# Syntax Basics & Control Flow

> **Phase 1 · Week 1**  
> Covers: variables, types, zero values, type inference, basic I/O, if/else, for, switch, defer

---

## What It Is

The building blocks of every Go program: how to declare values, what types exist, how to control execution flow, and how to defer cleanup. Go's syntax is intentionally minimal — fewer ways to do things means more consistent code.

---

## Java/Python Comparison

### Variables

| Concept | Java | Python | Go |
|---|---|---|---|
| Declare + assign | `int x = 5;` | `x = 5` | `x := 5` or `var x int = 5` |
| Declare only | `int x;` | N/A | `var x int` (zero value: `0`) |
| Constant | `final int X = 5;` | `X = 5` (convention) | `const X = 5` |
| Type inference | `var x = 5;` (Java 10+) | Always inferred | `:=` infers type |
| Unused variable | Compile warning | Fine | **Compile error** |

### Types

| Concept | Java | Python | Go |
|---|---|---|---|
| Integer | `int`, `long` | `int` | `int`, `int8/16/32/64`, `uint...` |
| Float | `double`, `float` | `float` | `float64`, `float32` |
| String | `String` (object) | `str` | `string` (immutable, UTF-8) |
| Boolean | `boolean` | `bool` | `bool` |
| Character | `char` | N/A | `rune` (int32, Unicode code point) |
| Byte | `byte` | `bytes` element | `byte` (alias for `uint8`) |
| No value | `null` | `None` | `nil` (only for pointers, slices, maps, channels, funcs, interfaces) |

### Control Flow

| Concept | Java | Python | Go |
|---|---|---|---|
| If condition | `if (x > 0) {}` | `if x > 0:` | `if x > 0 {}` (no parens, always braces) |
| Loop | `for`, `while`, `do-while` | `for`, `while` | `for` only (no while, no do-while) |
| Switch | Falls through by default | `match` (3.10+) | **Does NOT fall through** by default |
| Switch fallthrough | Need `break` to stop | N/A | Explicit `fallthrough` keyword |

---

## How It Works

### Variables

```go
// Three ways to declare a variable:

// 1. Full declaration (works at package level and inside functions)
var name string = "Alice"
var count int     // zero value: 0
var ok bool       // zero value: false

// 2. Short declaration — ONLY inside functions
x := 42           // inferred as int
greeting := "hi"  // inferred as string

// 3. Multiple assignment
a, b := 1, 2
a, b = b, a       // swap — no temp variable needed

// Constants
const Pi = 3.14159
const MaxRetries = 3
const Greeting = "Hello"

// Typed constants (less common)
const TypedPi float64 = 3.14159

// iota — auto-incrementing constant (like Java enums, simpler)
const (
    Sunday = iota  // 0
    Monday         // 1
    Tuesday        // 2
)
```

### Zero Values

Every type in Go has a zero value — variables are NEVER uninitialised.

```go
var i int       // 0
var f float64   // 0.0
var b bool      // false
var s string    // "" (empty string)
var p *int      // nil
var sl []int    // nil (but len(nil slice) == 0, it's usable)
var m map[string]int  // nil (reading returns zero value, writing panics)
```

**From Java:** Unlike Java, there's no `NullPointerException` for basic types. But nil pointer dereference still panics for pointers and interfaces.

### Types

```go
// Integer types — use plain int unless you have a specific reason
var a int    = 42
var b int64  = 9_000_000_000  // underscores for readability
var c uint   = 7              // unsigned

// Floating point — float64 is the default
var price float64 = 9.99
rate := 0.075   // inferred as float64

// Strings — UTF-8 encoded, immutable
name := "Gopher"
multiline := `This is a
raw string literal`          // backtick = no escape sequences

// Runes and bytes
r := 'A'          // rune (int32) — single quotes
b2 := byte('A')   // byte (uint8)

// Type conversion — explicit, never implicit
var n int = 42
var f2 float64 = float64(n)   // must convert explicitly
var n2 int = int(f2)
```

### Basic I/O

```go
import "fmt"

// Print
fmt.Println("Hello, World!")     // adds newline
fmt.Print("no newline")
fmt.Printf("Name: %s, Age: %d\n", name, age)

// Format a string (like Java's String.format or Python's f-strings)
msg := fmt.Sprintf("Score: %d", score)

// Common format verbs
// %v  — default format (works for any type)
// %T  — type of the value
// %d  — integer
// %f  — float (%0.2f for 2 decimal places)
// %s  — string
// %q  — quoted string
// %t  — bool
// %p  — pointer address
// %+v — struct with field names
// %#v — Go syntax representation

fmt.Printf("%T\n", 42)       // int
fmt.Printf("%v\n", []int{1,2,3})  // [1 2 3]
```

### If / Else

```go
// No parentheses around condition — always use braces
if x > 0 {
    fmt.Println("positive")
} else if x < 0 {
    fmt.Println("negative")
} else {
    fmt.Println("zero")
}

// Init statement — declare a variable scoped to the if block
// This is idiomatic Go — very common with error checks
if err := doSomething(); err != nil {
    fmt.Println("error:", err)
    return
}
// err is not accessible here

// From Java: no ternary operator (x > 0 ? "a" : "b") in Go
// Just use a plain if/else
```

### For — Go's Only Loop

```go
// 1. Classic C-style for
for i := 0; i < 5; i++ {
    fmt.Println(i)
}

// 2. While-style (condition only)
n := 1
for n < 100 {
    n *= 2
}

// 3. Infinite loop
for {
    // break to exit
    if done {
        break
    }
}

// 4. Range — iterate over slice, map, string, channel
fruits := []string{"apple", "banana", "cherry"}
for i, v := range fruits {
    fmt.Printf("%d: %s\n", i, v)
}

// Ignore index with _
for _, v := range fruits {
    fmt.Println(v)
}

// Range over string iterates by rune (Unicode code point), not byte
for i, ch := range "héllo" {
    fmt.Printf("%d: %c\n", i, ch)
}

// Range over map (order is random — by design)
m := map[string]int{"a": 1, "b": 2}
for k, v := range m {
    fmt.Printf("%s=%d\n", k, v)
}

// break and continue work as in Java/Python
for i := 0; i < 10; i++ {
    if i == 3 {
        continue   // skip 3
    }
    if i == 7 {
        break      // stop at 7
    }
    fmt.Println(i)
}

// Labeled break — exit outer loop from inner loop
outer:
for i := 0; i < 5; i++ {
    for j := 0; j < 5; j++ {
        if i+j == 6 {
            break outer
        }
    }
}
```

### Switch

```go
// Go switch does NOT fall through by default (opposite of Java!)
day := "Monday"
switch day {
case "Monday", "Tuesday", "Wednesday", "Thursday", "Friday":
    fmt.Println("Weekday")
case "Saturday", "Sunday":
    fmt.Println("Weekend")
default:
    fmt.Println("Unknown")
}

// Switch with no expression — acts like if/else chain
x := 42
switch {
case x < 0:
    fmt.Println("negative")
case x == 0:
    fmt.Println("zero")
default:
    fmt.Println("positive")
}

// Switch with init statement
switch os := runtime.GOOS; os {
case "darwin":
    fmt.Println("macOS")
case "linux":
    fmt.Println("Linux")
default:
    fmt.Println("Other:", os)
}

// Explicit fallthrough (unusual — use sparingly)
switch x {
case 1:
    fmt.Println("one")
    fallthrough   // executes case 2 body too
case 2:
    fmt.Println("two or fell through from one")
}
```

### Defer

```go
// defer pushes a function call onto a stack — runs AFTER the surrounding function returns
// Order: LIFO (last in, first out)

func readFile(path string) error {
    f, err := os.Open(path)
    if err != nil {
        return err
    }
    defer f.Close()   // runs when readFile returns, regardless of how it exits

    // ... process file
    return nil
}

// Multiple defers — run in LIFO order
func example() {
    defer fmt.Println("third")
    defer fmt.Println("second")
    defer fmt.Println("first")  // this runs first (LIFO)
}
// Output: first, second, third

// Defer with a loop — WRONG pattern (all defers stack up)
for _, f := range files {
    defer f.Close()   // BAD: all closes happen after function returns, not each iteration
}
// CORRECT: wrap in an anonymous function
for _, f := range files {
    func() {
        defer f.Close()   // closes after this anonymous func returns
        process(f)
    }()
}
```

---

## Key Gotchas

### 1. Unused variables are compile errors
```go
x := 5
// ... never use x
// → compile error: x declared and not used
// Fix: use _ if you need to discard a value
_, err := doSomething()
```

### 2. Short declaration `:=` requires at least one new variable on left
```go
x := 1
x := 2   // ERROR: no new variables on left side of :=
x = 2    // OK: plain assignment
x, y := 2, 3  // OK: y is new
```

### 3. Integer division truncates (same as Java, different from Python)
```go
fmt.Println(7 / 2)   // 3, not 3.5
fmt.Println(7 % 2)   // 1
// For float division, at least one operand must be float64:
fmt.Println(7.0 / 2) // 3.5
```

### 4. String indexing returns bytes, not characters
```go
s := "hello"
fmt.Println(s[0])        // 104 (byte value of 'h'), not 'h'
fmt.Printf("%c\n", s[0]) // 'h'
// For Unicode: use range or convert to []rune
```

### 5. Defer captures variables by reference, not by value
```go
x := 1
defer fmt.Println(x)   // prints 1 (arg evaluated now)
x = 2
// But deferred function body sees final x:
defer func() { fmt.Println(x) }()  // prints 2
```

### 6. Switch cases don't fall through — remember to add `fallthrough` if you want it
```go
// Coming from Java: remove all those `break` statements — Go doesn't need them
```

---

## Idiomatic Go

```go
// ✅ Use := inside functions, var at package level
func greet(name string) string {
    msg := fmt.Sprintf("Hello, %s", name)
    return msg
}

// ✅ Init statement in if — keeps error variable scoped
if err := json.Unmarshal(data, &v); err != nil {
    return err
}

// ✅ Early return / guard clauses (not deep nesting)
func process(x int) error {
    if x < 0 {
        return errors.New("negative input")
    }
    if x == 0 {
        return errors.New("zero input")
    }
    // happy path here
    doWork(x)
    return nil
}

// ✅ Blank identifier _ to discard values you don't need
for _, v := range items {
    fmt.Println(v)
}

// ✅ Defer for cleanup — always pair Open/Lock with a deferred Close/Unlock
f, _ := os.Open("file.txt")
defer f.Close()

// ❌ Don't use var for local variables when you have a value to assign
var x int = 5   // verbose
x := 5          // idiomatic

// ❌ Don't nest — prefer early returns
if ok {
    if valid {
        doThing()  // bad
    }
}
// vs
if !ok || !valid {
    return
}
doThing()   // better
```

---

## Exercises

See [`exercises/01-basics/`](../exercises/01-basics/README.md)

- Part 1: Variables & Types
- Part 2: Control Flow
- Part 3: Defer & Scope

---

## Related Files

- Cheatsheet: [`cheatsheets/syntax.md`](../cheatsheets/syntax.md)
- Next concept: [`concepts/functions.md`](functions.md)
