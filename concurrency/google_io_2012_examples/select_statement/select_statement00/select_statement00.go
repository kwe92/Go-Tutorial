package main

import (
	"fmt"
	"time"
)

func main() {

	n := 20

	ch := fanIn[string](say("hello"), say("goodbye"))

	for i := 0; i < n; i++ {
		// read from channel n times
		fmt.Println(<-ch)

	}

}

// say: write the passed in message to a channel n times and return a the channel as a read channel.

func say(msg string) <-chan string {

	out := make(chan string)

	go func() {

		defer func() {
			close(out)
			fmt.Println("channel closed")
		}()
		for {
			out <- msg
			time.Sleep(500 * time.Millisecond)
		}

	}()

	return out

}

// fanIn: write two channels to one channel concurrently and return the channel
func fanIn[T any](ch00, ch01 <-chan T) <-chan T {

	out := make(chan T)

	go func() {

		defer close(out)

		// receive from the first available channel, if no channels are available execute the default case
		for {
			select {
			case var00 := <-ch00:
				out <- var00

			case var01 := <-ch01:
				out <- var01
			default:
				fmt.Println("Waiting for a channel...")
				time.Sleep(time.Millisecond * 100)
			}
		}
	}()

	return out
}

// Select Statement

//   - a control structure unique to concurrency, used to route data along multiple channels
//   - select statements are switch statements for channel operations
//   - each case instead of being an expression is a communication `Send or Receive too or from a channel respectively`
//   - controls the behavior of a program based on what communications are able to proceed at any given moment
//   - provides a way to handle multiple channels
//   - all channels are evaluated
//   - the selection blocks or runs the default case until a channel is ready to communicate
//   - if multiple channels are ready to proceed one is choosen pseudo-randomly
//   - you can not depend on the order of which communication will proceed
