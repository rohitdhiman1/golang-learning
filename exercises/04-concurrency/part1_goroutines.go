package main

import (
	"fmt"
	"sync"
	"time"
)

func Part1Goroutines() {
	// Exercise 1: Basic goroutine
	fmt.Println("=== Exercise 1: Basic Goroutine ===")
	go func() {
		fmt.Println("  hello from goroutine!")
	}()
	time.Sleep(10 * time.Millisecond)

	// Exercise 2: WaitGroup — wait for goroutines to finish
	fmt.Println("\n=== Exercise 2: WaitGroup ===")
	var wg sync.WaitGroup

	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			fmt.Printf("  worker %d: started\n", id)
			time.Sleep(time.Duration(id) * 10 * time.Millisecond)
			fmt.Printf("  worker %d: done\n", id)
		}(i)
	}

	wg.Wait()
	fmt.Println("  all workers finished")

	// Exercise 3: Closure variable capture gotcha
	fmt.Println("\n=== Exercise 3: Closure Variable Capture ===")

	fmt.Println("  WRONG way (all goroutines see final value of i):")
	var wg2 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			fmt.Printf("    i = %d\n", i)
		}()
	}
	wg2.Wait()

	fmt.Println("  RIGHT way (pass i as parameter):")
	var wg3 sync.WaitGroup
	for i := 0; i < 3; i++ {
		wg3.Add(1)
		go func(n int) {
			defer wg3.Done()
			fmt.Printf("    n = %d\n", n)
		}(i)
	}
	wg3.Wait()

	// Exercise 4: Goroutines are concurrent, not necessarily parallel
	fmt.Println("\n=== Exercise 4: Concurrent Printing ===")
	var wg4 sync.WaitGroup

	for _, word := range []string{"hello", "world", "from", "goroutines"} {
		wg4.Add(1)
		go func(w string) {
			defer wg4.Done()
			fmt.Printf("  %s\n", w)
		}(word)
	}
	wg4.Wait()
	fmt.Println("  (order may vary between runs — that's concurrency!)")

	// Exercise 5: Named function as goroutine
	fmt.Println("\n=== Exercise 5: Named Function as Goroutine ===")
	var wg5 sync.WaitGroup
	wg5.Add(1)
	go greet(&wg5, "Rohit")
	wg5.Wait()
}

func greet(wg *sync.WaitGroup, name string) {
	defer wg.Done()
	fmt.Printf("  hello, %s! (from a named goroutine)\n", name)
}
