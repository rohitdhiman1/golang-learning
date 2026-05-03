# Pointers

> **Phase 2 · Week 3**  
> Covers: `&` and `*`, pointer receivers, nil pointers, no pointer arithmetic

---

## What It Is

A pointer holds the **memory address** of a value. Go has pointers but deliberately omits pointer arithmetic (unlike C). Pointers let you share data without copying and allow functions/methods to modify their arguments.

---

## Java/Python Comparison

| Concept | Java | Python | Go |
|---|---|---|---|
| References | All objects are references (hidden) | All objects are references (hidden) | Explicit: `*T` is a pointer to `T` |
| Take address | Not possible directly | Not possible directly | `p := &x` |
| Dereference | Automatic | Automatic | Explicit: `*p` (but auto-deref for methods) |
| Null/None | `null` (reference can be null) | `None` | `nil` (only for pointers, slices, maps, channels, interfaces, functions) |
| Pass by value | Primitives by value, objects by reference | Everything by "object reference" | **Everything by value** — but a pointer value IS the address |

**Key insight:** Go is **always pass-by-value**. When you pass a pointer, the pointer itself is copied — but the copy points to the same data. This is explicit, unlike Java where "is it a reference?" depends on whether it's a primitive or object.

---

## How It Works

### Basics: `&` and `*`

```go
x := 42
p := &x          // p is *int, holds address of x
fmt.Println(*p)  // 42 — dereference: read the value at the address
*p = 100         // modify x through the pointer
fmt.Println(x)   // 100
```

### `new(T)` vs `&T{}`

```go
// Both create a *User pointing to a zero-value User
p1 := new(User)          // rarely used
p2 := &User{}            // preferred — allows field initialisation
p3 := &User{Name: "Al"}  // most common
```

### Pointer receivers

```go
type Counter struct {
    n int
}

// Pointer receiver — modifies the actual Counter
func (c *Counter) Increment() {
    c.n++  // Go auto-dereferences: same as (*c).n++
}

c := Counter{}
c.Increment()  // Go auto-takes address: same as (&c).Increment()
fmt.Println(c.n) // 1
```

### When to use pointers

| Use a pointer | Use a value |
|---|---|
| Function needs to modify the argument | Small, immutable data (int, small structs) |
| Struct is large (avoid copying) | You want a copy (safety) |
| Need to represent "absent" with nil | Slices, maps, channels (already reference types) |
| Pointer receiver methods | Value semantics are fine |

### Pointer to pointer (rare)

```go
x := 42
p := &x
pp := &p  // **int — pointer to pointer
fmt.Println(**pp) // 42
```

Almost never needed in Go. If you see `**T`, reconsider the design.

---

## Key Gotchas

1. **nil pointer dereference** — the most common pointer bug:
   ```go
   var p *User     // p is nil
   fmt.Println(p.Name) // PANIC: nil pointer dereference
   ```
   Always check for nil before dereferencing when the pointer may be nil.

2. **Slices, maps, and channels are already reference types** — you don't need `*[]int`. Passing a slice passes a header (pointer to backing array + len + cap). But: appending to a slice inside a function won't be visible outside unless you return the new slice or pass `*[]int`.

3. **Return a pointer to a local variable is fine** — Go's escape analysis moves it to the heap:
   ```go
   func newUser() *User {
       u := User{Name: "Alice"} // allocated on heap, not stack
       return &u                 // safe — Go handles this
   }
   ```

4. **Auto-dereference convenience** — Go automatically dereferences struct pointers for field access:
   ```go
   p := &User{Name: "Alice"}
   fmt.Println(p.Name)     // no need for (*p).Name
   ```

5. **No pointer arithmetic** — you cannot do `p++` or `p + offset`. This is intentional — it eliminates a huge class of C bugs.

6. **Method set difference** — `T` values can only call value receiver methods. `*T` values can call both. This matters for interface satisfaction (see interfaces.md).

---

## Idiomatic Go

- **Don't overuse pointers** — Go's value semantics are a feature, not a limitation. Small structs are fine to copy.
- **Use pointer receivers consistently** — if one method needs a pointer receiver, make all methods on that type pointer receivers.
- **Prefer `&T{}` over `new(T)`** — it's more readable and allows field initialisation.
- **Return values, not pointers, for small types** — `func Min(a, b int) int`, not `*int`.
- **Check nil before use** when a pointer can be nil — especially in exported functions.
- **Don't use `*[]T` or `*map[K]V`** unless you specifically need to replace the entire slice/map.

---

## Exercises

See [`exercises/03-interfaces/`](../exercises/03-interfaces/README.md) — Part 2: Pointers

---

## Related Files

- Previous: [`concepts/interfaces.md`](interfaces.md)
- Next: [`concepts/generics.md`](generics.md)
- Related: [`concepts/structs-methods.md`](structs-methods.md) (pointer receivers)
