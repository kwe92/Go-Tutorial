package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	ch := getTime()

	var timeOut <-chan time.Time

	// timeout communication after two second.
	timeOut = time.After(2 * time.Second)

	var waitGroup sync.WaitGroup

	waitGroup.Add(1)

	go func() {

		defer waitGroup.Done()

		for {
			select {

			case val := <-ch:
				fmt.Println(val)
				return

			// after timeout if no other channel has returned the case statement is invoked breaking out of the loop
			case timeout := <-timeOut:
				fmt.Printf("%v: channel communication timeout...please try again.\n", timeout)
				return

			}
		}

	}()

	waitGroup.Wait()
}

func getTime() <-chan time.Time {

	out := make(chan time.Time)

	go func() {

		out <- <-time.After(3 * time.Second)

		close(out)
	}()

	return out
}

// Timing Out Select Statements

//   - you can timeout select statements by adding a default case that reads a call from time.After

// time.After

//   - after the passed in duration a read channel time.Timer returns
