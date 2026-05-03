# Interfaces

> **Phase 2 · Week 3**  
> Covers: implicit satisfaction (duck typing), empty interface (`any`), type assertions, type switches

---

## What It Is

An interface in Go is a **set of method signatures**. Any type that implements all the methods satisfies the interface — no `implements` keyword needed. This is called **implicit (structural) satisfaction**.

Interfaces are the primary abstraction mechanism in Go. They decouple callers from implementations.

---

## Java/Python Comparison

| Concept | Java | Python | Go |
|---|---|---|---|
| Define interface | `interface Foo { ... }` | `class Foo(ABC)` | `type Foo interface { ... }` |
| Implement | `class X implements Foo` | `class X(Foo)` | Just have the methods — nothing to declare |
| Check at compile time | Yes | No (duck typing at runtime) | Yes (structural, checked at compile time) |
| Empty interface | `Object` | `object` | `any` (alias for `interface{}`) |
| Type check | `instanceof` | `isinstance()` | Type assertion: `v.(T)` |

**Key insight:** Go interfaces combine Java's compile-time safety with Python's duck typing. You define the interface where it's *used* (the consumer), not where it's *implemented* (the provider).

---

## How It Works

### Defining and implementing

```go
type Stringer interface {
    String() string
}

type User struct {
    Name string
    Age  int
}

// User satisfies Stringer — no "implements" keyword
func (u User) String() string {
    return fmt.Sprintf("%s (age %d)", u.Name, u.Age)
}
```

### Key stdlib interfaces to know

```go
fmt.Stringer     // String() string
error            // Error() string
io.Reader        // Read(p []byte) (n int, err error)
io.Writer        // Write(p []byte) (n int, err error)
io.Closer        // Close() error
sort.Interface   // Len(), Less(i,j), Swap(i,j)
http.Handler     // ServeHTTP(ResponseWriter, *Request)
```

### The empty interface: `any`

`any` (alias for `interface{}`) is satisfied by every type. Use it when you genuinely don't know the type (like `json.Unmarshal`'s target). Avoid it as a crutch — if you know the type, use it.

```go
func printAnything(v any) {
    fmt.Println(v)
}
```

### Type assertions

Extract the concrete type from an interface:

```go
var s Stringer = User{Name: "Alice", Age: 30}

// Single-value form — panics if wrong type
u := s.(User)

// Comma-ok form — safe
u, ok := s.(User)
if ok {
    fmt.Println(u.Name)
}
```

### Type switches

Handle multiple possible types:

```go
func describe(v any) string {
    switch x := v.(type) {
    case int:
        return fmt.Sprintf("integer: %d", x)
    case string:
        return fmt.Sprintf("string: %q", x)
    case User:
        return fmt.Sprintf("user: %s", x.Name)
    default:
        return fmt.Sprintf("unknown: %T", x)
    }
}
```

### Interface internals: the (type, value) pair

An interface variable holds two things internally:
1. A pointer to the **type descriptor** (what concrete type is stored)
2. A pointer to the **value**

An interface is `nil` only when **both** are nil.

---

## Key Gotchas

1. **The nil interface trap** — the most common Go interface bug:
   ```go
   var p *User = nil
   var s Stringer = p
   fmt.Println(s == nil) // false! s has type=*User, value=nil
   ```
   An interface holding a nil pointer is NOT a nil interface.

2. **Pointer vs value receiver and interface satisfaction:**
   ```go
   type Speaker interface { Speak() string }
   func (d *Dog) Speak() string { return "Woof" }

   var s Speaker = Dog{}   // COMPILE ERROR — Dog value doesn't have *Dog methods
   var s Speaker = &Dog{}  // OK — *Dog has Speak()
   ```

3. **Accept interfaces, return structs** — a Go proverb. Functions should accept the smallest interface they need, but return concrete types so callers have full access.

4. **Don't define interfaces preemptively.** Define them at the call site when you need to decouple. One-method interfaces are ideal.

5. **Interface pollution** — defining large interfaces forces all implementors to implement everything. Keep interfaces small (1-3 methods).

---

## Idiomatic Go

- **Small interfaces** — `io.Reader` has ONE method. Most Go interfaces have 1-3 methods.
- **Define interfaces at the consumer** — the package that *uses* the interface, not the package that *implements* it.
- **Name single-method interfaces with `-er` suffix**: `Reader`, `Writer`, `Stringer`, `Handler`.
- **"The bigger the interface, the weaker the abstraction"** — Rob Pike.
- **Accept interfaces, return structs** — gives callers flexibility without hiding capabilities.
- **Don't return interfaces to hide implementation** — that's a Java pattern. In Go, return the concrete type.

---

## Exercises

See [`exercises/03-interfaces/`](../exercises/03-interfaces/README.md) — Parts 3 & 4

---

## Related Files

- Previous: [`concepts/structs-methods.md`](structs-methods.md)
- Next: [`concepts/pointers.md`](pointers.md)
- Cheatsheet: [`cheatsheets/idioms.md`](../cheatsheets/idioms.md)
