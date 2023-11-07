package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Fake Backend Service Calls
var (
	Web = fakeSearch("Web")

	Image = fakeSearch("Image")

	Video = fakeSearch("Video")
)

type Result string

type Search func(query string) Result

func main() {

	var totalDuration time.Duration

	runs := 20

	for i := 0; i < runs; i++ {

		startTime := time.Now()

		fmt.Println(Google("Allan Watts"))

		elapsedTime := time.Since(startTime)

		fmt.Println("time taken to complete search:", elapsedTime)

		totalDuration += elapsedTime
	}

	fmt.Printf("Average search time: %v\n", totalDuration/time.Duration(runs))

}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}

func Google(query string) []Result {

	results := []Result{}

	var resultChannel chan Result

	resultChannel = make(chan Result)

	// launch each backend service in an independant goroutine, calling and processing each backend service in parallel
	// the goroutines then fan-in to one channel, whichever goroutine is ready first writes to the channel first

	go func() {
		resultChannel <- Web(query)
		fmt.Println("Web result written")
	}()

	go func() {
		resultChannel <- Image(query)
		fmt.Println("Image result written")
	}()

	go func() {
		resultChannel <- Video(query)
		fmt.Println("Video result written")
	}()

	for i := 0; i < 3; i++ {

		// for each goroutine read the value written to the channel
		var result Result = <-resultChannel

		// append the read value from the channel to a slice of Results
		results = append(results, result)
	}

	return results

}

// Concurrent Fake Google Search | multiple backends running independant of one another

//   - in the none concurrent example the total duration was longer due to each
//     backend service call taking up to 200 ms to run synchronously

//   - in the concurrent example the backend service calls are launched in independant goroutines
//     that run in parallel, the total duration is dependant on the longest running proccess
//     not the accumulation of all synchronous processes as in our non-concurrent example
