package main

import (
	"fmt"
	"time"
)

func Part2UnbufferedChannels() {
	// Exercise 1: Basic send and receive
	fmt.Println("=== Exercise 1: Basic Send & Receive ===")
	ch := make(chan string)

	go func() {
		ch <- "hello from channel!"
	}()

	msg := <-ch
	fmt.Printf("  received: %s\n", msg)

	// Exercise 2: Channel as synchronisation
	fmt.Println("\n=== Exercise 2: Channel as Synchronisation ===")
	done := make(chan bool)

	go func() {
		fmt.Println("  doing work...")
		time.Sleep(20 * time.Millisecond)
		fmt.Println("  work complete")
		done <- true
	}()

	<-done
	fmt.Println("  main: received done signal")

	// Exercise 3: Passing data between goroutines
	fmt.Println("\n=== Exercise 3: Producer → Consumer ===")
	numbers := make(chan int)

	go func() {
		for i := 1; i <= 5; i++ {
			numbers <- i
		}
		close(numbers)
	}()

	for n := range numbers {
		fmt.Printf("  received: %d\n", n)
	}

	// Exercise 4: Directional channels in function signatures
	fmt.Println("\n=== Exercise 4: Directional Channels ===")
	ch2 := make(chan int)

	go produce(ch2)
	consume(ch2)

	// Exercise 5: Multiple goroutines, one channel
	fmt.Println("\n=== Exercise 5: Multiple Senders, One Receiver ===")
	results := make(chan string)

	for i := 1; i <= 3; i++ {
		go func(id int) {
			time.Sleep(time.Duration(id) * 10 * time.Millisecond)
			results <- fmt.Sprintf("result from worker %d", id)
		}(i)
	}

	for i := 0; i < 3; i++ {
		fmt.Printf("  %s\n", <-results)
	}
}

func produce(out chan<- int) {
	for i := 10; i <= 13; i++ {
		out <- i
	}
	close(out)
}

func consume(in <-chan int) {
	for v := range in {
		fmt.Printf("  consumed: %d\n", v)
	}
}
