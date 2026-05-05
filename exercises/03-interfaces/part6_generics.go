package main

import (
	"fmt"
	"strconv"
	"strings"
)

func GenericContains[T comparable](s []T, target T) bool {
	for _, v := range s {
		if v == target {
			return true
		}
	}
	return false
}

func GenericMap[T, U any](s []T, f func(T) U) []U {
	result := make([]U, len(s))
	for i, v := range s {
		result[i] = f(v)
	}
	return result
}

func GenericFilter[T any](s []T, pred func(T) bool) []T {
	var result []T
	for _, v := range s {
		if pred(v) {
			result = append(result, v)
		}
	}
	return result
}

type Number interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~float32 | ~float64
}

func GenericSum[T Number](nums []T) T {
	var total T
	for _, n := range nums {
		total += n
	}
	return total
}

type Celsius float64

type Stack[T any] struct {
	items []T
}

func (s *Stack[T]) Push(v T) {
	s.items = append(s.items, v)
}

func (s *Stack[T]) Pop() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	v := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]
	return v, true
}

func (s *Stack[T]) Peek() (T, bool) {
	if len(s.items) == 0 {
		var zero T
		return zero, false
	}
	return s.items[len(s.items)-1], true
}

func (s *Stack[T]) Len() int {
	return len(s.items)
}

func Part6Generics() {
	// Exercise 1: Generic Contains
	fmt.Println("=== Exercise 1: Generic Contains ===")
	ints := []int{1, 2, 3, 4, 5}
	fmt.Printf("  Contains(%v, 3) = %t\n", ints, GenericContains(ints, 3))
	fmt.Printf("  Contains(%v, 9) = %t\n", ints, GenericContains(ints, 9))

	words := []string{"go", "rust", "python"}
	fmt.Printf("  Contains(%v, \"go\") = %t\n", words, GenericContains(words, "go"))
	fmt.Printf("  Contains(%v, \"java\") = %t\n", words, GenericContains(words, "java"))

	// Exercise 2: Generic Map
	fmt.Println("\n=== Exercise 2: Generic Map ===")
	nums := []int{1, 2, 3, 4}
	strs := GenericMap(nums, strconv.Itoa)
	fmt.Printf("  Map(%v, Itoa) = %v\n", nums, strs)

	words2 := []string{"hello", "world", "go"}
	lengths := GenericMap(words2, func(s string) int { return len(s) })
	fmt.Printf("  Map(%v, len) = %v\n", words2, lengths)

	// Exercise 3: Generic Filter
	fmt.Println("\n=== Exercise 3: Generic Filter ===")
	allNums := []int{1, 2, 3, 4, 5, 6, 7, 8}
	evens := GenericFilter(allNums, func(n int) bool { return n%2 == 0 })
	fmt.Printf("  Filter(%v, isEven) = %v\n", allNums, evens)

	fruits := []string{"apple", "banana", "avocado", "cherry", "apricot"}
	aFruits := GenericFilter(fruits, func(s string) bool { return strings.HasPrefix(s, "a") })
	fmt.Printf("  Filter(%v, startsWithA) = %v\n", fruits, aFruits)

	// Exercise 4: Type constraint (Number)
	fmt.Println("\n=== Exercise 4: Number Constraint ===")
	intSlice := []int{10, 20, 30}
	fmt.Printf("  Sum(%v) = %d\n", intSlice, GenericSum(intSlice))

	floatSlice := []float64{1.5, 2.5, 3.0}
	fmt.Printf("  Sum(%v) = %.1f\n", floatSlice, GenericSum(floatSlice))

	temps := []Celsius{20.5, 22.0, 18.3}
	fmt.Printf("  Sum(%v) = %.1f  (Celsius — custom type with ~float64)\n", temps, GenericSum(temps))

	// Exercise 5: Generic Stack
	fmt.Println("\n=== Exercise 5: Generic Stack ===")
	is := Stack[int]{}
	is.Push(10)
	is.Push(20)
	is.Push(30)
	fmt.Printf("  int stack: len=%d\n", is.Len())

	top, ok := is.Peek()
	fmt.Printf("  peek: val=%d, ok=%t\n", top, ok)

	for is.Len() > 0 {
		v, _ := is.Pop()
		fmt.Printf("  pop: %d\n", v)
	}

	_, ok = is.Pop()
	fmt.Printf("  pop from empty: ok=%t\n", ok)

	ss := Stack[string]{}
	ss.Push("hello")
	ss.Push("world")
	v, _ := ss.Pop()
	fmt.Printf("  string stack pop: %q\n", v)
}
