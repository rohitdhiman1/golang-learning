package main

import (
	"errors"
	"fmt"
	"strconv"
)

var (
	ErrTooYoung = errors.New("too young")
	ErrTooOld   = errors.New("too old")
)

type FieldError struct {
	Field   string
	Message string
}

func (e *FieldError) Error() string {
	return fmt.Sprintf("%s: %s", e.Field, e.Message)
}

func Part5Errors() {
	// Exercise 1: Basic error with wrapping
	fmt.Println("=== Exercise 1: Basic Error + Wrapping ===")
	age, err := parseAge("25")
	fmt.Printf("  parseAge(\"25\"): age=%d, err=%v\n", age, err)

	age, err = parseAge("abc")
	fmt.Printf("  parseAge(\"abc\"): age=%d, err=%v\n", age, err)

	// Exercise 2: Sentinel errors
	fmt.Println("\n=== Exercise 2: Sentinel Errors ===")
	for _, testAge := range []int{10, 25, 200} {
		err := validateAge(testAge)
		if err == nil {
			fmt.Printf("  validateAge(%d): OK\n", testAge)
		} else if errors.Is(err, ErrTooYoung) {
			fmt.Printf("  validateAge(%d): too young (errors.Is matched ErrTooYoung)\n", testAge)
		} else if errors.Is(err, ErrTooOld) {
			fmt.Printf("  validateAge(%d): too old (errors.Is matched ErrTooOld)\n", testAge)
		}
	}

	// Exercise 3: Custom error type with errors.As
	fmt.Println("\n=== Exercise 3: Custom Error Type ===")
	err = validateUser("", "bademail")
	fmt.Printf("  validateUser(\"\", \"bademail\"): %v\n", err)
	var fe *FieldError
	if errors.As(err, &fe) {
		fmt.Printf("  errors.As → Field=%q, Message=%q\n", fe.Field, fe.Message)
	}

	// Exercise 4: Error wrapping chain
	fmt.Println("\n=== Exercise 4: Error Wrapping Chain ===")
	err = handleRequest("age", "notanumber")
	fmt.Printf("  full chain: %v\n", err)

	var numErr *strconv.NumError
	if errors.As(err, &numErr) {
		fmt.Printf("  errors.As found *strconv.NumError at root: Func=%s, Num=%q\n", numErr.Func, numErr.Num)
	}

	// Exercise 5: panic and recover
	fmt.Println("\n=== Exercise 5: Panic & Recover ===")
	result, err := safeDivide(10, 3)
	fmt.Printf("  safeDivide(10, 3): result=%d, err=%v\n", result, err)

	result, err = safeDivide(10, 0)
	fmt.Printf("  safeDivide(10, 0): result=%d, err=%v\n", result, err)
}

func parseAge(s string) (int, error) {
	age, err := strconv.Atoi(s)
	if err != nil {
		return 0, fmt.Errorf("parseAge(%q): %w", s, err)
	}
	return age, nil
}

func validateAge(age int) error {
	if age < 18 {
		return fmt.Errorf("age %d: %w", age, ErrTooYoung)
	}
	if age > 120 {
		return fmt.Errorf("age %d: %w", age, ErrTooOld)
	}
	return nil
}

func validateUser(name, email string) error {
	if name == "" {
		return &FieldError{Field: "name", Message: "required"}
	}
	if len(email) < 3 {
		return &FieldError{Field: "email", Message: "too short"}
	}
	return nil
}

func handleRequest(fieldName, rawValue string) error {
	err := validateInput(fieldName, rawValue)
	if err != nil {
		return fmt.Errorf("handleRequest: %w", err)
	}
	return nil
}

func validateInput(fieldName, rawValue string) error {
	_, err := parseField(rawValue)
	if err != nil {
		return fmt.Errorf("validateInput(%s): %w", fieldName, err)
	}
	return nil
}

func parseField(raw string) (int, error) {
	v, err := strconv.Atoi(raw)
	if err != nil {
		return 0, fmt.Errorf("parseField: %w", err)
	}
	return v, nil
}

func safeDivide(a, b int) (result int, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("recovered panic: %v", r)
		}
	}()
	return a / b, nil
}
