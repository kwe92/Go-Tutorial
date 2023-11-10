package main

import "fmt"

func main() {

	// Final Stage: Setup the pipeline.

	generatorChannel := generator(1, 2, 3, 4, 5)

	ch := square(generatorChannel)

	// consume output from stage 2 until the channel is closed

	for val := range ch {
		fmt.Println(val)
	}

}

// Stage One

// generator: variadic function that converts a Slice of integers into a receive-only channel
//
//	that emits the integers written from the Slice
func generator(nums ...int) <-chan int {

	// create out channel
	out := make(chan int)

	// launch a goroutine that writes all passed in values to the out channel and closes it
	go func() {
		for _, num := range nums {

			// write to the out channel
			out <- num
		}

		// close channel after sending all values
		close(out)
	}()

	return out
}

// Stage Two

// square: receive a channel of integers and return a channel that emits the square of each received integer.
func square(in <-chan int) <-chan int {

	// create an out channel
	out := make(chan int)

	// launch a goroutine that reads from the input channel does some computation and writes the results to the out channel
	go func() {

		for num := range in {

			out <- num * num
		}

		// close channel after writing all values
		close(out)
	}()

	return out
}

// Go Pipelines

//   - loosely defined as a way to construct concurrent programs in GO
//   - composed of stages connected by channels
//   - each stage is composed of a group of goroutines
//     running the same function

// Stages and Their Routines

//   - responsibility of goroutines within a stage:

//       - receive values from upstream via inbound channels
//       - perform some function on received upstream values, producing new values
//       - write the new values to an outbound channel and send the values of the outbound channel downstream
//       - the sender should always close the channel when they are done writing values
