package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func Part2Pointers() {
	// Exercise 1: Basics
	fmt.Println("=== Exercise 1: Pointer Basics ===")
	x := 42
	p := &x
	fmt.Printf("  x = %d, p = %p, *p = %d\n", x, p, *p)
	*p = 100
	fmt.Printf("  after *p = 100: x = %d\n", x)

	// Exercise 2: Swap function
	fmt.Println("\n=== Exercise 2: Swap via Pointers ===")
	a, b := 10, 20
	fmt.Printf("  before: a=%d, b=%d\n", a, b)
	swap(&a, &b)
	fmt.Printf("  after swap(&a, &b): a=%d, b=%d\n", a, b)

	// Exercise 3: Nil pointer guard
	fmt.Println("\n=== Exercise 3: Nil Pointer Guard ===")
	u := &User{Name: "Alice"}
	fmt.Printf("  safeName(user):  %q\n", safeName(u))
	fmt.Printf("  safeName(nil):   %q\n", safeName(nil))

	// Exercise 4: Pointer to struct vs value
	fmt.Println("\n=== Exercise 4: Pointer vs Value Struct Argument ===")
	r := Rect{Width: 5, Height: 3}
	fmt.Printf("  before: %+v\n", r)
	doubleWidthValue(r)
	fmt.Printf("  after doubleWidthValue (by value): %+v  ← unchanged\n", r)
	doubleWidthPointer(&r)
	fmt.Printf("  after doubleWidthPointer (by ptr):  %+v  ← width doubled\n", r)

	// Exercise 5: Slice append visibility
	fmt.Println("\n=== Exercise 5: Slice Append Visibility ===")
	nums := []int{1, 2, 3}
	fmt.Printf("  original: %v\n", nums)
	appendBroken(nums, 4)
	fmt.Printf("  after appendBroken(4): %v  ← unchanged (slice header was copied)\n", nums)
	nums = appendFixed(nums, 4)
	fmt.Printf("  after appendFixed(4):  %v  ← visible (returned new slice)\n", nums)
}

func swap(a, b *int) {
	*a, *b = *b, *a
}

func safeName(u *User) string {
	if u == nil {
		return "<nil>"
	}
	return u.Name
}

func doubleWidthValue(r Rect) {
	r.Width *= 2
}

func doubleWidthPointer(r *Rect) {
	r.Width *= 2
}

func appendBroken(s []int, v int) {
	s = append(s, v)
	_ = s
}

func appendFixed(s []int, v int) []int {
	return append(s, v)
}
