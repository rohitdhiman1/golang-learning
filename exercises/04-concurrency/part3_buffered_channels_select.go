package main

import (
	"fmt"
	"time"
)

func Part3BufferedChannelsSelect() {
	// Exercise 1: Buffered channel basics
	fmt.Println("=== Exercise 1: Buffered Channel ===")
	ch := make(chan int, 3)

	ch <- 10
	ch <- 20
	ch <- 30
	// ch <- 40  // would block — buffer is full and no receiver
	fmt.Printf("  buffer len=%d, cap=%d\n", len(ch), cap(ch))

	fmt.Printf("  received: %d\n", <-ch)
	fmt.Printf("  received: %d\n", <-ch)
	fmt.Printf("  received: %d\n", <-ch)

	// Exercise 2: Select — wait on multiple channels
	fmt.Println("\n=== Exercise 2: Select Statement ===")
	ch1 := make(chan string)
	ch2 := make(chan string)

	go func() {
		time.Sleep(20 * time.Millisecond)
		ch1 <- "from ch1"
	}()
	go func() {
		time.Sleep(10 * time.Millisecond)
		ch2 <- "from ch2"
	}()

	for i := 0; i < 2; i++ {
		select {
		case msg1 := <-ch1:
			fmt.Printf("  %s\n", msg1)
		case msg2 := <-ch2:
			fmt.Printf("  %s\n", msg2)
		}
	}

	// Exercise 3: Select with default (non-blocking)
	fmt.Println("\n=== Exercise 3: Non-Blocking Select ===")
	messages := make(chan string)

	select {
	case msg := <-messages:
		fmt.Printf("  received: %s\n", msg)
	default:
		fmt.Println("  no message ready — default case hit")
	}

	// Exercise 4: Timeout with select
	fmt.Println("\n=== Exercise 4: Timeout ===")
	slow := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		slow <- "slow result"
	}()

	select {
	case result := <-slow:
		fmt.Printf("  got: %s\n", result)
	case <-time.After(50 * time.Millisecond):
		fmt.Println("  timed out after 50ms")
	}

	// Exercise 5: Done channel pattern
	fmt.Println("\n=== Exercise 5: Done Channel Pattern ===")
	done := make(chan struct{})
	data := make(chan int)

	go func() {
		i := 0
		for {
			select {
			case <-done:
				fmt.Println("  generator: received done signal, stopping")
				return
			case data <- i:
				i++
			}
		}
	}()

	for i := 0; i < 5; i++ {
		fmt.Printf("  received: %d\n", <-data)
	}
	close(done)
	time.Sleep(10 * time.Millisecond)
}
