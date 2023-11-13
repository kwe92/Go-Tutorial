package main

import "fmt"

func main() {
	buffered := make(chan int, 1)

	buffered <- 42

	fmt.Println(<-buffered)

	buffered <- 42

	fmt.Println(<-buffered)

	// The bellow code will cause a deadlock as we are trying to write two values in a buffer that can only hold one

	buffered <- 42

	buffered <- 43

}

// Buffered Channel

//   - a channel with a finite capacity
//   - buffered channels can be written to until full without an explicit receiver unlike unbuffered channels
//   - buffered channels can be written to and read from within the same goroutine

// Buffered Channel Deadlocking

//   - buffered channels deadlock when:

//       ~ you try to read from an empty channel
//       ~ you try to exceed the capacity of the buffered channel
