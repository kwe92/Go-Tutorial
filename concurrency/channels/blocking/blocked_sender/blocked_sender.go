package main

import (
	"fmt"
)

func main() {

	// create a channel of integers
	unbuffered := make(chan int)

	// write to the channel | there is no other goroutine to read the written value so the channel blocks
	unbuffered <- 1

	// read from the channel
	fmt.Println(<-unbuffered)

}

// Writing and Reading Within The Same goroutine

//   - Writing to a unbuffered channel and then reading from that channel within the same goroutine will deadlock
//   - you must write to and read from separate goroutines
//   - writing to an unbuffered channel without reading the value in another goroutine will also deadlock
