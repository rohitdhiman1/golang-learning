# Cheatsheet — Go Idioms & Proverbs

> Fill as you progress through all phases. The "Go way" of thinking.

---

## Go Proverbs (Rob Pike)

> "Don't communicate by sharing memory; share memory by communicating."  
> "Concurrency is not parallelism."  
> "The bigger the interface, the weaker the abstraction."  
> "Make the zero value useful."  
> "A little copying is better than a little dependency."  
> "Clear is better than clever."  
> "Errors are values."  
> "Don't just check errors, handle them gracefully."  
> "Design the architecture, name the components, document the details."

---

## Naming Conventions

_Fill as you learn — coming in Phase 1+._

Topics to capture here:
- `MixedCaps` not `snake_case` (unlike Python)
- Short variable names in small scopes: `i`, `v`, `k`, `err`, `ok`
- Acronyms all-caps: `URL`, `HTTP`, `ID` (so `ServeHTTP`, not `ServeHttp`)
- Interface naming: single-method interfaces end in `-er` (`Reader`, `Stringer`, `Handler`)
- Package names: short, lowercase, no underscores
- Error variable: `err` (always — don't name it `e` or `error`)
- Constructor: `NewXxx` returns `*Xxx` or `Xxx`
- Don't repeat package name: `http.Handler` not `http.HTTPHandler`

---

## Common Idioms

_Fill as you encounter them._

```go
// Comma-ok pattern (map lookup, type assertion, channel receive)
v, ok := m[key]
v, ok := x.(Type)
v, ok := <-ch

// Early return (guard clauses — avoid deep nesting)
if err != nil { return err }

// Defer for cleanup
defer f.Close()
defer mu.Unlock()

// Functional options (substitute for default params)
// (fill with template when learned in Phase 2)

// Table-driven tests
// (fill with template when learned in Phase 4)
```

---

_Placeholder — fill continuously as you progress._
