package main

import (
	"fmt"
	"sort"
	"strings"
)

func Part2Maps() {
	// Exercise 1: Word frequency
	fmt.Println("=== Exercise 1: Word Frequency ===")
	input := "go is great go is fast go"
	freq := wordFrequency(input)
	keys := make([]string, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	fmt.Printf("  input: %q\n  ", input)
	for _, k := range keys {
		fmt.Printf("%s:%d ", k, freq[k])
	}
	fmt.Println()

	// Exercise 2: Grouping by first letter
	fmt.Println("\n=== Exercise 2: Grouping by First Letter ===")
	fruits := []string{"apple", "avocado", "banana", "blueberry", "cherry", "cranberry"}
	groups := groupByFirstLetter(fruits)
	sortedLetters := make([]string, 0, len(groups))
	for k := range groups {
		sortedLetters = append(sortedLetters, k)
	}
	sort.Strings(sortedLetters)
	for _, letter := range sortedLetters {
		fmt.Printf("  %s: %v\n", letter, groups[letter])
	}

	// Exercise 3: Two-sum
	fmt.Println("\n=== Exercise 3: Two-Sum ===")
	nums := []int{2, 7, 11, 15}
	target := 9
	i, j := twoSum(nums, target)
	fmt.Printf("  input: %v, target=%d\n", nums, target)
	fmt.Printf("  result: [%d, %d]  (%d + %d = %d)\n", i, j, nums[i], nums[j], target)

	// Exercise 4: Inverted index
	fmt.Println("\n=== Exercise 4: Inverted Index ===")
	sentences := []string{"go is fast", "go is simple", "python is dynamic"}
	index := invertedIndex(sentences)
	fmt.Printf("  sentences: %v\n", sentences)
	fmt.Printf("  \"is\" appears in sentences: %v\n", index["is"])
	fmt.Printf("  \"go\" appears in sentences: %v\n", index["go"])
	fmt.Printf("  \"python\" appears in sentences: %v\n", index["python"])

	// Exercise 5: Default values & two-value lookup
	fmt.Println("\n=== Exercise 5: Default Values ===")
	words := []string{"hello", "world", "go"}
	lengths := wordLengths(words)
	fmt.Printf("  wordLengths(%v) = %v\n", words, lengths)

	val := lengths["missing"]
	fmt.Printf("  lengths[\"missing\"] = %d  ← zero value for int\n", val)

	val, ok := lengths["missing"]
	fmt.Printf("  two-value: val=%d, ok=%t  ← key does not exist\n", val, ok)

	val, ok = lengths["hello"]
	fmt.Printf("  two-value: val=%d, ok=%t  ← key exists\n", val, ok)
}

func wordFrequency(s string) map[string]int {
	freq := make(map[string]int)
	for _, w := range strings.Fields(s) {
		freq[w]++
	}
	return freq
}

func groupByFirstLetter(words []string) map[string][]string {
	groups := make(map[string][]string)
	for _, w := range words {
		key := string(w[0])
		groups[key] = append(groups[key], w)
	}
	return groups
}

func twoSum(nums []int, target int) (int, int) {
	seen := make(map[int]int)
	for i, n := range nums {
		if j, ok := seen[target-n]; ok {
			return j, i
		}
		seen[n] = i
	}
	return -1, -1
}

func invertedIndex(sentences []string) map[string][]int {
	index := make(map[string][]int)
	for i, s := range sentences {
		for _, w := range strings.Fields(s) {
			index[w] = append(index[w], i)
		}
	}
	return index
}

func wordLengths(words []string) map[string]int {
	m := make(map[string]int)
	for _, w := range words {
		m[w] = len(w)
	}
	return m
}
