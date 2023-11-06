package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})

	readchannel := say("begin.", done)

	for i := 0; i < 4; i++ {
		fmt.Println(<-readchannel)
	}

	// write to the done channel to indicate you are done using the servie
	done <- struct{}{}

}

func say(msg string, done chan struct{}) <-chan string {

	writeChannel := make(chan string)

	go func() {
		for {
			select {
			case writeChannel <- msg:
				time.Sleep(500 * time.Millisecond)
			case <-done:
				return
			}
		}
	}()

	// convert write channel to read channel after all values are written
	readChannel := (<-chan string)(writeChannel)

	return readChannel

}

// Quit / Done Channel

//   - deterministically say that you are done using a service
//   - the caller to a channel can specify when they are done using the channel as a service
