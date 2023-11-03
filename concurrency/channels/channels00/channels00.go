package main

import "fmt"

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

	// the index to split the Slice

	splitAtIndex := len(values) / 2

	// split data in half | if odd length the second chunk will have one additional element

	chunk00 := values[:splitAtIndex]

	chunk01 := values[splitAtIndex:]

	fmt.Println(chunk00)

	fmt.Println(chunk01)

	// channels must be created before use

	c := make(chan int)

	// distribute the summation of the values Slice between two GO routines

	go sum(chunk00, c)

	go sum(chunk01, c)

	// receive from the channel for each go-routine executed

	x := <-c

	y := <-c

	fmt.Println(x + y)

}

// Channels

//   - a typed conduit that sends and receives values

// Channel Operator: <-

//   - use the channel operator to send and receive values to and from variables respectively
//   - the direction does not change but the operand to the left determines which action you are taking
//   - the data flows in the direction of the arrow

// Synchronize Without Conditions

//   - sends and receives are blocked until the other side is ready
//   - implying that they execute synchronously even though go routines are an asynchronous operation
