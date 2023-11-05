package main

import (
	"fmt"
	"time"
)

func main() {

	var tick <-chan time.Time = time.Tick(100 * time.Millisecond)

	var boom <-chan time.Time = time.After(500 * time.Millisecond)

	for {
		select {
		case tickTimeStamp := <-tick:
			fmt.Println("TICK! at: ", tickTimeStamp)
		case boomTimeStamp := <-boom:
			fmt.Println("BOOM! at: ", boomTimeStamp)
			return
		default:
			fmt.Println("		.")
			time.Sleep(50 * time.Millisecond)
		}
	}

}

// Default Selection

//   - a select statements default case runs if no other case is ready to communicate
//   - can be used to attempt a send or receive without blocking

// time.Tick

//   - provides access to the ticking channel only and returns a time.Timer of channel type

// time.After

//   - after the elapsed time returns a time.Timer of channel type with the current time

// time.Sleep

//   - pauses a goroutine for the duration
