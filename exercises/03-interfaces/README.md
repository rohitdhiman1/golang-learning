# Exercises — 03: Interfaces & Type System

> **Phase 2 · Week 3–4**  
> Read [`concepts/structs-methods.md`](../../concepts/structs-methods.md), [`concepts/interfaces.md`](../../concepts/interfaces.md), [`concepts/pointers.md`](../../concepts/pointers.md), [`concepts/generics.md`](../../concepts/generics.md), [`concepts/error-handling.md`](../../concepts/error-handling.md) before starting.

---

## Part 1 — Structs & Methods

**File:** `part1_structs_methods.go`

1. **Value vs pointer receiver** — Define `type Rect struct { Width, Height float64 }`. Add `Area() float64` (value receiver) and `Scale(factor float64)` (pointer receiver). Create a Rect, scale it, print the area before and after.

2. **Constructor pattern** — Define `type Circle struct { Radius float64 }` with `NewCircle(r float64) (*Circle, error)` that returns an error if radius ≤ 0. Add `Area() float64` and `Perimeter() float64` methods.

3. **Embedding** — Define `type Animal struct { Name string }` with `Speak() string` returning `"..."`. Define `type Dog struct { Animal; Breed string }` with its own `Speak()` that returns `"Woof!"`. Show that Dog shadows Animal's Speak. Access both.

4. **Method set demo** — Create an interface `type Resizer interface { Scale(float64) }`. Show that `Rect` value does NOT satisfy it (Scale has pointer receiver), but `*Rect` does.

5. **Zero-value usefulness** — Define `type Counter struct { n int }` with `Increment()` and `Value() int`. Show that `var c Counter` is immediately usable without a constructor.

---

## Part 2 — Pointers

**File:** `part2_pointers.go`

1. **Basics** — Declare `x := 42`. Take its address, print the pointer and dereferenced value. Modify `x` through the pointer. Print `x` to confirm.

2. **Swap function** — Write `func swap(a, b *int)` that swaps two values via pointers. Test it.

3. **Nil pointer guard** — Write `func safeName(u *User) string` that returns `"<nil>"` if u is nil, otherwise `u.Name`. Test with both nil and non-nil.

4. **Pointer to struct vs value** — Create a function that takes `*Rect` and doubles its width. Show the caller sees the change. Then create a version that takes `Rect` (value) — show the caller does NOT see the change.

5. **Slice append visibility** — Demonstrate that appending inside a function is not visible to the caller (slice header is copied). Then fix it by returning the new slice.

---

## Part 3 — Interfaces

**File:** `part3_interfaces.go`

1. **Shape interface** — Define `type Shape interface { Area() float64; Perimeter() float64 }`. Make `Rect` and `Circle` (from Part 1) satisfy it. Write `func printShape(s Shape)` that prints both. Test with both types.

2. **Stringer** — Implement `fmt.Stringer` for a `type Colour struct { R, G, B uint8 }`. Printing a Colour should output `#RRGGBB` hex format.

3. **Sort interface** — Define `type ByAge []Person` (with `Name string; Age int`). Implement `sort.Interface` (Len, Less, Swap). Sort a slice of people by age.

4. **Interface composition** — Define `type ReadWriter interface { io.Reader; io.Writer }`. Create a `type Buffer struct { data []byte }` that satisfies `ReadWriter`.

5. **Accept interfaces, return structs** — Write `func summarise(r io.Reader) (string, error)` that reads all bytes from any Reader and returns the content as a string. Test with `strings.NewReader`.

---

## Part 4 — Type Assertions & Switches

**File:** `part4_type_assertions.go`

1. **Safe type assertion** — Given `var v any = "hello"`, extract the string using comma-ok form. Then try asserting to `int` — confirm ok is false.

2. **Type switch** — Write `func describe(v any) string` that returns a description for `int`, `string`, `bool`, `[]int`, and a default. Test with several values.

3. **The nil interface trap** — Define a custom error type. Write a function that returns `error` but internally uses a typed nil pointer. Demonstrate that the returned error is non-nil even though the pointer is nil. Then show the fix.

4. **Interface-based polymorphism** — Define `type Notification interface { Send() string }`. Implement `EmailNotification` and `SMSNotification`. Write `func broadcast(notifications []Notification)` that calls Send on each. Use a type switch inside broadcast to also print the notification type.

---

## Part 5 — Error Handling

**File:** `part5_errors.go`

1. **Basic error** — Write `func parseAge(s string) (int, error)` using `strconv.Atoi`. Wrap the error with context using `fmt.Errorf` and `%w`. Test with valid and invalid inputs.

2. **Sentinel errors** — Define `var ErrTooYoung = errors.New("too young")` and `var ErrTooOld = errors.New("too old")`. Write `func validateAge(age int) error` that returns the appropriate sentinel. Use `errors.Is` to check.

3. **Custom error type** — Define `type FieldError struct { Field, Message string }` implementing the `error` interface. Write a validation function that returns `*FieldError`. Use `errors.As` to extract the field name.

4. **Error wrapping chain** — Build a 3-layer call chain: `handleRequest → validateInput → parseField`. Each layer wraps the error with context. At the top, use `errors.Is` to find the root cause, and `fmt.Println(err)` to see the full chain.

5. **panic and recover** — Write `func safeDivide(a, b int) (result int, err error)` that uses defer+recover to catch a divide-by-zero panic and return it as an error instead.

---

## Part 6 — Generics

**File:** `part6_generics.go`

1. **Generic Contains** — Write `func Contains[T comparable](s []T, target T) bool`. Test with `[]int` and `[]string`.

2. **Generic Map** — Write `func Map[T, U any](s []T, f func(T) U) []U`. Test by mapping `[]int` to `[]string` and `[]string` to `[]int` (lengths).

3. **Generic Filter** — Write `func Filter[T any](s []T, pred func(T) bool) []T`. Test by filtering even numbers and strings starting with "a".

4. **Type constraint** — Define `type Number interface { ~int | ~float64 }`. Write `func Sum[T Number](nums []T) T`. Test with `int`, `float64`, and a custom type `type Celsius float64`.

5. **Generic Stack** — Implement `type Stack[T any]` with `Push(T)`, `Pop() (T, bool)`, `Peek() (T, bool)`, `Len() int`. Test with `Stack[int]` and `Stack[string]`.

---

## Checklist

After completing all parts:

- [ ] All files compile: `go build ./exercises/03-interfaces/`
- [ ] All files pass vet: `go vet ./exercises/03-interfaces/`
- [ ] You can explain value vs pointer receivers and when to use each
- [ ] You can explain implicit interface satisfaction
- [ ] You can explain the nil interface trap
- [ ] You know when to use `errors.Is` vs `errors.As`
- [ ] You can explain when generics are appropriate vs overkill
- [ ] Add anything still unclear to [`docs/gaps.md`](../../docs/gaps.md)
