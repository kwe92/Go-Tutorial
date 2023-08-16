package greetings

import "fmt"

func Hello(name string) string {
	message := fmt.Sprintf("Hello %v. Welcome to the Go Developer Community", name)
	return message
}
