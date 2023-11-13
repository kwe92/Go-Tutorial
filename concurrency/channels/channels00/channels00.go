package main

import "fmt"

// sum: sums the slice of integers passed in and write the result to the channel
func sum(values []int, c chan int) {

	var result int // initialized to the zero-value : 0

	for _, ele := range values {
		result += ele
	}

	// send result to channel
	c <- result
}

func main() {
	values := []int{1, 2, 3, 4, 5, 6, 7}

	// split index
	splitAtIndex := len(values) / 2 // 3

	fmt.Println(splitAtIndex)

	// split data in half | if odd length the second chunk will have one additional element

	chunk00 := values[:splitAtIndex] // {1,2,3}

	chunk01 := values[splitAtIndex:] // {4, 5, 6, 7}

	fmt.Println(chunk00)

	fmt.Println(chunk01)

	// channels must be created before use

	in := make(chan int)

	// fan-out: distribute the summation of the values Slice between two GO routines

	go sum(chunk00, in) // write to channel once

	go sum(chunk01, in) // write to channel a second time

	// read each value stored in the channel

	chunk00Result := <-in

	chunk01Result := <-in

	fmt.Println(chunk00Result + chunk01Result)

}

// Channels

//   - a typed conduit or pipe of communication that connects concurrent goroutines
//   - provides bidirectional communication between two goroutines
//   - values can be sent or received between connected goroutines
//   - a value from a channel can only be read once even when using the underscore pattern
//   - channels are a reference type / pointer type
//   - the zero value for a channel is nil
//   - the components of a channel are:
//     ~ Sender ~ Buffer (optional) ~ Receiver
//   - channels can be viewed as streams of data
//   - unbuffered channels must have a 1 to 1 mapping for reading and writing or the goroutines will deadlock

// Channel Operator: <-

//   - use the channel operator to read from or write to a channel
//   - the operand to the left determines which action you are taking
//   - the data flows in the direction of the arrow

// Reading Channel Values

//   - a channels value can only be read once regardless of use or ignoring the value with the underscore pattern
//   - if multiple goroutines read from the same channel only one of them will receive the value regardless of use

// Channel Types

//   - a channels static type can be of any type
//   - a channel can also be explicitly declared as a send channel `write channel` or receive channel `read channel`

// Synchronization (Goroutines Communicate & Synchronize)

//   - a sender or receiver can initiate synchronization
//   - receiver channels `<-ch` wait for a value to be sent
//   - sender channels ` ch <-` wait for a receiver to be ready
//   - sends and receives are blocked until the one side is ready
//   - senders and receivers execute synchronously even though go routines are an asynchronous operation
//   - a sender and receiver must be ready to work together or they will wait on one another indefinitely
//     or until the channel is closed

// Go Approach to Concurrent Software

//   - "Don't communicate by sharing memory, share memory by communicating."
//   - You dont have memory that needs protecting by puting locks and mutex's
//     around the memory to protect it from parallel access
//   - channels are used to pass data back and fourth
//     between goroutines
