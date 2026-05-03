# Interview Questions & Answers

> **Phase 6 · Week 11–12**  
> Fill answers as you study each topic. Use for self-testing before interviews.

---

## How to Use

Cover the Answer column, read the Question, answer out loud, then reveal. If you hesitate, add the topic to [`docs/gaps.md`](../docs/gaps.md).

---

## Phase 1 — Foundations

| # | Question | Answer |
|---|---|---|
| 1 | What is the zero value of `int`, `bool`, `string`, `*int`? | `0`, `false`, `""`, `nil` |
| 2 | What's the difference between `var x int` and `x := 0`? | `var` works at package level; `:=` is short declaration, function-only. Both produce `int` with value `0`. |
| 3 | Why are unused variables a compile error in Go? | Forces clean code — no "declare-and-forget". Reduces bugs from typos. |
| 4 | What does `defer` do? When does it run? | Pushes a function call onto a stack. Runs after the enclosing function returns, LIFO order. |
| 5 | Can Go's `for` loop replace `while`? | Yes — `for condition {}` is Go's while loop. `for {}` is infinite. |
| 6 | Does Go's `switch` fall through by default? | No. Use explicit `fallthrough` if needed. Opposite of Java. |

---

## Phase 2 — Type System

| # | Question | Answer |
|---|---|---|
| 7 | How does Go implement interfaces differently from Java? | Implicitly — no `implements` keyword. Any type that has the required methods satisfies the interface. |
| 8 | What is the `error` type in Go? | An interface: `type error interface { Error() string }`. Any type with an `Error()` method is an error. |
| 9 | When should you use a pointer receiver vs value receiver? | Pointer if the method modifies the receiver, if the struct is large (avoid copy), or to be consistent with other methods on the type. |
| 10 | What is the "nil interface" trap? | An interface holding a typed nil pointer is NOT nil itself — `interface != nil` even if the value inside is nil. |
| 11 | What's `errors.Is` vs `errors.As`? | `Is` checks if an error in the chain matches a target value/sentinel. `As` checks if any error in the chain can be assigned to a target type (for extracting typed errors). |

---

## Phase 3 — Concurrency

| # | Question | Answer |
|---|---|---|
| 12 | What is a goroutine? How is it different from a thread? | A goroutine is a lightweight function executing concurrently, multiplexed onto OS threads by the Go scheduler. Starts at 2KB stack (vs MB for threads), grows as needed. |
| 13 | What is the GMP scheduling model? | G=goroutines, M=OS threads, P=logical processors. P limits parallel execution (GOMAXPROCS). Go uses work-stealing. |
| 14 | What's the difference between unbuffered and buffered channels? | Unbuffered: sender blocks until receiver is ready (synchronous rendezvous). Buffered: sends don't block until buffer is full. |
| 15 | How do you detect a goroutine leak? | Use `go test -race`, check `runtime.NumGoroutine()`, use `pprof` goroutine profile, ensure all channels are closed / goroutines have an exit path. |
| 16 | What does `select` do? | Waits on multiple channel operations simultaneously, executes the first one that's ready. `default` case makes it non-blocking. |

---

## Phase 4 — Standard Library

| # | Question | Answer |
|---|---|---|
| 17 | What is `io.Reader`? Why is it important? | Interface: `Read(p []byte) (n int, err error)`. Anything readable implements it — files, HTTP bodies, strings. Enables composable I/O pipelines. |
| 18 | How do you make HTTP requests time out in Go? | Set `Timeout` on `http.Client`: `client := &http.Client{Timeout: 10 * time.Second}`. Or use `context.WithTimeout` and `req.WithContext(ctx)`. |
| 19 | What is `context.Context` used for? | Carries deadlines, cancellations, and request-scoped values across API boundaries and goroutines. |
| 20 | What is the reference time in Go's `time.Format`? | `Mon Jan 2 15:04:05 MST 2006` — each field is a specific number (month=1, day=2, hour=15, minute=04, second=05, year=2006). |

---

## Phase 5–6 — Advanced / Internals

| # | Question | Answer |
|---|---|---|
| 21 | What are the three fields of a slice header? | Pointer to backing array, length, capacity. |
| 22 | What happens when you `append` past capacity? | Go allocates a new (roughly 2x) backing array, copies elements, returns a new slice header. |
| 23 | Why is map iteration order random? | Intentionally randomised by the runtime since Go 1.0 to prevent code from relying on implementation-specific ordering. |
| 24 | How does Go's garbage collector work? | Tri-color mark-and-sweep (concurrent). Objects are white (unmarked), grey (queued), black (reachable). White objects after mark phase are swept. |
| 25 | What is escape analysis? | The compiler determines whether a variable can live on the stack or must escape to the heap. Use `go build -gcflags="-m"` to see decisions. |
| 26 | What is `init()` and when does it run? | A special function (can be multiple per package) that runs after package-level vars are initialised and before `main()`. Cannot be called manually. |

---

## Code Review — "What's wrong with this code?"

See [`interview/code-review.md`](code-review.md)
