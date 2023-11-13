package main

import "fmt"

func main() {

	// create a channel | zero-value nil
	unbuffered := make(chan int)

	// reading from nil channel `channel without data` blocks causing a panic | runtime error
	fmt.Println(<-unbuffered)
}
