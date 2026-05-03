# Structs & Methods

> **Phase 2 · Week 3**  
> Covers: receiver functions (value vs pointer), constructor patterns, embedding, composition over inheritance

---

## What It Is

Go has no classes. Instead, you define **structs** (data) and attach **methods** (behaviour) to them via receiver functions. Composition replaces inheritance — you embed one struct inside another to reuse fields and methods.

---

## Java/Python Comparison

| Concept | Java/Python | Go |
|---|---|---|
| Class | `class Foo { ... }` | `type Foo struct { ... }` + methods |
| Constructor | `new Foo(...)` / `Foo(...)` | `func NewFoo(...) *Foo` (convention, not language) |
| `this`/`self` | Implicit | Explicit receiver: `func (f *Foo) Bar()` |
| Inheritance | `extends` / `class Child(Parent)` | No inheritance — use embedding |
| Override | Override method in subclass | Outer struct's method shadows embedded method |
| Access control | `private`/`public`/`protected` | Uppercase = exported, lowercase = unexported (package-level, not type-level) |

Key shift: In Java, you think in class hierarchies. In Go, you think in **small structs + interfaces**. There is no `super`, no `abstract class`, no `protected`.

---

## How It Works

### Defining a struct

```go
type User struct {
    Name  string
    Email string
    Age   int
}
```

### Value receiver vs pointer receiver

```go
// Value receiver — operates on a COPY of the struct
func (u User) Greeting() string {
    return "Hi, " + u.Name
}

// Pointer receiver — can MODIFY the original struct
func (u *User) SetEmail(email string) {
    u.Email = email
}
```

**When to use which:**

| Use value receiver | Use pointer receiver |
|---|---|
| Method doesn't modify the struct | Method modifies the struct |
| Struct is small (few fields, no slices/maps) | Struct is large (avoids copying) |
| You want immutability semantics | You need consistency (if one method is pointer, make all pointer) |

**Rule of thumb:** If in doubt, use a pointer receiver. Be consistent within a type — if any method needs a pointer receiver, make them all pointer receivers.

### Constructor pattern

Go has no constructors. By convention, use `NewXxx` functions:

```go
func NewUser(name, email string, age int) *User {
    return &User{
        Name:  name,
        Email: email,
        Age:   age,
    }
}
```

Return `*T` (pointer) when:
- The struct is large
- Methods use pointer receivers
- You need nil as a sentinel (uninitialized)

Return `T` (value) when:
- The struct is small and immutable (like `time.Time`)

### Embedding (composition)

```go
type Address struct {
    Street string
    City   string
}

type Employee struct {
    User            // embedded — Employee "has a" User
    Address         // embedded — Employee "has an" Address
    Department string
}
```

Embedding promotes fields and methods:

```go
e := Employee{
    User:       User{Name: "Alice", Email: "a@co.com", Age: 30},
    Address:    Address{Street: "123 Main", City: "NYC"},
    Department: "Engineering",
}
e.Name         // promoted from User — no need for e.User.Name
e.Greeting()   // promoted from User
e.City         // promoted from Address
```

If both embedded types have a field/method with the same name, you must disambiguate: `e.User.Name` vs `e.Address.Name`.

### Method sets

The **method set** of a type determines which interfaces it satisfies:

| Type | Method set includes |
|---|---|
| `T` (value) | Only value receiver methods |
| `*T` (pointer) | Both value AND pointer receiver methods |

This is why `*User` satisfies more interfaces than `User`.

---

## Key Gotchas

1. **Value receiver doesn't mutate** — a common bug for Java/Python developers:
   ```go
   func (u User) SetName(name string) { u.Name = name } // BUG: modifies copy
   ```

2. **Embedding is not inheritance** — the outer type does NOT become the inner type. An `Employee` is not a `User`. You cannot pass `Employee` where `User` is expected (but you can pass `e.User`).

3. **nil pointer receiver** — methods can be called on nil pointers (no NPE like Java). Handle it:
   ```go
   func (u *User) String() string {
       if u == nil {
           return "<nil user>"
       }
       return u.Name
   }
   ```

4. **Exported vs unexported** — access control is at the package level, not the type level. All code in the same package can see unexported fields. There is no `protected`.

5. **Embedding and shadowing** — if the outer struct defines a method that the embedded struct also has, the outer method wins (shadows). This looks like overriding but isn't — the inner method is still accessible via `e.User.Method()`.

---

## Idiomatic Go

- **Prefer composition over inheritance** — embed small, focused structs
- **Keep structs small** — if a struct has 10+ fields, consider splitting
- **Constructors return pointers** (`NewXxx() *Xxx`) when methods use pointer receivers
- **Be consistent** with receiver types — don't mix value and pointer receivers on the same type
- **Don't embed just to save typing** — embed when the outer type genuinely "has a" relationship
- **Name receivers short** — `u` for `User`, `s` for `Server`, not `this` or `self`
- **Zero value should be useful** — design structs so `var x MyType` is a valid, usable value (like `sync.Mutex`, `bytes.Buffer`)

---

## Exercises

See [`exercises/03-interfaces/`](../exercises/03-interfaces/README.md) — Part 1: Structs & Methods

---

## Related Files

- Previous: [`concepts/slices-maps.md`](slices-maps.md)
- Next: [`concepts/interfaces.md`](interfaces.md)
- Cheatsheet: [`cheatsheets/idioms.md`](../cheatsheets/idioms.md)
