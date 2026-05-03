package main

import (
	"fmt"
	"strconv"
)

func Part1Variables() {
	// Exercise 1: Hello, World
	fmt.Println("=== Exercise 1: Hello, World ===")
	fmt.Println("Hello, World!")

	// Exercise 2: Variable declaration — all three methods
	fmt.Println("\n=== Exercise 2: Variable Declaration ===")

	var name string = "Rohit"
	fmt.Printf("var with type:       name = %v (type: %T)\n", name, name)

	var age = 25
	fmt.Printf("var with inference:  age  = %v (type: %T)\n", age, age)

	score := 99.5
	fmt.Printf("short declaration:   score = %v (type: %T)\n", score, score)

	// Exercise 3: Zero values
	fmt.Println("\n=== Exercise 3: Zero Values ===")

	var zeroInt int
	var zeroFloat float64
	var zeroBool bool
	var zeroString string

	fmt.Printf("int:     %v (zero? %t)\n", zeroInt, zeroInt == 0)
	fmt.Printf("float64: %v (zero? %t)\n", zeroFloat, zeroFloat == 0)
	fmt.Printf("bool:    %v (zero? %t)\n", zeroBool, zeroBool == false)
	fmt.Printf("string:  %q (zero? %t)\n", zeroString, zeroString == "")

	// Exercise 4: Constants & iota
	fmt.Println("\n=== Exercise 4: Constants & iota ===")

	const (
		Sunday = iota
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)
	fmt.Printf("Wednesday = %d\n", Wednesday)

	// Exercise 5: Type conversion
	fmt.Println("\n=== Exercise 5: Type Conversion ===")

	var x int = 42

	asFloat := float64(x)
	fmt.Printf("int → float64: %v (type: %T)\n", asFloat, asFloat)

	asString := strconv.Itoa(x)
	fmt.Printf("int → string (strconv.Itoa): %q (type: %T)\n", asString, asString)

	asBadString := string(rune(x))
	fmt.Printf("int → string (cast):         %q ← this is the rune '*', not \"42\"!\n", asBadString)

	// Exercise 6: Multiple assignment & swap
	fmt.Println("\n=== Exercise 6: Multiple Assignment & Swap ===")

	a, b := 10, 20
	fmt.Printf("Before swap: a=%d, b=%d\n", a, b)
	a, b = b, a
	fmt.Printf("After swap:  a=%d, b=%d\n", a, b)
}
