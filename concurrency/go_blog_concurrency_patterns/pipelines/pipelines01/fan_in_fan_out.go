package main

// TODO: Review

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	intArray := []int{}

	inChannel := gen(1, 2, 3, 5, 8, 13)

	// fan-out: distribute work across two goroutines who both read from inChannel

	squaredChannel00 := square(inChannel)

	squaredChannel01 := square(inChannel)

	mergedChannels := merge(squaredChannel00, squaredChannel01)

	for value := range mergedChannels {
		intArray = append(intArray, value)
	}

	fmt.Println(intArray)

	// for val := range mergedChannels {
	// 	fmt.Println(val)
	// }

	// for value := range squaredChannel00 {
	// 	intArray = append(intArray, value)
	// }

}

func gen(nums ...int) <-chan int {
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
			out <- num * num
		}

		close(out)

	}()

	return out
}

func merge(cs ...<-chan int) <-chan int {

	var waitGroup sync.WaitGroup

	out := make(chan int)

	numRoutinesRan := 0

	output := func(in <-chan int) {
		time.Sleep(50 * time.Millisecond)
		for num := range in {
			out <- num
		}

		numRoutinesRan += 1

		fmt.Println("Goroutines ran:", numRoutinesRan)

		waitGroup.Done()
	}

	waitGroup.Add(len(cs))

	for _, ch := range cs {
		go output(ch)
	}

	go func() {
		fmt.Println("waiting for all goroutines...")
		waitGroup.Wait()
		fmt.Println("all goroutines are finished!")
		close(out)
	}()

	return out

}

// fan-out

//   - multiple goroutines read from the same channel until the channel is closed
//   - a way of distributing work across goroutines, parallelizing CPU use and I/O

// fan-in

//   - receive from multiple channels funneled into one channel
