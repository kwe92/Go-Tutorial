package main

import (
	"fmt"
	"sync"
)

func main() {
	generator := gen(8, 13, 21)

	done := make(chan struct{})

	defer close(done)

	squaredGenerator00 := square(done, generator)

	squaredGenerator01 := square(done, generator)

	ch := merge(done, squaredGenerator00, squaredGenerator01)

	fmt.Println(<-ch)

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

func square(done <-chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)

	go func() {
		defer close(out)
		for num := range in {
			select {
			case out <- num * num:
			case <-done:
				fmt.Println("done from square")
				return
			}
		}

	}()

	return out
}

func merge(done <-chan struct{}, cs ...<-chan int) <-chan int {

	var waitGroup sync.WaitGroup

	out := make(chan int)

	output := func(in <-chan int) {
		defer waitGroup.Done()

		for num := range in {
			select {
			case out <- num:
			case <-done:
				fmt.Println("done from merge")
				return
			}
		}
	}

	waitGroup.Add(len(cs))

	for _, ch := range cs {
		go output(ch)
	}

	go func() {
		waitGroup.Wait()
		close(out)
	}()

	return out
}
