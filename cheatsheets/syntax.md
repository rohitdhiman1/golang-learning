# Cheatsheet — Go Syntax (vs Java & Python)

> Fill as you complete Phase 1. Use as a quick-reference before interviews.

---

## Variables

| Task | Java | Python | Go |
|---|---|---|---|
| Declare + assign | `int x = 5;` | `x = 5` | `x := 5` |
| Declare only | `int x;` | — | `var x int` |
| Constant | `final int X = 5;` | `X = 5` | `const X = 5` |
| Swap | `int t=a; a=b; b=t;` | `a,b = b,a` | `a,b = b,a` |
| Multiple assign | — | `a,b = 1,2` | `a,b := 1,2` |
| Unused var | Warning | OK | **Compile error** |

---

## Types

| Go Type | Size | Java Equiv | Python Equiv | Zero Value |
|---|---|---|---|---|
| `bool` | 1 byte | `boolean` | `bool` | `false` |
| `int` | platform (64-bit usually) | `long` | `int` | `0` |
| `int8/16/32/64` | fixed | `byte/short/int/long` | — | `0` |
| `uint`, `uint8`… | unsigned | — | — | `0` |
| `float32` | 4 bytes | `float` | — | `0.0` |
| `float64` | 8 bytes | `double` | `float` | `0.0` |
| `string` | immutable, UTF-8 | `String` | `str` | `""` |
| `rune` | int32, Unicode point | `char` | — | `0` |
| `byte` | uint8 | `byte` | — | `0` |

---

## Type Conversions (always explicit in Go)

```go
var i int = 42
f := float64(i)       // int → float64
n := int(3.9)         // float → int (truncates, not rounds → 3)
s := strconv.Itoa(i)  // int → string ("42")
i2, _ := strconv.Atoi("42")  // string → int
s2 := fmt.Sprintf("%d", i)   // any → string (formatted)
```

---

## fmt Verbs

| Verb | Meaning | Example |
|---|---|---|
| `%v` | Default format | `[1 2 3]`, `{Alice 30}` |
| `%+v` | Struct with field names | `{Name:Alice Age:30}` |
| `%#v` | Go syntax | `main.Person{Name:"Alice"}` |
| `%T` | Type | `int`, `[]string` |
| `%d` | Integer | `42` |
| `%f` | Float | `3.140000` |
| `%.2f` | Float, 2 decimal places | `3.14` |
| `%s` | String | `hello` |
| `%q` | Quoted string | `"hello"` |
| `%t` | Bool | `true` |
| `%p` | Pointer address | `0xc000...` |
| `%c` | Rune as character | `A` |
| `%b/%o/%x` | Binary/Octal/Hex | `101010` / `52` / `2a` |

---

## Control Flow

```go
// if — no parentheses, always braces
if x > 0 { ... } else if x < 0 { ... } else { ... }

// if with init statement
if err := fn(); err != nil { return err }

// for — three styles
for i := 0; i < n; i++ { ... }          // classic
for condition { ... }                     // while-style
for { ... }                               // infinite

// range
for i, v := range slice { ... }
for k, v := range mapVar { ... }
for i, ch := range str { ... }           // ch is rune
for v := range channel { ... }

// switch — no fallthrough by default
switch x {
case 1, 2:  ...
default:    ...
}

switch {                  // no expression — like if/else chain
case x < 0: ...
default:    ...
}

// explicit fallthrough (rare)
switch x {
case 1:
    doOne()
    fallthrough
case 2:
    doTwo()
}

// defer — LIFO, runs after function returns
defer cleanup()
defer fmt.Println("done")
```

---

## Functions

```go
func name(a, b int) int { return a + b }
func name(a int) (int, error) { return a, nil }     // multiple return
func name(nums ...int) int { ... }                   // variadic
func(x int) int { return x * 2 }                    // anonymous
```

---

## Key Mental Model Differences from Java/Python

| "I'd do X in Java/Python" | Go's way |
|---|---|
| `try { } catch (Exception e) { }` | `val, err := fn(); if err != nil { ... }` |
| `null` — forgot to initialise | Zero value — always initialised |
| Ternary `x > 0 ? "a" : "b"` | No ternary — use `if/else` |
| `for`, `while`, `do-while` | `for` only |
| Switch falls through | Switch does NOT fall through |
| String is an object/class | String is a primitive-ish value type |
| `int i; // uninitialised` | `var i int // = 0` always |

---

## Common Pitfalls

- `x := y := z` — not valid, use `x, y := ..., ...`
- Unused variable → compile error
- `string(65)` → `"A"` not `"65"` (use `strconv.Itoa`)
- `len("héllo")` → 6 (bytes), not 5 (chars)
- Switch in Go ≠ switch in Java (no auto-fallthrough)
