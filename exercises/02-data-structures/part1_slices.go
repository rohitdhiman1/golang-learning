package main

import (
	"fmt"
	"sort"
	"strings"
)

func Part1Slices() {
	// Exercise 1: Basic operations
	fmt.Println("=== Exercise 1: Basic Operations ===")
	s := []int{5, 3, 1, 4, 2}
	fmt.Printf("  slice: %v  len=%d  cap=%d\n", s, len(s), cap(s))
	s = append(s, 6, 7)
	fmt.Printf("  after append(6,7): %v  len=%d  cap=%d\n", s, len(s), cap(s))

	// Exercise 2: Slice expressions
	fmt.Println("\n=== Exercise 2: Slice Expressions ===")
	s2 := []int{0, 1, 2, 3, 4, 5, 6}
	fmt.Printf("  original:    %v\n", s2)
	fmt.Printf("  first 3:     %v  (s[:3])\n", s2[:3])
	fmt.Printf("  last 3:      %v  (s[len-3:])\n", s2[len(s2)-3:])
	fmt.Printf("  elements 2–4:%v  (s[2:5])\n", s2[2:5])

	// Exercise 3: Shared backing array
	fmt.Println("\n=== Exercise 3: Shared Backing Array ===")
	a := []int{1, 2, 3, 4, 5}
	b := a[1:3]
	fmt.Printf("  a = %v\n  b = a[1:3] = %v\n", a, b)

	b[0] = 99
	fmt.Printf("  after b[0] = 99:\n")
	fmt.Printf("  a = %v  ← a[1] changed!\n", a)
	fmt.Printf("  b = %v\n", b)

	a[1] = 2
	c := make([]int, len(b))
	copy(c, a[1:3])
	c[0] = 99
	fmt.Printf("  fix with copy: c = %v, a = %v  ← a unchanged\n", c, a)

	// Exercise 4: Word counter
	fmt.Println("\n=== Exercise 4: Word Counter ===")
	sentence := "the quick brown fox jumps over the lazy dog the"
	unique := wordCounter(sentence)
	fmt.Printf("  input:  %q\n", sentence)
	fmt.Printf("  unique words: %v, count: %d\n", unique, len(unique))

	// Exercise 5: Dedup in-place
	fmt.Println("\n=== Exercise 5: Dedup In-Place ===")
	sorted := []int{1, 1, 2, 3, 3, 3, 4, 5, 5}
	fmt.Printf("  input:  %v\n", sorted)
	n := dedupInPlace(sorted)
	fmt.Printf("  result: %v (length: %d)\n", sorted[:n], n)

	// Exercise 6: Rotate left
	fmt.Println("\n=== Exercise 6: Rotate Left ===")
	r := []int{1, 2, 3, 4, 5}
	fmt.Printf("  input:  %v, k=2\n", r)
	rotateLeft(r, 2)
	fmt.Printf("  result: %v\n", r)
}

func wordCounter(sentence string) []string {
	words := strings.Fields(sentence)
	seen := make(map[string]bool)
	var unique []string
	for _, w := range words {
		if !seen[w] {
			seen[w] = true
			unique = append(unique, w)
		}
	}
	sort.Strings(unique)
	return unique
}

func dedupInPlace(s []int) int {
	if len(s) == 0 {
		return 0
	}
	w := 1
	for r := 1; r < len(s); r++ {
		if s[r] != s[r-1] {
			s[w] = s[r]
			w++
		}
	}
	return w
}

func rotateLeft(s []int, k int) {
	k = k % len(s)
	reverse(s[:k])
	reverse(s[k:])
	reverse(s)
}

func reverse(s []int) {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
}
