package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web = fakeSearch("Web")

	Image = fakeSearch("Image")

	Video = fakeSearch("Video")
)

type Result string

type Search func(query string) Result

func main() {

	runs := 10

	var totalElapsedTime time.Duration

	for i := 0; i < runs; i++ {

		startTime := time.Now()

		fmt.Println(Google("Shunryū Suzuki"))

		elapsedTime := time.Since(startTime)

		fmt.Println(elapsedTime)

		totalElapsedTime += elapsedTime
	}

	fmt.Println("Average elapsed time:", totalElapsedTime/time.Duration(runs))

}

func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}

func Google(query string) []Result {
	results := []Result{}

	resultChannel := make(chan Result)

	timeout := time.After(100 * time.Millisecond)

	// fan-in pattern
	go func() { resultChannel <- Web(query) }()

	go func() { resultChannel <- Image(query) }()

	go func() { resultChannel <- Video(query) }()

	// timeout select pattern
	for i := 0; i < 3; i++ {
		select {

		case result := <-resultChannel:
			results = append(results, result)

			// end all communication after timeout
		case <-timeout:
			fmt.Println("Service timeout...")
			return results
		}
	}

	return results
}

// Timeout Concurrent Service Call

//   - in our example above we timeout all communcation to our backend service calls after the timeout expiration
