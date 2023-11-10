package main

import (
	"fmt"
	"time"
)

func main() {

	readChannel := getTime()

	var timeOut <-chan time.Time

	// timeout communication after one second.
	timeOut = time.After(time.Second)

	for {
		select {

		case val := <-readChannel:
			fmt.Println(val)

		// after timeout if no other channel has returned the case statement is invoked breaking out of the loop
		case timeout := <-timeOut:
			fmt.Printf("%v: channel communication timeout...please try again.\n", timeout)
			return

		}
	}
}

func getTime() <-chan time.Time {

	writeChannel := make(chan time.Time)

	go func() {

		writeChannel <- <-time.After(3 * time.Second)

		close(writeChannel)
	}()

	readChannel := (<-chan time.Time)(writeChannel)

	return readChannel
}

// Timing Out Select Statements

//   - you can timeout select statements by adding a default case that reads a call from time.After

// time.After

//   - after the passed in duration a read channel time.Timer returns
