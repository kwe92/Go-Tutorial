package main

import "fmt"

func main() {

	ch := fibonacci(10)

	// read all values written to the channel until closed
	for ele := range ch {
		fmt.Println("read from channel")
		fmt.Println("Ele:", ele)
	}

	// check if a channel is open | false indicates a closed channel
	fmt.Println("open channel:", isopen(ch))

}

func fibonacci(n int) <-chan int {

	out := make(chan int)

	x, y := 0, 1

	// run goroutine in anonymous function

	go func() {

		// fibonacci: write n fibonacci numbers to a channel.
		for i := 0; i < n; i++ {

			// write to the channel n times
			out <- x

			x, y = y, x+y

		}

		// close the channel after writing all values
		close(out)
	}()

	return out

}

func isopen[T interface{}](in <-chan T) bool {

	if _, ok := <-in; ok {
		return true
	} else {
		return false
	}

}

// Channel Close | Sender

//   - a sender to a channel can close the channel indicating there will be no more values sent
//   - senders should only close channels trying to send to a closed channel results in a panic
//   - Channels are not like files where the stream needs to be closed once opened
//   - if a for range loop reads a channel the channel must be closed or the for range loop will listen infinitely

// Check Channel Close | Receiver

//   - receivers can use the comma ok idiom to check if a channel was closed
//   - false indicates that the channel is closed

// Channels & for range

//   - receives values from a channel until the channel is closed
//   - unlike other for range operations you only receive the one value instead of two
