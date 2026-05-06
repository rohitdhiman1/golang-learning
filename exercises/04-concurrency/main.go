package main

import "fmt"

func main() {
	fmt.Println("========================================")
	fmt.Println("  PHASE 3 — EXERCISES 04: CONCURRENCY")
	fmt.Println("========================================")

	fmt.Println("\n--- Part 1: Goroutines ---")
	fmt.Println()
	Part1Goroutines()

	fmt.Println("\n--- Part 2: Unbuffered Channels ---")
	fmt.Println()
	Part2UnbufferedChannels()

	fmt.Println("\n--- Part 3: Buffered Channels & Select ---")
	fmt.Println()
	Part3BufferedChannelsSelect()

	fmt.Println("\n--- Part 4: Sync Primitives ---")
	fmt.Println()
	Part4SyncPrimitives()

	fmt.Println("\n--- Part 5: Concurrency Patterns ---")
	fmt.Println()
	Part5ConcurrencyPatterns()

	fmt.Println("\n--- Part 6: Context ---")
	fmt.Println()
	Part6Context()
}
