package main

import (
	"fmt"
	"io"
	"sort"
	"strings"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

func printShape(s Shape) {
	fmt.Printf("    type=%-10T  area=%.2f  perimeter=%.2f\n", s, s.Area(), s.Perimeter())
}

type Colour struct {
	R, G, B uint8
}

func (c Colour) String() string {
	return fmt.Sprintf("#%02X%02X%02X", c.R, c.G, c.B)
}

type Person struct {
	Name string
	Age  int
}

type ByAge []Person

func (a ByAge) Len() int           { return len(a) }
func (a ByAge) Less(i, j int) bool { return a[i].Age < a[j].Age }
func (a ByAge) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type Buffer struct {
	data []byte
	pos  int
}

func (b *Buffer) Read(p []byte) (int, error) {
	if b.pos >= len(b.data) {
		return 0, io.EOF
	}
	n := copy(p, b.data[b.pos:])
	b.pos += n
	return n, nil
}

func (b *Buffer) Write(p []byte) (int, error) {
	b.data = append(b.data, p...)
	return len(p), nil
}

func summarise(r io.Reader) (string, error) {
	data, err := io.ReadAll(r)
	if err != nil {
		return "", fmt.Errorf("summarise: %w", err)
	}
	return string(data), nil
}

func Part3Interfaces() {
	// Exercise 1: Shape interface
	fmt.Println("=== Exercise 1: Shape Interface ===")
	shapes := []Shape{
		Rect{Width: 10, Height: 5},
		&Circle{Radius: 7},
	}
	for _, s := range shapes {
		printShape(s)
	}

	// Exercise 2: Stringer
	fmt.Println("\n=== Exercise 2: fmt.Stringer ===")
	red := Colour{255, 0, 0}
	green := Colour{0, 255, 0}
	white := Colour{255, 255, 255}
	fmt.Printf("  red:   %s\n", red)
	fmt.Printf("  green: %s\n", green)
	fmt.Printf("  white: %s\n", white)

	// Exercise 3: Sort interface
	fmt.Println("\n=== Exercise 3: sort.Interface ===")
	people := []Person{
		{"Charlie", 35},
		{"Alice", 25},
		{"Bob", 30},
		{"Diana", 20},
	}
	fmt.Printf("  before: %v\n", people)
	sort.Sort(ByAge(people))
	fmt.Printf("  after:  %v\n", people)

	// Exercise 4: Interface composition (ReadWriter)
	fmt.Println("\n=== Exercise 4: Interface Composition (ReadWriter) ===")
	var rw io.ReadWriter = &Buffer{}
	rw.Write([]byte("hello from Buffer"))
	buf := make([]byte, 64)
	n, _ := rw.Read(buf)
	fmt.Printf("  wrote then read: %q\n", string(buf[:n]))

	// Exercise 5: Accept interfaces, return structs
	fmt.Println("\n=== Exercise 5: Accept Interfaces, Return Structs ===")
	content, err := summarise(strings.NewReader("Go interfaces are powerful"))
	if err != nil {
		fmt.Printf("  error: %v\n", err)
	} else {
		fmt.Printf("  summarise(reader) = %q\n", content)
	}
}
