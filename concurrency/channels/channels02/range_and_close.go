package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {

		// write to the channel n times
		ch <- x

		x, y = y, x+y

	}

}

func main() {

	bufferSize := 10

	ch := make(chan int, bufferSize)

	// run goroutine in anonymous function
	go func() {

		// fibonacci: write n fibonacci numbers to a channel.
		fibonacci(cap(ch), ch)

		// close the channel after writing all values
		close(ch)

	}()

	// read all values written to the channel
	for ele := range ch {

		fmt.Println("read from channel")
		fmt.Println("Ele:", ele)
	}

	// check if a channel is open | false indicates a closed channel
	_, ok := <-ch

	fmt.Println("open channel:", isopenChannel(ok))

}

func isopenChannel(isOpen bool) string {
	if isOpen {
		return "yes"
	} else {
		return "no"
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
//   - unlike other for range operations you only receive the value and not the index
