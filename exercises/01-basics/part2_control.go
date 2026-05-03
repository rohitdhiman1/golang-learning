package main

import "fmt"

func Part2Control() {
	// Exercise 1: FizzBuzz
	fmt.Println("=== Exercise 1: FizzBuzz (1–30) ===")
	for i := 1; i <= 30; i++ {
		switch {
		case i%15 == 0:
			fmt.Print("FizzBuzz ")
		case i%3 == 0:
			fmt.Print("Fizz ")
		case i%5 == 0:
			fmt.Print("Buzz ")
		default:
			fmt.Printf("%d ", i)
		}
	}
	fmt.Println()

	// Exercise 2: Infinite loop with break
	fmt.Println("\n=== Exercise 2: Infinite Loop with Break ===")
	val := 1
	for {
		val *= 2
		if val > 1000 {
			break
		}
	}
	fmt.Printf("Final value (first power of 2 > 1000): %d\n", val)

	// Exercise 3: Range over string
	fmt.Println("\n=== Exercise 3: Range over String ===")
	s := "héllo, 世界"
	fmt.Printf("String: %q (len in bytes: %d)\n", s, len(s))
	for i, r := range s {
		fmt.Printf("  index=%2d  rune=%c  (U+%04X)\n", i, r, r)
	}

	// Exercise 4: Switch — grade classifier
	fmt.Println("\n=== Exercise 4: Grade Classifier ===")
	scores := []int{95, 82, 73, 65, 45}
	for _, sc := range scores {
		grade := classifyGrade(sc)
		fmt.Printf("  Score %3d → Grade %s\n", sc, grade)
	}

	// Exercise 5: Labeled break
	fmt.Println("\n=== Exercise 5: Labeled Break ===")
	iterations := 0
outer:
	for i := 0; i < 4; i++ {
		for j := 0; j < 4; j++ {
			if i+j == 5 {
				fmt.Printf("  Breaking at i=%d, j=%d (i+j=5)\n", i, j)
				break outer
			}
			iterations++
		}
	}
	fmt.Printf("  Completed %d iterations before break\n", iterations)

	// Exercise 6: Defer order (LIFO)
	fmt.Println("\n=== Exercise 6: Defer Order (LIFO) ===")
	fmt.Println("  Calling deferDemo()...")
	deferDemo()
}

func classifyGrade(score int) string {
	switch {
	case score >= 90:
		return "A"
	case score >= 80:
		return "B"
	case score >= 70:
		return "C"
	case score >= 60:
		return "D"
	default:
		return "F"
	}
}

func deferDemo() {
	defer fmt.Println("  deferred: a (deferred first, runs last)")
	defer fmt.Println("  deferred: b (deferred second)")
	defer fmt.Println("  deferred: c (deferred third, runs first)")
	fmt.Println("  deferDemo body executes")
}
