package main

import (
	"fmt"
	"math/rand"
	"time"
)

var (
	Web   = fakeSearch("Web")
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

		fmt.Println(Google("Marcus Aurelius"))

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
	results := make([]Result, 0)

	results = append(results, Web(query))

	results = append(results, Image(query))

	results = append(results, Video(query))

	return results

}

// System Software

//   - Go is designed to write system software

// Example: Fake Google Search Engine | nonconcurrent

// Q: What does Google search do?

// A: Given a query, return a page of search results.

// Q: How are search results received?

// A: Send the given query to independantly executing backends and merge the results. (Web Search, Image Search, News Search etc.)

// In parallel you want to send the query to the backends and merge the results from all backends

// Implement the behavior specified above.
