# Slices, Maps & Structs

> **Phase 1 · Week 2**  
> Covers: arrays, slices (length vs capacity, append, copy), maps, structs (basic)

---

## What It Is

The core data structures you'll use in 90% of Go programs. Slices are Go's primary collection type (not arrays). Maps are hash tables. Structs are typed collections of fields — Go's substitute for classes. Understanding slice internals (the three-field header) is critical for avoiding subtle bugs.

---

## Java/Python Comparison

| Concept | Java | Python | Go |
|---|---|---|---|
| Fixed array | `int[]` | N/A | `[3]int` |
| Dynamic list | `ArrayList<T>` | `list` | `[]T` (slice) |
| Hash table | `HashMap<K,V>` | `dict` | `map[K]V` |
| Struct/Record | Class (no methods) | dataclass | `struct` |
| Null key in map | Yes | Yes | No — keys must be comparable |
| Ordered map | `LinkedHashMap` | `dict` (3.7+) | No (use slice of structs) |

---

## How It Works

### Arrays

Arrays in Go have a **fixed size that's part of the type**. Rarely used directly — slices are preferred.

```go
// [N]T — size is fixed and part of the type
var a [3]int              // [0 0 0] — zero values
b := [3]int{1, 2, 3}     // literal
c := [...]int{4, 5, 6}   // compiler counts the elements

fmt.Println(len(b))   // 3
fmt.Println(b[0])     // 1

// Arrays are values — copying an array copies all elements
x := [3]int{1, 2, 3}
y := x   // y is a full copy
y[0] = 99
fmt.Println(x[0])   // 1 — unchanged

// [3]int and [4]int are DIFFERENT types — can't assign one to the other
```

### Slices

A slice is a view into an underlying array. It has three fields: **pointer to array, length, capacity**.

```go
// Declare a nil slice (pointer is nil, len=0, cap=0)
var s []int   // nil slice — usable, just empty

// Slice literal — creates and references an array
s = []int{1, 2, 3, 4, 5}
fmt.Println(len(s))   // 5
fmt.Println(cap(s))   // 5

// make([]T, length, capacity) — preallocate
s2 := make([]int, 3)      // [0 0 0], len=3, cap=3
s3 := make([]int, 3, 10)  // [0 0 0], len=3, cap=10

// Slicing — creates a new slice header pointing to the SAME backing array
s := []int{0, 1, 2, 3, 4}
s1 := s[1:3]   // [1 2], shares backing array with s
s1[0] = 99
fmt.Println(s)  // [0 99 2 3 4] — mutation visible through original!

// s[low:high]  → elements from index low up to (not including) high
// s[:high]     → from 0 to high
// s[low:]      → from low to end
// s[:]         → full slice (same backing array)

// Three-index slice — also limits capacity
s2 := s[1:3:4]   // s2 has len=2, cap=3 (4-1)
```

### Append

```go
s := []int{1, 2, 3}
s = append(s, 4)         // [1 2 3 4]
s = append(s, 5, 6, 7)   // [1 2 3 4 5 6 7]

// Append another slice — spread with ...
other := []int{8, 9}
s = append(s, other...)  // [1 2 3 4 5 6 7 8 9]

// ⚠️ append may or may not allocate a new backing array
// When cap is exceeded, Go doubles the capacity (roughly) and copies to new array
// Always assign back: s = append(s, val)

// Pre-allocate when you know the size to avoid repeated reallocations
result := make([]int, 0, len(input))
for _, v := range input {
    result = append(result, transform(v))
}
```

### Copy

```go
src := []int{1, 2, 3, 4, 5}
dst := make([]int, 3)   // only 3 elements
n := copy(dst, src)     // copies min(len(dst), len(src)) elements
fmt.Println(n)    // 3
fmt.Println(dst)  // [1 2 3] — independent copy, not sharing backing array

// Trick: copy a slice to get an independent duplicate
clone := make([]int, len(src))
copy(clone, src)
// Or: clone := append([]int{}, src...)
```

### Slice Patterns

```go
// Nil vs empty slice — different but both have len == 0
var nilSlice []int         // nil == true
emptySlice := []int{}      // nil == false
// Both work with range, append, len — but nil slice marshals to JSON null

// Delete element at index i (order doesn't matter)
s := []int{1, 2, 3, 4, 5}
i := 2
s[i] = s[len(s)-1]
s = s[:len(s)-1]   // [1 2 5 4]

// Delete element preserving order
s = append(s[:i], s[i+1:]...)

// Filter in-place (no allocation)
n := 0
for _, v := range s {
    if keep(v) {
        s[n] = v
        n++
    }
}
s = s[:n]
```

### Maps

```go
// map[KeyType]ValueType
// Keys must be comparable (==): string, int, bool, struct with comparable fields
// NOT valid as keys: slice, map, function

// Declare a nil map — reading returns zero value, WRITING PANICS
var m map[string]int   // nil — don't write to this!

// Create with make
m = make(map[string]int)
m["alice"] = 42
m["bob"] = 7

// Map literal
scores := map[string]int{
    "alice": 95,
    "bob":   87,
    "carol": 91,
}

// Read — returns zero value if key doesn't exist (no panic, no error)
v := scores["dave"]   // 0 — dave not in map
fmt.Println(v)

// Two-value read — check if key exists
v, ok := scores["alice"]
if ok {
    fmt.Println("alice:", v)
} else {
    fmt.Println("alice not found")
}

// Delete a key
delete(scores, "bob")

// Iterate — order is RANDOM (by design, to prevent reliance on ordering)
for k, v := range scores {
    fmt.Printf("%s: %d\n", k, v)
}

// Length
fmt.Println(len(scores))

// Maps are reference types — assigning copies the reference
m2 := scores         // m2 and scores point to same underlying map
m2["alice"] = 0      // also changes scores["alice"]
```

