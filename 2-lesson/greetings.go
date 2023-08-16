package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello %v. Welcome to the Go Developer Community", name)
	return message
}

// Function Names

//   - function names start with a capital letter
