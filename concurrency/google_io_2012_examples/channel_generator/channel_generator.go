package main

import "fmt"

func main() {

	n := 5

	// receive read-only channels
	readChannel00 := say("readChannel00")

	readChannel01 := say("readChannel01")

	// read n values from the channels and initiate synchronization
	for i := 0; i < n; i++ {

		// read from channel00
		fmt.Println("you say: ", <-readChannel00)

		// read from channel01
		fmt.Println("you say: ", <-readChannel01)

	}
}

// say: writes message to channel n times and returns a read-only channel.
func say(msg string) <-chan string {

	// define channel
	var writeChannel chan string

	// initialize channel | zero-value is nil
	writeChannel = make(chan string)

	// launch goroutine to write to the channel
	go func() {
		// infinite loop with index to write n number of values to the channel
		// will write as many values as read in the calling function
		for i := 0; ; i++ {
			// write to channel
			writeChannel <- fmt.Sprintf("%s %d", msg, i)
		}

	}()

	// convert write channel to read channel after all values are written
	readChannel := (<-chan string)(writeChannel)

	return readChannel

}

// Synchronization

//   - synchronization allows the channels to communicate synchronously `in lockstep`
//   - channels will take turns executing and returning their values
//   - in the example above if readChannel01 is ready to send a value but readChannel00 has not sent
//     a value readChannel01 is blocked until readChannel00 sends a value

// Channel Generator

//   - a function that returns a channel
