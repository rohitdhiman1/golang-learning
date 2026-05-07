package main

import (
	"fmt"
	"sync"
	"time"
)

func Part4SyncPrimitives() {
	// Exercise 1: Mutex — safe counter
	fmt.Println("=== Exercise 1: Mutex — Safe Counter ===")

	var mu sync.Mutex
	counter := 0

	var wg sync.WaitGroup
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			mu.Lock()
			counter++
			mu.Unlock()
		}()
	}
	wg.Wait()
	fmt.Printf("  counter = %d (expected 1000)\n", counter)

	// Exercise 2: Mutex embedded in struct
	fmt.Println("\n=== Exercise 2: Mutex in Struct ===")
	sc := &SafeCounter{}
	var wg2 sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			sc.Inc()
		}()
	}
	wg2.Wait()
	fmt.Printf("  SafeCounter value = %d (expected 100)\n", sc.Value())

	// Exercise 3: RWMutex — concurrent reads
	fmt.Println("\n=== Exercise 3: RWMutex ===")
	cache := &ReadHeavyCache{data: map[string]string{
		"go":   "fast",
		"java": "verbose",
	}}

	var wg3 sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg3.Add(1)
		go func(id int) {
			defer wg3.Done()
			val := cache.Get("go")
			fmt.Printf("  reader %d: go = %q\n", id, val)
		}(i)
	}

	wg3.Add(1)
	go func() {
		defer wg3.Done()
		cache.Set("python", "dynamic")
		fmt.Println("  writer: set python = dynamic")
	}()

	wg3.Wait()
	fmt.Printf("  python = %q\n", cache.Get("python"))

	// Exercise 4: sync.Once — lazy init
	fmt.Println("\n=== Exercise 4: sync.Once ===")
	var once sync.Once
	var config string

	init := func() {
		fmt.Println("  initialising config... (this prints only once)")
		config = "production"
	}

	var wg4 sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg4.Add(1)
		go func(id int) {
			defer wg4.Done()
			once.Do(init)
			fmt.Printf("  goroutine %d sees config = %q\n", id, config)
		}(i)
	}
	wg4.Wait()

	// Exercise 5: Demonstrating a race condition (without mutex)
	fmt.Println("\n=== Exercise 5: Race Condition Demo ===")
	fmt.Println("  (run with `go run -race .` to detect races)")
	unsafeCounter := 0
	var mu2 sync.Mutex
	var wg5 sync.WaitGroup

	for i := 0; i < 1000; i++ {
		wg5.Add(1)
		go func() {
			defer wg5.Done()
			mu2.Lock()
			unsafeCounter++
			mu2.Unlock()
		}()
	}
	wg5.Wait()
	fmt.Printf("  safe counter = %d (expected 1000)\n", unsafeCounter)
	fmt.Println("  tip: remove mu2.Lock()/Unlock() and run with -race to see the detector catch it")

	time.Sleep(10 * time.Millisecond)
}

type SafeCounter struct {
	mu sync.Mutex
	n  int
}

func (c *SafeCounter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.n++
}

func (c *SafeCounter) Value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.n
}

type ReadHeavyCache struct {
	mu   sync.RWMutex
	data map[string]string
}

func (c *ReadHeavyCache) Get(key string) string {
	c.mu.RLock()
	defer c.mu.RUnlock()
	return c.data[key]
}

func (c *ReadHeavyCache) Set(key, value string) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.data[key] = value
}
