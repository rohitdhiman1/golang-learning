package main

import (
	"errors"
	"fmt"
)

func Part3Functions() {
	// Exercise 1: Multiple return
	fmt.Println("=== Exercise 1: Multiple Return (divide) ===")
	result, err := divide(10, 3)
	fmt.Printf("  divide(10, 3) = %.4f, err = %v\n", result, err)

	result, err = divide(10, 0)
	fmt.Printf("  divide(10, 0) = %.4f, err = %v\n", result, err)

	// Exercise 2: Named return
	fmt.Println("\n=== Exercise 2: Named Return (minMax) ===")
	nums := []int{5, 1, 9, 3, 7}
	lo, hi := minMax(nums)
	fmt.Printf("  minMax(%v) → min=%d, max=%d\n", nums, lo, hi)

	// Exercise 3: Variadic sum
	fmt.Println("\n=== Exercise 3: Variadic Sum ===")
	fmt.Printf("  sum(1, 2, 3)       = %d\n", sum(1, 2, 3))
	fmt.Printf("  sum(10, 20, 30, 40)= %d\n", sum(10, 20, 30, 40))

	vals := []int{100, 200, 300}
	fmt.Printf("  sum(vals...)       = %d  (spread a slice)\n", sum(vals...))

	// Exercise 4: First-class function
	fmt.Println("\n=== Exercise 4: First-class Function (applyAll) ===")
	input := []int{1, 2, 3, 4, 5}
	doubled := applyAll(input, func(n int) int { return n * 2 })
	fmt.Printf("  applyAll(%v, double) = %v\n", input, doubled)

	squared := applyAll(input, func(n int) int { return n * n })
	fmt.Printf("  applyAll(%v, square) = %v\n", input, squared)

	// Exercise 5: Closure counter
	fmt.Println("\n=== Exercise 5: Closure Counter ===")
	counterA := counter()
	counterB := counter()
	fmt.Printf("  counterA: %d, %d, %d\n", counterA(), counterA(), counterA())
	fmt.Printf("  counterB: %d, %d\n", counterB(), counterB())
	fmt.Printf("  counterA: %d  ← independent, continues from 3\n", counterA())

	// Exercise 6: Closure loop gotcha
	fmt.Println("\n=== Exercise 6: Closure Loop Gotcha ===")
	closureLoopGotcha()

	// Exercise 7: Defer in a loop
	fmt.Println("\n=== Exercise 7: Defer in a Loop ===")
	deferInLoop()
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("division by zero")
	}
	return a / b, nil
}

func minMax(nums []int) (min, max int) {
	min = nums[0]
	max = nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return
}

func sum(nums ...int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func applyAll(nums []int, fn func(int) int) []int {
	result := make([]int, len(nums))
	for i, n := range nums {
		result[i] = fn(n)
	}
	return result
}

func counter() func() int {
	count := 0
	return func() int {
		count++
		return count
	}
}

func closureLoopGotcha() {
	// Go 1.22+ changed loop variable semantics: each iteration gets its own copy.
	// In Go <1.22, the following would print "5 5 5 5 5" because all closures
	// captured the same variable. Now it prints "0 1 2 3 4" correctly.
	// The fixes below are still important to know for interviews and pre-1.22 code.

	fmt.Println("  Pre-1.22 BUG (now fixed by language spec — loop vars are per-iteration):")
	funcs := make([]func(), 5)
	for i := 0; i < 5; i++ {
		funcs[i] = func() {
			fmt.Printf("    i=%d", i)
		}
	}
	fmt.Print("  ")
	for _, f := range funcs {
		f()
	}
	fmt.Println("  ← Go 1.22+: correct (per-iteration var). Pre-1.22: all would print 5")

	fmt.Println("  FIX 1 (pre-1.22) — shadow with i := i (still valid, just redundant now):")
	for i := 0; i < 5; i++ {
		i := i
		funcs[i] = func() {
			fmt.Printf("    i=%d", i)
		}
	}
	fmt.Print("  ")
	for _, f := range funcs {
		f()
	}
	fmt.Println("  ← 0,1,2,3,4")

	fmt.Println("  FIX 2 (pre-1.22) — pass as function argument:")
	for i := 0; i < 5; i++ {
		funcs[i] = func(n int) func() {
			return func() {
				fmt.Printf("    i=%d", n)
			}
		}(i)
	}
	fmt.Print("  ")
	for _, f := range funcs {
		f()
	}
	fmt.Println("  ← 0,1,2,3,4")
}

func deferInLoop() {
	fmt.Println("  WRONG: defer in a loop (defers pile up until function returns)")
	fmt.Println("  Simulating 3 file opens with defer close in loop:")
	wrongDeferLoop()

	fmt.Println("\n  CORRECT: wrap in IIFE so defer runs each iteration")
	correctDeferLoop()
}

func wrongDeferLoop() {
	for i := 0; i < 3; i++ {
		fmt.Printf("    opening file_%d\n", i)
		defer fmt.Printf("    closing file_%d (deferred)\n", i)
	}
	fmt.Println("    — function body done, now defers run in LIFO:")
}

func correctDeferLoop() {
	for i := 0; i < 3; i++ {
		func(n int) {
			fmt.Printf("    opening file_%d\n", n)
			defer fmt.Printf("    closing file_%d (deferred — runs immediately via IIFE)\n", n)
		}(i)
	}
	fmt.Println("    — each file closed right after opening")
}
