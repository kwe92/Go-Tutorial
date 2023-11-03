package main

import "fmt"

//TODO: understand how the value is being passed around

//! Explain this code works

func main() {

	ch := make(chan int)

	done := make(chan struct{})

	go func() {
		for i := 0; i < 10; i++ {
			// channel read
			fmt.Println(<-ch)
		}
		// channel done pattern
		done <- struct{}{}
	}()

	for i := 0; i < 10; i++ {
		// channel write
		ch <- (i)

	}

	// ! explain more complicated examples

	// fibonacci(ch, done)
	// addOne(ch, done)

}

func addOne(ch chan int, done chan struct{}) {

	x := 0

	for {
		select {
		case ch <- x:
			x++
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

//   - gives a go-routine the ability to wait on multiple communication operations
//   - the select blocks until one of the cases can run and then runs that case
//   - if multiple cases are ready then one is executed at random
