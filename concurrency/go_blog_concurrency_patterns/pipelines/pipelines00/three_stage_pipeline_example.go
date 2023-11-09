package main

import "fmt"

func main() {

	// Final Stage: Setup the pipeline.

	generatorChannel := gen(1, 2, 3, 4, 5)

	ch := square(generatorChannel)

	// consume output from stage 2 until the channel is closed

	for val := range ch {
		fmt.Println(val)
	}

}

// Stage One

// gen: variadic function that converts a Slice of integers into a receive-only channel
//
//	that emits the integers written from the Slice
func gen(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, ele := range nums {
			out <- ele
		}

		// close channel after sending all values
		close(out)
	}()

	return out
}

// Stage Two

// square: receives integers from the passed in channel and returns a channel that emits the square of each received integer.
func square(in <-chan int) <-chan int {

	out := make(chan int)

	go func() {
		for num := range in {
			out <- num * num
		}

		// close channel after sending all values
		close(out)
	}()

	return out
}

// Go Pipelines

//   - Loosely defined as a way to construct concurrent programs in GO
//   - Composed of stages connected by channels
//   - each stage is comprised of a group of goroutines
//     running the same function

// Stages and Their Routines

//   - responsibility of goroutines within a stage:

//       - Receive values from upstream via inbound channels
//       - Perfom some function on received upstream values, producing new values
//       - write the new values to an outbound channel and send the values of the outbound channel downstream
