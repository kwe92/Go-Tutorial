package main

// TODO: review

import (
	"fmt"
	"time"
)

func main() {

	n := 20

	readChannel := fanIn[string](say("hello"), say("goodbye"))

	for i := 0; i < n; i++ {
		// read from channel n times
		fmt.Println(<-readChannel)

	}

}

func say(msg string) <-chan string {

	writeChannel := make(chan string)

	go func() {
		for {
			writeChannel <- msg
			time.Sleep(500 * time.Millisecond)
		}
	}()

	// convert write channel to read channel after all values are written
	readChannel := (<-chan string)(writeChannel)

	return readChannel

}

func fanIn[T any](ch00, ch01 <-chan T) <-chan T {

	// declare and initialize channels
	writeChannel := make(chan T)

	// run two independent goroutines that write to the same channel

	go func() {
		for {
			writeChannel <- <-ch00
		}
	}()

	go func() {
		for {
			writeChannel <- <-ch01
		}
	}()

	// convert write channel to read channel after all values are written
	readChannel := (<-chan T)(writeChannel)

	return readChannel

}

// Multiplexer | fan-in function

//   - two independant communications are combined into one over a shared medium
//   - whoever is ready first speaks first
//   - run independent goroutines out of sequence
