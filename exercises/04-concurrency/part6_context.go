package main

import (
	"context"
	"fmt"
	"time"
)

func Part6Context() {
	// Exercise 1: context.WithCancel
	fmt.Println("=== Exercise 1: WithCancel ===")
	ctx, cancel := context.WithCancel(context.Background())

	go func(ctx context.Context) {
		for i := 0; ; i++ {
			select {
			case <-ctx.Done():
				fmt.Printf("  generator stopped: %v\n", ctx.Err())
				return
			default:
				if i < 5 {
					fmt.Printf("  generating: %d\n", i)
				}
				time.Sleep(5 * time.Millisecond)
			}
		}
	}(ctx)

	time.Sleep(30 * time.Millisecond)
	cancel()
	time.Sleep(10 * time.Millisecond)

	// Exercise 2: context.WithTimeout
	fmt.Println("\n=== Exercise 2: WithTimeout ===")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 50*time.Millisecond)
	defer cancel2()

	result := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		result <- "operation complete"
	}()

	select {
	case res := <-result:
		fmt.Printf("  got: %s\n", res)
	case <-ctx2.Done():
		fmt.Printf("  timed out: %v\n", ctx2.Err())
	}

	// Exercise 3: context.WithValue (typed keys)
	fmt.Println("\n=== Exercise 3: WithValue ===")
	type contextKey string
	const requestIDKey contextKey = "requestID"

	ctx3 := context.WithValue(context.Background(), requestIDKey, "req-abc-123")
	processRequest(ctx3, requestIDKey)

	// Exercise 4: Context propagation — parent cancels children
	fmt.Println("\n=== Exercise 4: Propagation ===")
	parent, parentCancel := context.WithCancel(context.Background())

	child1, child1Cancel := context.WithCancel(parent)
	defer child1Cancel()

	child2, child2Cancel := context.WithTimeout(parent, 5*time.Second)
	defer child2Cancel()

	parentCancel()
	time.Sleep(10 * time.Millisecond)

	fmt.Printf("  parent err: %v\n", parent.Err())
	fmt.Printf("  child1 err: %v\n", child1.Err())
	fmt.Printf("  child2 err: %v\n", child2.Err())
	fmt.Println("  (cancelling parent cancelled both children)")

	// Exercise 5: Simulated graceful shutdown
	fmt.Println("\n=== Exercise 5: Graceful Shutdown ===")
	ctx5, shutdown := context.WithCancel(context.Background())

	done := make(chan struct{})
	go func() {
		defer close(done)
		server(ctx5)
	}()

	time.Sleep(60 * time.Millisecond)
	fmt.Println("  main: sending shutdown signal...")
	shutdown()
	<-done
	fmt.Println("  main: server stopped cleanly")
}

func processRequest(ctx context.Context, key interface{}) {
	if reqID, ok := ctx.Value(key).(string); ok {
		fmt.Printf("  handling request: %s\n", reqID)
	} else {
		fmt.Println("  no request ID found")
	}
}

func server(ctx context.Context) {
	ticker := time.NewTicker(20 * time.Millisecond)
	defer ticker.Stop()

	for i := 1; ; i++ {
		select {
		case <-ctx.Done():
			fmt.Printf("  server: shutting down after %d ticks (%v)\n", i-1, ctx.Err())
			return
		case <-ticker.C:
			fmt.Printf("  server: tick %d\n", i)
		}
	}
}
