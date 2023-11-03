package main

import "fmt"

func fibonacci(n int, ch chan int) {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)
}

func main() {

	bufferSize := 10

	ch := make(chan int, bufferSize)

	go func() {
		fibonacci(bufferSize, ch)
	}()

	for ele := range ch {

		fmt.Println("Channel open: ")
		fmt.Println("Ele: ", ele)
	}

}

// Channel Close | Sender

//   - a sender to a channel can close the channel indicating there will be no more values sent
//   - senders should only close channels trying to send to a closed channel is like walking through a closed door
//     it will cause a panic
//   - Channels are not like files where the stream needs to be closed once opened
//   - if a for range loop reads a channel the channel should be closed

// Check Channel Close | Receiver

//   - receivers can use the comma ok idiom to check if a channel was closed
//   - false indicates that the channel is closed

// Channels  & for range

//   - receive values from a channel until it is closed

//
