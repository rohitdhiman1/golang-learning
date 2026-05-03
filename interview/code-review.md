# Code Review Exercises — "What's wrong with this code?"

> **Phase 6 · Week 11–12**  
> Read each snippet, identify the bug(s) before reading the answer.

---

## 1 — The goroutine loop capture

```go
func main() {
    for i := 0; i < 5; i++ {
        go func() {
            fmt.Println(i)
        }()
    }
    time.Sleep(time.Second)
}
```

<details>
<summary>What's wrong?</summary>

All goroutines share the same `i` variable. By the time they run, the loop has finished and `i == 5`, so all goroutines print `5`.

**Fix:** Pass `i` as an argument or shadow it:
```go
go func(n int) { fmt.Println(n) }(i)
// or:
i := i
go func() { fmt.Println(i) }()
```

</details>

---

## 2 — Writing to a nil map

```go
func wordCount(s string) map[string]int {
    var counts map[string]int
    for _, word := range strings.Fields(s) {
        counts[word]++
    }
    return counts
}
```

<details>
<summary>What's wrong?</summary>

`counts` is a nil map. Assigning to a nil map panics: `assignment to entry in nil map`.

**Fix:**
```go
counts := make(map[string]int)
```

</details>

---

## 3 — Ignoring the error

```go
func readConfig(path string) string {
    data, _ := os.ReadFile(path)
    return string(data)
}
```

<details>
<summary>What's wrong?</summary>

The error is silently discarded. If the file doesn't exist or can't be read, `data` is nil and the function returns an empty string with no indication of failure.

**Fix:** Return the error or handle it explicitly.
```go
func readConfig(path string) (string, error) {
    data, err := os.ReadFile(path)
    if err != nil {
        return "", fmt.Errorf("readConfig: %w", err)
    }
    return string(data), nil
}
```

</details>

---

## 4 — Defer in a loop

```go
func processFiles(paths []string) error {
    for _, p := range paths {
        f, err := os.Open(p)
        if err != nil {
            return err
        }
        defer f.Close()
        // process f...
    }
    return nil
}
```

<details>
<summary>What's wrong?</summary>

All `defer f.Close()` calls are queued and execute when `processFiles` returns — not at the end of each loop iteration. For large numbers of files this holds all file descriptors open until the function exits, potentially exhausting OS limits.

**Fix:** Wrap the body in a closure:
```go
for _, p := range paths {
    if err := func() error {
        f, err := os.Open(p)
        if err != nil { return err }
        defer f.Close()
        // process f...
        return nil
    }(); err != nil {
        return err
    }
}
```

</details>

---

## 5 — Slice mutation surprise

```go
func firstTwo(s []int) []int {
    result := s[:2]
    result[0] = 999
    return result
}

func main() {
    data := []int{1, 2, 3, 4}
    sub := firstTwo(data)
    fmt.Println(data)   // what prints?
    fmt.Println(sub)
}
```

<details>
<summary>What's wrong?</summary>

`s[:2]` shares the backing array with `data`. Mutating `result[0] = 999` also changes `data[0]`. Output:
```
[999 2 3 4]
[999 2]
```

**Fix:** If you want an independent copy:
```go
result := make([]int, 2)
copy(result, s[:2])
```

</details>

---

## 6 — The nil interface trap

```go
type MyError struct{ msg string }
func (e *MyError) Error() string { return e.msg }

func mayFail(fail bool) error {
    var err *MyError
    if fail {
        err = &MyError{"something went wrong"}
    }
    return err   // BUG: always non-nil!
}

func main() {
    err := mayFail(false)
    if err != nil {
        fmt.Println("got error:", err)   // prints even when fail=false!
    }
}
```

<details>
<summary>What's wrong?</summary>

`return err` returns an `error` interface value with type `*MyError` and value `nil`. The interface itself is NOT nil because it has a type. So `err != nil` is `true` even when `fail=false`.

**Fix:** Return an untyped nil:
```go
if fail {
    return &MyError{"something went wrong"}
}
return nil   // untyped nil — interface is truly nil
```

</details>

---

## 7 — Race condition

```go
var count int

func increment(wg *sync.WaitGroup) {
    defer wg.Done()
    count++
}

func main() {
    var wg sync.WaitGroup
    for i := 0; i < 1000; i++ {
        wg.Add(1)
        go increment(&wg)
    }
    wg.Wait()
    fmt.Println(count)
}
```

<details>
<summary>What's wrong?</summary>

`count++` is not atomic — it's a read-modify-write that can be interrupted. Multiple goroutines race on `count`. Run with `-race` to confirm.

**Fix:** Use `sync/atomic` or protect with a mutex:
```go
var count int64
// in increment:
atomic.AddInt64(&count, 1)
```

</details>

---

_Add more snippets as you progress through phases._
