package main

import "fmt"

type EmailNotification struct {
	To      string
	Subject string
}

func (e EmailNotification) Send() string {
	return fmt.Sprintf("Email to %s: %s", e.To, e.Subject)
}

type SMSNotification struct {
	Phone   string
	Message string
}

func (s SMSNotification) Send() string {
	return fmt.Sprintf("SMS to %s: %s", s.Phone, s.Message)
}

type Notification interface {
	Send() string
}

type MyError struct {
	Code    int
	Message string
}

func (e *MyError) Error() string {
	return fmt.Sprintf("error %d: %s", e.Code, e.Message)
}

func Part4TypeAssertions() {
	// Exercise 1: Safe type assertion
	fmt.Println("=== Exercise 1: Safe Type Assertion ===")
	var v any = "hello"

	s, ok := v.(string)
	fmt.Printf("  v.(string): val=%q, ok=%t\n", s, ok)

	i, ok := v.(int)
	fmt.Printf("  v.(int):    val=%d, ok=%t  ← wrong type, zero value returned\n", i, ok)

	// Exercise 2: Type switch
	fmt.Println("\n=== Exercise 2: Type Switch ===")
	values := []any{42, "hello", true, []int{1, 2, 3}, 3.14}
	for _, val := range values {
		fmt.Printf("  describe(%v) = %s\n", val, describe(val))
	}

	// Exercise 3: The nil interface trap
	fmt.Println("\n=== Exercise 3: Nil Interface Trap ===")
	err := returnsNilPointerBug()
	fmt.Printf("  returnsNilPointerBug(): err == nil? %t  ← false! (type=*MyError, value=nil)\n", err == nil)
	if err != nil {
		fmt.Printf("  err.Error() would panic if we called it (nil pointer dereference)\n")
	}

	err = returnsNilFixed()
	fmt.Printf("  returnsNilFixed():      err == nil? %t  ← correct\n", err == nil)

	// Exercise 4: Interface-based polymorphism with type switch
	fmt.Println("\n=== Exercise 4: Polymorphism + Type Switch ===")
	notifications := []Notification{
		EmailNotification{To: "alice@example.com", Subject: "Meeting"},
		SMSNotification{Phone: "+1234567890", Message: "Reminder"},
		EmailNotification{To: "bob@example.com", Subject: "Deploy done"},
	}
	broadcast(notifications)
}

func describe(v any) string {
	switch x := v.(type) {
	case int:
		return fmt.Sprintf("int(%d)", x)
	case string:
		return fmt.Sprintf("string(%q)", x)
	case bool:
		return fmt.Sprintf("bool(%t)", x)
	case []int:
		return fmt.Sprintf("[]int(len=%d)", len(x))
	default:
		return fmt.Sprintf("unknown(%T)", x)
	}
}

func returnsNilPointerBug() error {
	var err *MyError = nil
	return err // NOT nil — interface has type=*MyError, value=nil
}

func returnsNilFixed() error {
	var err *MyError = nil
	if err == nil {
		return nil // explicitly return nil interface
	}
	return err
}

func broadcast(notifications []Notification) {
	for _, n := range notifications {
		result := n.Send()
		switch n.(type) {
		case EmailNotification:
			fmt.Printf("  [EMAIL] %s\n", result)
		case SMSNotification:
			fmt.Printf("  [SMS]   %s\n", result)
		}
	}
}
