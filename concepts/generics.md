# Generics

> **Phase 2 · Week 4**  
> Covers: type parameters, constraints, `comparable`, when to use vs when not to

---

## What It Is

Generics (added in Go 1.18) let you write functions and types that work with **any type** satisfying a **constraint**. Before generics, you'd use `interface{}` (losing type safety) or write duplicate code for each type.

Go generics are deliberately simpler than Java/C# generics — no variance, no wildcards, no specialisation.

---

## Java/Python Comparison

| Concept | Java | Python | Go |
|---|---|---|---|
| Syntax | `<T extends Comparable>` | `TypeVar` / `Generic[T]` (3.12: `def f[T]()`) | `[T comparable]` |
| Constraints | `extends` + interfaces | Protocol / bound | Interface-based constraints |
| Type erasure | Yes (at runtime, generic info gone) | N/A (no compile-time generics) | Monomorphised at compile time (like C++) |
| Wildcards | `? extends`, `? super` | Not applicable | Not available — simpler model |
| Variance | Covariant/contravariant | N/A | No variance — Go is invariant |

**Key insight:** Go generics exist to reduce code duplication in a type-safe way. They're not meant for building elaborate type hierarchies. The Go philosophy: use concrete types first, add generics only when you have 3+ callers that differ only in type.

---

## How It Works

### Basic generic function

```go
func Map[T, U any](s []T, f func(T) U) []U {
    result := make([]U, len(s))
    for i, v := range s {
        result[i] = f(v)
    }
    return result
}

// Usage — type inference figures out T and U
names := Map(users, func(u User) string { return u.Name })
```

### Constraints

A constraint is an **interface** that restricts which types a type parameter accepts.

```go
// any — no restriction (alias for interface{})
func Print[T any](v T) { fmt.Println(v) }

// comparable — supports == and != (maps need this for keys)
func Contains[T comparable](s []T, target T) bool {
    for _, v := range s {
        if v == target {
            return true
        }
    }
    return false
}

// Custom constraint using type union (Go 1.18+)
type Number interface {
    ~int | ~int8 | ~int16 | ~int32 | ~int64 |
    ~float32 | ~float64
}

func Sum[T Number](nums []T) T {
    var total T
    for _, n := range nums {
        total += n
    }
    return total
}
```

The `~` prefix means "any type whose underlying type is X" — so `type Celsius float64` satisfies `~float64`.

### Generic types

```go
type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(v T) {
    s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
    if len(s.items) == 0 {
        var zero T
        return zero, false
    }
    v := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return v, true
}
```

### The `cmp` and `slices` stdlib packages (Go 1.21+)

```go
import (
    "cmp"
    "slices"
)

// Sort any ordered slice
slices.Sort([]int{3, 1, 2})              // [1, 2, 3]
slices.SortFunc(users, func(a, b User) int {
    return cmp.Compare(a.Name, b.Name)
})

// Contains, Index, Compact, etc.
slices.Contains([]string{"a", "b"}, "a") // true
```

---

## Key Gotchas

1. **Type inference works for functions, not types:**
   ```go
   Map(nums, double)     // OK — T and U inferred
   Stack[int]{}          // Must specify — no inference for types
   ```

2. **No method type parameters** — you can't add type params to methods, only functions and types:
   ```go
   // NOT ALLOWED:
   // func (s Stack[T]) Map[U any](f func(T) U) Stack[U]
   
   // Instead, use a top-level function:
   func MapStack[T, U any](s Stack[T], f func(T) U) Stack[U] { ... }
   ```

3. **No specialisation** — you can't write a different implementation for a specific type.

4. **Zero value of generic type** — use `var zero T` to get the zero value:
   ```go
   func First[T any](s []T) (T, bool) {
       if len(s) == 0 {
           var zero T
           return zero, false
       }
       return s[0], true
   }
   ```

5. **`comparable` includes structs and arrays** — any type with all comparable fields satisfies `comparable`. Slices and maps do not.

---

## Idiomatic Go

- **Concrete types first** — don't reach for generics as your first tool. If you have one caller, use the concrete type.
- **Generics for data structures and utility functions** — `Map`, `Filter`, `Contains`, `Stack[T]` are ideal use cases.
- **Don't use generics to avoid implementing an interface** — if `sort.Interface` works, use it.
- **Prefer the `slices` and `maps` stdlib packages** over writing your own generic utilities (Go 1.21+).
- **Keep constraints narrow** — use `comparable` when you need `==`, `cmp.Ordered` when you need `<`, `any` when you need nothing.
- **"A little copying is better than a little dependency"** — applies to generics too. Three copy-pasted functions are often clearer than one generic function with a complex constraint.

---

## Exercises

See [`exercises/03-interfaces/`](../exercises/03-interfaces/README.md) — Part 6: Generics

---

## Related Files

- Previous: [`concepts/pointers.md`](pointers.md)
- Next: [`concepts/error-handling.md`](error-handling.md)
