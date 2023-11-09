package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Search func(query string) Result

// Fake replicated backend server calls
var (
	Web00 = fakeSearch("Web00")
	Web01 = fakeSearch("Web01")

	Image00 = fakeSearch("Image00")
	Image01 = fakeSearch("Image01")

	Audio00 = fakeSearch("Audio00")
	Audio01 = fakeSearch("Audio01")
)

func main() {
	n := 10

	var totalDuration time.Duration

	queryString := "Marcus Aurelius"

	for i := 0; i < n; i++ {

		start := time.Now()

		fmt.Println(Google(queryString))

		elapsed := time.Since(start)

		fmt.Println(elapsed)

		totalDuration += elapsed

	}

	fmt.Println("Average search time:", totalDuration/time.Duration(n))
}

// fakeSearch: save the state of the kind of search and return a closure of Search type.
func fakeSearch(kind string) Search {
	return func(query string) Result {
		time.Sleep(time.Duration(rand.Intn(200)) * time.Millisecond)
		return Result(fmt.Sprintf("%s result for %q", kind, query))
	}
}

// first: launch the same search multiple times in different goroutines and return the first result attained
func first(query string, replicas ...Search) Result {

	// declare and initialize channel
	resultChannel := make(chan Result)

	// launch a goroutine for every replicated server call
	for i := range replicas {
		go func(i int) { resultChannel <- replicas[i](query) }(i)
	}

	// read from the channel once, the first value that was written by the launched goroutines
	result := <-resultChannel

	return result

}

func Google(query string) []Result {

	results := []Result{}

	resultsChannel := make(chan Result)

	timeout := time.After(100 * time.Millisecond)

	// fan-in pattern

	go func() { resultsChannel <- first(query, Web00, Web01) }()

	go func() { resultsChannel <- first(query, Image00, Image01) }()

	go func() { resultsChannel <- first(query, Audio00, Audio01) }()

	for i := 0; i < 3; i++ {

		select {
		case result := <-resultsChannel:
			results = append(results, result)

		// select timeout pattern | timeout ALL
		case <-timeout:
			fmt.Println("server timeout...")
			return results
		}
	}

	return results

}

// Reduce Tail Latency

//   - if you have a service that has multiple available servers you can call n number of servers per service and
//      return the result from the first server that provides the service
//   - these service calls should happen in parallel and only return the first result
