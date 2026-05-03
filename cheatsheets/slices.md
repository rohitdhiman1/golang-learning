# Cheatsheet — Slices

> Fill as you complete Phase 1. Critical for interviews — slice internals are a favourite question.

---

## Slice Header (3 fields)

```
┌───────────┬───────┬──────────┐
│  *array   │  len  │   cap    │
└───────────┴───────┴──────────┘
  pointer     length  capacity
  to backing  (valid  (total
  array       elems)  space)
```

A slice is a **descriptor** — it does not own data, it points to an array.

---

## Creation

```go
var s []int                  // nil slice: ptr=nil, len=0, cap=0
s := []int{}                 // empty slice: ptr≠nil, len=0, cap=0
s := []int{1, 2, 3}         // literal: len=3, cap=3
s := make([]int, 5)          // len=5, cap=5, all zeros
s := make([]int, 3, 10)      // len=3, cap=10
```

**nil vs empty:**
```go
var s []int    // s == nil → true
s2 := []int{}  // s2 == nil → false
// Both: len=0, range works, append works
// Difference: nil marshals to JSON null; empty marshals to []
```

---

## Slicing

```go
s := []int{0, 1, 2, 3, 4}
s[1:3]   // [1 2]      len=2, cap=4  (shares backing array)
s[:2]    // [0 1]      len=2, cap=5
s[2:]    // [2 3 4]    len=3, cap=3
s[:]     // [0 1 2 3 4] full view
s[1:3:4] // [1 2]      len=2, cap=3  (limits cap — safe append)
```

**Key rule:** `len = high - low`, `cap = cap(original) - low`

---

## Append

```go
s = append(s, x)           // append one element
s = append(s, a, b, c)     // append multiple
s = append(s, other...)    // append another slice

// When len == cap: Go allocates new array (~2x cap), copies, returns new header
// Always assign back: s = append(s, ...)
```

---

## Copy

```go
dst := make([]int, len(src))
n := copy(dst, src)   // copies min(len(dst), len(src)) elements
                       // returns number of elements copied
// dst is now INDEPENDENT of src
```

---

## Common Patterns

```go
// Append to nil/empty slice (idiomatic)
var result []string
result = append(result, "a")

// Pre-allocate for performance
result := make([]T, 0, len(input))

// Filter in-place (no allocation)
n := 0
for _, v := range s {
    if keep(v) { s[n] = v; n++ }
}
s = s[:n]

// Delete at index i (order doesn't matter)
s[i] = s[len(s)-1]
s = s[:len(s)-1]

// Delete preserving order
s = append(s[:i], s[i+1:]...)

// Clone / independent copy
clone := append([]int{}, s...)
// or:
clone := make([]int, len(s))
copy(clone, s)

// Stack (LIFO)
stack = append(stack, v)           // push
top := stack[len(stack)-1]         // peek
stack = stack[:len(stack)-1]       // pop

// Queue (FIFO) — use carefully, may leave large backing array
queue = append(queue, v)           // enqueue
front := queue[0]; queue = queue[1:] // dequeue (leaks backing array)
```

---

## Gotchas

### 1. Sub-slice shares backing array
```go
a := []int{1, 2, 3, 4, 5}
b := a[1:3]   // [2 3] — SAME backing array
b[0] = 99
// a is now [1 99 3 4 5]  ← surprise!
// Fix: use copy or three-index slice to limit cap
```

### 2. Append into a sub-slice can overwrite the original
```go
a := make([]int, 3, 6)  // [0 0 0], cap=6
b := a[:2]              // [0 0], cap=6
b = append(b, 99)       // still within cap → writes to a[2]!
// a is now [0 0 99]
// Fix: b := a[:2:2]  — cap limited to 2, append forces new array
```

### 3. `append` return value must be captured
```go
append(s, 4)        // compile error
s = append(s, 4)    // correct
```

### 4. `len()` vs `cap()` vs number of underlying elements
```go
s := make([]int, 3, 10)
len(s)  // 3 — accessible elements
cap(s)  // 10 — total before reallocation
```

---

## Interview Questions

**Q: What are the three fields of a slice header?**  
A: Pointer to backing array, length, capacity.

**Q: What happens when you append past capacity?**  
A: Go allocates a new (larger) backing array, copies all elements, returns a new slice header. The old backing array is still referenced by any slices that point to it.

**Q: What's the difference between `nil` and empty slice?**  
A: Nil slice has a nil pointer and zero len/cap. Empty slice has a non-nil pointer but zero len/cap. Both are functionally equivalent for `append`, `range`, `len` — but nil marshals to JSON `null`, empty marshals to `[]`.

**Q: Why is `delete(s[i])` written as `s = append(s[:i], s[i+1:]...)`?**  
A: Slices don't have a built-in delete. This expression re-slices around the element and appends the suffix, overwriting index i.

**Q: When should you use `copy` vs slicing?**  
A: Use `copy` when you need an independent copy of data. Sub-slicing gives you a view (shared data).
