package main

import (
	"fmt"
	"sync"
	"time"
)

func Part5ConcurrencyPatterns() {
	// Exercise 1: Pipeline pattern
	fmt.Println("=== Exercise 1: Pipeline (generate → double → print) ===")

	nums := generateNums(1, 2, 3, 4, 5)
	doubled := double(nums)

	for v := range doubled {
		fmt.Printf("  %d\n", v)
	}

	// Exercise 2: Fan-out — multiple workers reading from one channel
	fmt.Println("\n=== Exercise 2: Fan-Out ===")
	jobs := make(chan int, 10)
	var wg sync.WaitGroup

	for w := 1; w <= 3; w++ {
		wg.Add(1)
		go func(id int) {
			defer wg.Done()
			for j := range jobs {
				fmt.Printf("  worker %d processed job %d\n", id, j)
				time.Sleep(5 * time.Millisecond)
			}
		}(w)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}
	close(jobs)
	wg.Wait()

	// Exercise 3: Fan-in — merge multiple channels into one
	fmt.Println("\n=== Exercise 3: Fan-In (Merge) ===")
	c1 := generateNums(1, 3, 5)
	c2 := generateNums(2, 4, 6)
	merged := mergeChannels(c1, c2)

	for v := range merged {
		fmt.Printf("  %d\n", v)
	}

	// Exercise 4: Worker pool
	fmt.Println("\n=== Exercise 4: Worker Pool ===")
	const numJobs = 8
	const numWorkers = 3

	jobsCh := make(chan int, numJobs)
	resultsCh := make(chan string, numJobs)

	var wg2 sync.WaitGroup
	for w := 1; w <= numWorkers; w++ {
		wg2.Add(1)
		go func(id int) {
			defer wg2.Done()
			for j := range jobsCh {
				time.Sleep(10 * time.Millisecond)
				resultsCh <- fmt.Sprintf("worker %d finished job %d", id, j)
			}
		}(w)
	}

	for j := 1; j <= numJobs; j++ {
		jobsCh <- j
	}
	close(jobsCh)

	go func() {
		wg2.Wait()
		close(resultsCh)
	}()

	for r := range resultsCh {
		fmt.Printf("  %s\n", r)
	}

	// Exercise 5: Pipeline with done channel (cancellable)
	fmt.Println("\n=== Exercise 5: Cancellable Pipeline ===")
	done := make(chan struct{})

	infinite := func(done <-chan struct{}) <-chan int {
		out := make(chan int)
		go func() {
			defer close(out)
			i := 0
			for {
				select {
				case <-done:
					return
				case out <- i:
					i++
				}
			}
		}()
		return out
	}

	stream := infinite(done)
	for i := 0; i < 5; i++ {
		fmt.Printf("  %d\n", <-stream)
	}
	close(done)
	time.Sleep(10 * time.Millisecond)
	fmt.Println("  pipeline cancelled — no goroutine leak")
}

func generateNums(nums ...int) <-chan int {
	out := make(chan int)
	go func() {
		for _, n := range nums {
			out <- n
		}
		close(out)
	}()
	return out
}

func double(in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		for n := range in {
			out <- n * 2
		}
		close(out)
	}()
	return out
}

func mergeChannels(channels ...<-chan int) <-chan int {
	var wg sync.WaitGroup
	merged := make(chan int)

	for _, ch := range channels {
		wg.Add(1)
		go func(c <-chan int) {
			defer wg.Done()
			for v := range c {
				merged <- v
			}
		}(ch)
	}

	go func() {
		wg.Wait()
		close(merged)
	}()

	return merged
}
