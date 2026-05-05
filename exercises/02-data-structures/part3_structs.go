package main

import (
	"encoding/json"
	"fmt"
)

type Student struct {
	Name  string
	Grade int
	Score float64
}

type Point struct {
	X, Y int
}

type RGB struct {
	R, G, B uint8
}

type Person struct {
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email,omitempty"`
}

func Part3Structs() {
	// Exercise 1: Student roster
	fmt.Println("=== Exercise 1: Student Roster ===")
	students := []Student{
		{"Alice", 10, 92.5},
		{"Bob", 10, 88.0},
		{"Charlie", 11, 95.3},
		{"Diana", 11, 91.0},
		{"Eve", 10, 87.5},
	}
	best := students[0]
	for _, s := range students[1:] {
		if s.Score > best.Score {
			best = s
		}
	}
	fmt.Printf("  students: %v\n", students)
	fmt.Printf("  highest score: %s (%.1f)\n", best.Name, best.Score)

	// Exercise 2: Value vs pointer copy
	fmt.Println("\n=== Exercise 2: Value vs Pointer Copy ===")

	p1 := Point{1, 2}
	p2 := p1
	p2.X = 99
	fmt.Printf("  value copy: p1=%v, p2=%v  ← p1 unchanged\n", p1, p2)

	p3 := &Point{1, 2}
	p4 := p3
	p4.X = 99
	fmt.Printf("  pointer copy: p3=%v, p4=%v  ← both changed (same underlying data)\n", *p3, *p4)

	// Exercise 3: Anonymous struct
	fmt.Println("\n=== Exercise 3: Anonymous Struct ===")
	resp := struct {
		Status  int
		Body    string
		Headers map[string]string
	}{
		Status: 200,
		Body:   `{"ok": true}`,
		Headers: map[string]string{
			"Content-Type": "application/json",
			"X-Request-Id": "abc-123",
		},
	}
	fmt.Printf("  %+v\n", resp)

	// Exercise 4: Struct comparison
	fmt.Println("\n=== Exercise 4: Struct Comparison ===")
	c1 := RGB{255, 128, 0}
	c2 := RGB{255, 128, 0}
	c3 := RGB{0, 0, 0}
	fmt.Printf("  c1=%v, c2=%v → c1==c2? %t\n", c1, c2, c1 == c2)
	fmt.Printf("  c1=%v, c3=%v → c1==c3? %t\n", c1, c3, c1 == c3)
	fmt.Println("  NOTE: adding a []string field would make RGB non-comparable (won't compile with ==)")

	// Exercise 5: JSON marshal/unmarshal
	fmt.Println("\n=== Exercise 5: JSON Marshal/Unmarshal ===")

	alice := Person{Name: "Alice", Age: 30, Email: "alice@example.com"}
	data, _ := json.Marshal(alice)
	fmt.Printf("  marshal (with email):    %s\n", data)

	bob := Person{Name: "Bob", Age: 25}
	data, _ = json.Marshal(bob)
	fmt.Printf("  marshal (empty email):   %s  ← omitempty drops it\n", data)

	var decoded Person
	_ = json.Unmarshal([]byte(`{"name":"Charlie","age":35}`), &decoded)
	fmt.Printf("  unmarshal: %+v  ← Email is zero value %q\n", decoded, decoded.Email)
}
