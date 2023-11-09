package main

import (
	"fmt"
	"time"
)

func main() {

	n := 5

	readChannel := getTime()

	var timeOut <-chan time.Time

	// timeout communication after one second.
	timeOut = time.After(time.Second)

	for i := 0; i < n; i++ {
		select {

		case val := <-readChannel:
			fmt.Println(val)

		// after timeout if no other channel has returned the case statement is invoked breaking out of the loop
		case <-timeOut:
			fmt.Println("channel communication timeout...please try again.")
			return

		}
	}
}

func getTime() <-chan time.Time {

	writeChannel := make(chan time.Time)

	go func() {
		for {
			writeChannel <- <-time.After(3 * time.Second)
		}
	}()

	readChannel := (<-chan time.Time)(writeChannel)

	return readChannel
}

// Timing Out Select Statements

//   - you can timeout select statements by adding a case that reads a call from time.After

// time.After

//   - after the passed in duration a read channel time.Timer returns