### Structs

```go
// Define a struct type
type Person struct {
    Name string
    Age  int
    Email string
}

// Create a struct — positional (fragile, avoid)
p1 := Person{"Alice", 30, "alice@example.com"}

// Create with field names (idiomatic — field order doesn't matter)
p2 := Person{
    Name:  "Bob",
    Age:   25,
    Email: "bob@example.com",
}

// Zero-value struct — all fields are zero values
var p3 Person   // Person{"", 0, ""}

// Access fields
fmt.Println(p2.Name)   // Bob
p2.Age = 26

// Anonymous structs — for one-off use (e.g., JSON, test data)
point := struct {
    X, Y int
}{X: 10, Y: 20}

// Structs are values — copying creates an independent copy
a := Person{Name: "Alice", Age: 30}
b := a
b.Name = "Bob"
fmt.Println(a.Name)   // Alice — unchanged

// Pointer to struct — modifications affect original
p := &Person{Name: "Alice", Age: 30}
p.Age = 31   // auto-dereferenced, same as (*p).Age = 31

// Struct comparison — structs with all comparable fields can be compared with ==
type Point struct{ X, Y int }
p1 := Point{1, 2}
p2 := Point{1, 2}
fmt.Println(p1 == p2)   // true
```

### Struct Embedding (preview — Phase 2 deep dive)

```go
// Embedding is Go's form of composition — not inheritance
type Animal struct {
    Name string
}

func (a Animal) Speak() string {
    return a.Name + " speaks"
}

type Dog struct {
    Animal   // embedded — Dog "inherits" Name and Speak()
    Breed string
}

d := Dog{Animal: Animal{Name: "Rex"}, Breed: "Labrador"}
fmt.Println(d.Speak())   // Rex speaks — promoted method
fmt.Println(d.Name)      // Rex — promoted field
```

---

## Key Gotchas

### 1. Slices share the backing array — mutations are visible through all slice headers
```go
s := []int{1, 2, 3, 4, 5}
a := s[1:3]   // [2 3]
a[0] = 99
fmt.Println(s)   // [1 99 3 4 5] — shared!

// To get an independent copy:
b := make([]int, len(a))
copy(b, a)
```

### 2. Append after slicing — capacity surprise
```go
s := make([]int, 3, 6)   // len=3, cap=6
a := s[:2]               // len=2, cap=6 (shares backing array)
a = append(a, 99)        // appends IN PLACE — s[2] is now 99!
fmt.Println(s)            // [0 0 99]

// Use three-index slice to limit capacity and force copy on append:
a = s[:2:2]              // len=2, cap=2
a = append(a, 99)        // cap exceeded → new allocation → safe
```

### 3. Map read on a nil map doesn't panic, but write does
```go
var m map[string]int
v := m["key"]   // OK, returns 0
m["key"] = 1    // PANIC: assignment to entry in nil map
```

### 4. Map iteration order is random
```go
// Don't rely on iteration order. If you need sorted keys:
keys := make([]string, 0, len(m))
for k := range m {
    keys = append(keys, k)
}
sort.Strings(keys)
for _, k := range keys {
    fmt.Println(k, m[k])
}
```

### 5. Struct fields must be exported (uppercase) to be visible outside the package
```go
type person struct {   // lowercase type — unexported
    name string        // lowercase field — unexported
    Age  int           // uppercase field — exported
}
// json.Marshal only sees exported fields — a common source of empty JSON
```

### 6. `append` return value must be captured
```go
s := []int{1, 2, 3}
append(s, 4)        // compile error: append evaluated but not used
s = append(s, 4)    // correct
```

---

## Idiomatic Go

```go
// ✅ Use slices, not arrays (arrays are for specific performance/size needs)
items := []string{"a", "b", "c"}

// ✅ Pre-allocate slices when size is known
result := make([]int, 0, len(input))

// ✅ Use two-value map lookup to check key existence
if v, ok := m[key]; ok {
    use(v)
}

// ✅ Use struct field names in literals — never positional for exported structs
// Positional breaks if fields are reordered
p := Person{Name: "Alice", Age: 30}   // good
p  = Person{"Alice", 30, ""}          // fragile — avoid

// ✅ Use nil slice (not empty slice) as zero value for function returns
func getItems() []string {
    if nothing {
        return nil   // preferred over return []string{}
    }
    return items
}

// ✅ Range over map — ignore value with _ if you only need keys
for k := range m {
    fmt.Println(k)
}

// ✅ Struct embedding for composition — not a class hierarchy
type Logger struct{ prefix string }
type Server struct {
    Logger          // gets Logger's methods promoted
    addr string
}
```

---

## Exercises

See [`exercises/02-data-structures/`](../exercises/02-data-structures/README.md)

- Part 1: Slices — word counter, dedup, rotate
- Part 2: Maps — frequency table, group-by, word index
- Part 3: Structs — student roster, JSON marshal/unmarshal

---

## Related Files

- Cheatsheet: [`cheatsheets/slices.md`](../cheatsheets/slices.md)
- Next concept: [`concepts/structs-methods.md`](structs-methods.md) (deep dive — Phase 2)
