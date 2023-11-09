package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	// main function start time
	start := time.Now()

	in := generator(1, 1, 2, 3, 5, 8)

	// Run three goroutines processing the same channel distributing the work

	squaredChannel00 := square(in)

	squaredChannel01 := square(in)

	squaredChannel02 := square(in)

	// fanin the results of all three goroutines.
	out := merge[int](squaredChannel00, squaredChannel01, squaredChannel02)

	// for num := range squaredChannel00 {
	// 	fmt.Println(num)
	// }

	for num := range out {
		fmt.Println(num)
	}

	// time taken to run main
	elapsed := time.Since(start)

	fmt.Println(elapsed)
}

func generator(nums ...int) <-chan int {

	out := make(chan int)

	go func() {
		for _, num := range nums {
			out <- num
		}
		close(out)
	}()

	return out
}

func square(in <-chan int) <-chan int {
	out := make(chan int)

	go func() {

		for num := range in {
			time.Sleep(500 * time.Millisecond)

			out <- num * num
		}
		close(out)
	}()

	return out
}

func merge[T interface{}](cs ...<-chan T) <-chan T {

	// wait for a collection of goroutines to finish
	var waitGroup sync.WaitGroup

	out := make(chan T)

	// define a closure that writes all values from a channel into the input channel and reduces the wait group counter by 1
	outputClosure := func(in <-chan T) {
		for val := range in {
			out <- val
		}
		// reduce the wait group counter by 1
		waitGroup.Done()
	}

	// set the number of goroutines to wait for
	waitGroup.Add(len(cs))

	for _, channel := range cs {
		go outputClosure(channel)
	}

	go func() {

		// blocks until the wait group counter is 0
		waitGroup.Wait()

		// close the output channel after all routines have finished
		close(out)

	}()

	return out

}
