package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan string)

	readchannel := say("begin.", done)

	for i := 0; i < 4; i++ {
		fmt.Println(<-readchannel)
	}

	// indicate you are done using the channel as a service
	done <- "finished"

	// read from done channel to ensure any clean up is complete
	fmt.Println(<-done)

}

func say(msg string, done chan string) <-chan string {

	writeChannel := make(chan string)

	go func() {
		for {
			select {
			case writeChannel <- msg:
				time.Sleep(500 * time.Millisecond)
			case <-done:
				someCleanUpFunction()

				// write back to the caller to indicate cleanup is finished
				done <- "clean up complete."
				return
			}
		}
	}()

	// convert write channel to read channel after all values are written
	readChannel := (<-chan string)(writeChannel)

	return readChannel

}

func someCleanUpFunction() {
	fmt.Println("cleaning up...")
	time.Sleep(3 * time.Second)
}

// Quit / Done Channel

//   - deterministically say that you are done using a service
//   - receiving on quit / done allows the two channels to communicate their status
