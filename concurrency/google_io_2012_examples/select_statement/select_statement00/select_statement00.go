package main

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

// say: write the passed in message to a channel n times and return a the channel as a read channel.

func say(msg string) <-chan string {

	writeChannel := make(chan string)

	go func() {
		for {
			writeChannel <- msg
			time.Sleep(500 * time.Millisecond)
		}
	}()

	readChannel := (<-chan string)(writeChannel)

	return readChannel

}

// fanIn: write two channels to one channel concurrently and return the channel
func fanIn[T any](ch00, ch01 <-chan T) <-chan T {
	writeChannel := make(chan T)

	go func() {
		for {
			select {
			case var00 := <-ch00:
				writeChannel <- var00

			case var01 := <-ch01:
				writeChannel <- var01
			default:
				fmt.Println("Waiting for a channel...")
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	readChannel := (<-chan T)(writeChannel)

	return readChannel
}

// Select Statement

//   - a control structure unique to concurrency
//   - select statements are somewhat like switch statements
//   - each case instead of being an expression is a communication `Send or Receive to a channel`
//   - controls the behavior of a program based on what communications
//     are able to proceed at any given moment
//   - provides a way to handle multiple channels
//   - all channels are evaluated
//   - the selection blocks or runs the default case until a channel is ready to communicate
//   - if multiple channels are ready to proceed one is choosen pseudo-randomly
//   - you can not depend on the order of which communication will proceed
