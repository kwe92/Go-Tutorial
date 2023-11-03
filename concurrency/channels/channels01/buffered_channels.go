package main

import "fmt"

func main() {
	ch := make(chan int, 2)

	ch <- 1

	ch <- 2

	fmt.Println(<-ch)

	fmt.Println(<-ch)

}

// Buffered Channel

//   - channels with capacity
//   - sends are blocked when the channel is full
//   - receives are blocked when the channel is full
// ?  - don't requred reads
