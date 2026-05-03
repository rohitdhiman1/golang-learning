# Exercises — 02: Data Structures

> **Phase 1 · Week 2**  
> Read [`concepts/slices-maps.md`](../../concepts/slices-maps.md) before starting.

Work through the parts in order. Create one `.go` file per part. Run with `go run <file>.go`.

---

## Part 1 — Slices

**File:** `part1_slices.go`

1. **Basic operations** — Create a slice `[]int{5, 3, 1, 4, 2}`. Print its length and capacity. Append `6` and `7`. Print the updated slice.

2. **Slice expressions** — Given `s := []int{0, 1, 2, 3, 4, 5, 6}`, extract:
   - The first 3 elements
   - The last 3 elements (use `len(s)` in the expression)
   - Elements 2 through 4 (inclusive)
   Print each sub-slice.

3. **Shared backing array** — Create `a := []int{1, 2, 3, 4, 5}` and `b := a[1:3]`. Modify `b[0]`. Print both `a` and `b` to confirm they share the backing array. Then fix it using `copy`.

4. **Word counter** — Given a sentence string, split it into words, count the words, and return the sorted unique words. Use `strings.Fields`, `append`, and `sort.Strings`.
   ```
   Input:  "the quick brown fox jumps over the lazy dog the"
   Output: unique words: [brown dog fox jumps lazy over quick the], count: 8
   ```

5. **Dedup in-place** — Given `[]int{1, 1, 2, 3, 3, 3, 4, 5, 5}` (sorted), remove consecutive duplicates in-place without allocating a new slice. Print the result using the returned length.

6. **Rotate left** — Rotate a slice left by `k` positions in-place (no extra allocation).
   ```
   Input:  [1 2 3 4 5], k=2
   Output: [3 4 5 1 2]
   ```
   Hint: three-reversal trick.

---

## Part 2 — Maps

**File:** `part2_maps.go`

1. **Word frequency** — Count how many times each word appears in a string. Print results in sorted key order.
   ```
   Input:  "go is great go is fast go"
   Output: fast:1 go:3 great:1 is:2
   ```

2. **Grouping** — Given `[]string{"apple", "avocado", "banana", "blueberry", "cherry", "cranberry"}`, group them by their first letter into `map[string][]string`. Print each group.

3. **Two-sum** — Given a slice of ints and a target, find the indices of the two numbers that add up to target (guaranteed one solution). Use a map for O(n) time.
   ```
   Input:  [2, 7, 11, 15], target=9
   Output: [0, 1]
   ```

4. **Inverted index** — Given `[]string{"go is fast", "go is simple", "python is dynamic"}`, build a `map[string][]int` mapping each word to the sentence indices it appears in. Print the index for `"is"` and `"go"`.

5. **Default values** — Write a function `wordLengths(words []string) map[string]int` that maps each word to its length. Demonstrate reading a missing key (returns 0) vs the two-value lookup.

---

## Part 3 — Structs

**File:** `part3_structs.go`

1. **Student roster** — Define a `Student` struct with `Name string`, `Grade int`, `Score float64`. Create a slice of 5 students. Print the one with the highest score.

2. **Value vs pointer copy** — Create a `Point struct{ X, Y int }`. Assign it to another variable and modify the copy. Confirm the original is unchanged. Then repeat with `*Point` and confirm the original IS changed.

3. **Anonymous struct** — Create an anonymous struct for a temporary HTTP response summary: `{ Status int; Body string; Headers map[string]string }`. Populate it and print with `%+v`.

4. **Struct comparison** — Define `type RGB struct{ R, G, B uint8 }`. Create two RGB values representing the same colour. Compare with `==`. Now add a `[]string` field — observe that the struct becomes non-comparable.

5. **JSON marshal/unmarshal** — Define a `Person` struct with JSON tags:
   ```go
   type Person struct {
       Name  string `json:"name"`
       Age   int    `json:"age"`
       Email string `json:"email,omitempty"`
   }
   ```
   - Marshal a `Person` to JSON, print it
   - Unmarshal `{"name":"Alice","age":30}` into a `Person`, print it
   - Confirm `omitempty` works when `Email` is empty

---

## Checklist

After completing all parts:

- [ ] All files compile: `go build ./exercises/02-data-structures/`
- [ ] All files pass vet: `go vet ./exercises/02-data-structures/`
- [ ] You can explain what a slice header is (3 fields)
- [ ] You can explain why modifying a sub-slice modifies the original
- [ ] You can explain when `append` allocates a new array vs reuses the existing one
- [ ] You know why map iteration order is random
- [ ] You can explain what `omitempty` does in JSON struct tags
- [ ] Add anything still unclear to [`docs/gaps.md`](../../docs/gaps.md)
