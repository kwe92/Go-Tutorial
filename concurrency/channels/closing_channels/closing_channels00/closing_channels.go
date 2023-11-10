package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	generatedChannel := generator(1, 1, 2, 3, 5, 8)

	out := square(generatedChannel)

	fmt.Println("Channel is open:", isOpen(out))

	for num := range out {
		fmt.Println(num)
	}

	fmt.Println("Channel is open:", isOpen(out))

}

func isOpen(in <-chan int) bool {

	if _, ok := <-in; ok {

		return true

	} else {

		return false

	}

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
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)

		for num := range in {
			out <- num * num
		}

		close(out)
	}()
	return out
}

// Closing Channels

//   - closing a channel will generate a message readable by the receiver using the comma ok idiom pattern
//   - if a channel is closed the zero value of the channels static type and false are returned
//   - used by for range loops to read values from a channel until the channel is closed
//   - sending to a closed channel will panic, this is why sender close and receivers read values
//   - if a receiver ranges over a channel without the sender closing the channel
//     the receiver will attempt to read values indefinitely eventually deadlocking
//     as the receiver tries to read from an empty channel
