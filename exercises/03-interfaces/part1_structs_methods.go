package main

import (
	"fmt"
	"math"
)

type Rect struct {
	Width, Height float64
}

func (r Rect) Area() float64 {
	return r.Width * r.Height
}

func (r Rect) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

func (r *Rect) Scale(factor float64) {
	r.Width *= factor
	r.Height *= factor
}

type Circle struct {
	Radius float64
}

func NewCircle(r float64) (*Circle, error) {
	if r <= 0 {
		return nil, fmt.Errorf("radius must be positive, got %f", r)
	}
	return &Circle{Radius: r}, nil
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

type Animal struct {
	Name string
}

func (a Animal) Speak() string {
	return "..."
}

type Dog struct {
	Animal
	Breed string
}

func (d Dog) Speak() string {
	return "Woof!"
}

type Resizer interface {
	Scale(float64)
}

type Counter struct {
	n int
}

func (c *Counter) Increment() {
	c.n++
}

func (c Counter) Value() int {
	return c.n
}

func Part1StructsMethods() {
	// Exercise 1: Value vs pointer receiver
	fmt.Println("=== Exercise 1: Value vs Pointer Receiver ===")
	r := Rect{Width: 10, Height: 5}
	fmt.Printf("  before scale: %+v, area=%.1f\n", r, r.Area())
	r.Scale(2)
	fmt.Printf("  after scale(2): %+v, area=%.1f\n", r, r.Area())

	// Exercise 2: Constructor pattern
	fmt.Println("\n=== Exercise 2: Constructor Pattern ===")
	c, err := NewCircle(5)
	if err != nil {
		fmt.Printf("  error: %v\n", err)
	} else {
		fmt.Printf("  circle: radius=%.1f, area=%.2f, perimeter=%.2f\n", c.Radius, c.Area(), c.Perimeter())
	}

	_, err = NewCircle(-1)
	fmt.Printf("  NewCircle(-1): err=%v\n", err)

	// Exercise 3: Embedding
	fmt.Println("\n=== Exercise 3: Embedding ===")
	d := Dog{
		Animal: Animal{Name: "Rex"},
		Breed:  "German Shepherd",
	}
	fmt.Printf("  d.Name = %q (promoted from Animal)\n", d.Name)
	fmt.Printf("  d.Speak() = %q (Dog's own method — shadows Animal)\n", d.Speak())
	fmt.Printf("  d.Animal.Speak() = %q (explicitly access Animal's method)\n", d.Animal.Speak())

	// Exercise 4: Method set demo
	fmt.Println("\n=== Exercise 4: Method Set Demo ===")
	rect := Rect{Width: 3, Height: 4}
	// var _ Resizer = rect   // WON'T COMPILE: Rect doesn't have Scale (pointer receiver)
	var _ Resizer = &rect // OK: *Rect has Scale
	fmt.Println("  var _ Resizer = rect   → compile error (Scale has pointer receiver)")
	fmt.Println("  var _ Resizer = &rect  → OK (*Rect method set includes Scale)")

	// Exercise 5: Zero-value usefulness
	fmt.Println("\n=== Exercise 5: Zero-Value Usefulness ===")
	var ctr Counter
	fmt.Printf("  zero value: Counter{n: %d}\n", ctr.Value())
	ctr.Increment()
	ctr.Increment()
	ctr.Increment()
	fmt.Printf("  after 3 increments: Counter{n: %d}\n", ctr.Value())
}
