package main

import "fmt"

func main() {

	n := 10

	ch := make(chan int)

	done := make(chan struct{})

	go func() {
		for i := 0; i < n; i++ {
			// reading from the channel n times
			fmt.Println(<-ch)
		}

		// channel done pattern | initialize the done channel so it is no longer nil
		done <- struct{}{}

	}()

	addOne(ch, done)

}

func addOne(ch chan int, done chan struct{}) {

	x := 0

	for {
		select {

		// write to channel
		case ch <- x:
			x++

		// if done is not nil print done and return out of the loop
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

func fibonacci(ch chan int, done chan struct{}) {
	x, y := 0, 1

	for {
		select {
		case ch <- x:
			x, y = y, x+y
		case <-done:
			fmt.Println("done")
			return
		}
	}
}

// Select Statement

//   - gives a goroutine the ability to wait on multiple communication operations
//   - the select blocks until one of the cases can run and then runs that case
//   - if multiple cases are ready then one is executed at random
//   - cases with a nil value are skipped

// Done Pattern
