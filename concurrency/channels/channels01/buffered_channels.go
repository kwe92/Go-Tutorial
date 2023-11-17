package main

import "fmt"

func main() {

	// declare and initialize a channel of integers with a capacity of 2
	ch := make(chan int, 2)

	// launch goroutine that writes to the channel
	go func() {
		functionThatWritesToChannel(ch)

		close(ch)
	}()

	// read from the buffered channel based on its capacity
	for i := 0; i < cap(ch); i++ {
		fmt.Println("reading value from channel:", <-ch)
	}

}

func functionThatWritesToChannel(in chan<- int) {

	for i := 0; ; i++ {
		in <- i + 1
	}

}

// Buffered Channel

//   - not used often
//   - channels with capacity
//   - sends are blocked `deadlock` when the buffed channel is full
//   - receives are blocked `deadlock` when the buffed channel is empty
//   - buffered channels remove synchronization
//   - doesn't requred reads
