# Exercises — 01: Basics

> **Phase 1 · Week 1**  
> Read [`concepts/syntax-basics.md`](../../concepts/syntax-basics.md) and [`concepts/functions.md`](../../concepts/functions.md) before starting.

Work through the parts in order. Create one `.go` file per part in this folder (e.g., `part1_variables.go`). Run with `go run part1_variables.go`.

---

## Part 1 — Variables & Types

**File:** `part1_variables.go`

1. **Hello, World** — Print `Hello, World!` to stdout.

2. **Variable declaration** — Declare variables using all three methods:
   - `var` with explicit type
   - `var` with initialiser (type inferred)
   - `:=` short declaration
   Print each variable's value and type using `%T`.

3. **Zero values** — Declare one variable of each type (`int`, `float64`, `bool`, `string`) without initialising. Print each and confirm they are zero values.

4. **Constants & iota** — Define a `DayOfWeek` constant group using `iota` for Sunday through Saturday. Print the value of `Wednesday`.

5. **Type conversion** — Start with `var x int = 42`. Convert it to `float64`, then to `string` using `strconv.Itoa` (not a cast — try `string(x)` too and observe the difference).

6. **Multiple assignment** — Declare `a, b := 10, 20`. Swap their values in one line without a temp variable. Print before and after.

---

## Part 2 — Control Flow

**File:** `part2_control.go`

1. **FizzBuzz** — Print numbers 1–30. For multiples of 3 print `Fizz`, multiples of 5 print `Buzz`, multiples of both print `FizzBuzz`.

2. **Infinite loop with break** — Write a loop that starts at 1, doubles each iteration, and breaks when the value exceeds 1000. Print the final value.

3. **Range over string** — Range over the string `"héllo, 世界"` and print each character's index and rune value. Observe that byte index jumps for multi-byte characters.

4. **Switch — grade classifier** — Given a score (int), print the letter grade:
   - 90–100 → A, 80–89 → B, 70–79 → C, 60–69 → D, below 60 → F
   Use a switch with no expression (condition-based cases).

5. **Labeled break** — Write a 4×4 nested loop. Break out of both loops when `i + j == 5`. Print how many iterations completed.

6. **Defer order** — Write a function that defers three `fmt.Println` calls with values `"a"`, `"b"`, `"c"`. Call the function and confirm LIFO output order.

---

## Part 3 — Functions

**File:** `part3_functions.go`

1. **Multiple return** — Write `divide(a, b float64) (float64, error)` that returns an error if `b == 0`. Test with valid and invalid inputs.

2. **Named return** — Write `minMax(nums []int) (min, max int)`. Use a naked return. Test with `[]int{5, 1, 9, 3, 7}`.

3. **Variadic sum** — Write `sum(nums ...int) int`. Call it with individual values, then spread a slice with `...`.

4. **First-class function** — Write `applyAll(nums []int, fn func(int) int) []int` that applies `fn` to every element and returns a new slice. Test with a doubling function.

5. **Closure counter** — Write a `counter()` function that returns a closure. Each call to the closure increments and returns an independent count. Create two counters and confirm they're independent.

6. **Closure loop gotcha** — Demonstrate the loop variable capture bug. Then show the fix (either shadowing `i := i` or passing as an argument to the goroutine/func).

7. **Defer in a loop** — Write a function that opens 3 "files" (simulated with `fmt.Println`). Show the wrong pattern (defer in loop) and the correct IIFE fix.

---

## Checklist

After completing all parts:

- [ ] All files compile without errors: `go build ./exercises/01-basics/`
- [ ] All files pass vet: `go vet ./exercises/01-basics/`
- [ ] You can explain why `string(42)` gives `"*"` not `"42"`
- [ ] You can explain LIFO defer order
- [ ] You can explain the closure loop capture bug and two ways to fix it
- [ ] Add anything still unclear to [`docs/gaps.md`](../../docs/gaps.md)
