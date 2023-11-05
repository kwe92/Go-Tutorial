package main

import (
	"fmt"
	"time"
)

// TODO: review

// Message: contains a channel for a reply.
type Message struct {
	str  string
	wait chan bool
}

func main() {

	n := 5

	readChannel := fanIn[Message](say("hello"), say("goodbye"))

	for i := 0; i < n; i++ {
		// read from channel n times
		msg00 := <-readChannel
		fmt.Println(msg00.str)

		msg01 := <-readChannel
		fmt.Println(msg01.str)

		msg00.wait <- true
		msg01.wait <- true

	}

}

func say(msg string) <-chan Message {

	writeChannel := make(chan Message)

	waitForit := make(chan bool)

	go func() {
		for {
			writeChannel <- Message{
				str:  msg,
				wait: waitForit,
			}

			<-waitForit

			time.Sleep(time.Second)
		}
	}()

	readChannel := writeChannel

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

	readChannel := writeChannel

	return readChannel

}

// Restoring Sequencing

//   - goroutines are lockstep
//   - Sends a channel on a channel forcing goroutine to wait its turn
//   - Receive all messages then enable channels by sending on a private channel
