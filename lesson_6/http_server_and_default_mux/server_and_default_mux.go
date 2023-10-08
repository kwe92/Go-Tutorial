package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

const (
	Addr = ":8080"
)

func main() {

	// register a handler to a pattern on defaultServeMux
	http.HandleFunc("/", homeRouteHandler)

	fmt.Println("Server has started successfully, listening at", Addr)

	// start HTTP server on port 8080, using the default ServeMux
	log.Fatal(http.ListenAndServe(Addr, nil))

}

// homeRouteHandler handles requests to the home path
func homeRouteHandler(w http.ResponseWriter, r *http.Request) {

	// set response headers
	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set("X-Custom-Header", "In the beginning was the Word.")

	var responseData00 []byte

	var responseData01 []byte

	responseData00 = []byte("Why are ye fearful, O ye of little faith?\n\nThen he arose, and rebuked the winds and the sea; and there was a great calm.")

	responseData01 = []byte(" And all things, whatsoever ye shall ask in prayer, believing, ye shall receive.")

	responses := [][]byte{
		responseData00,
		responseData01,
	}

	// access request URL
	requestURL := r.URL

	host := r.URL.Host

	hostName := r.URL.Hostname()

	path := r.URL.Path

	port := r.URL.Port()

	query := r.URL.Query()

	verse := query.Get("verse")

	// convert verse to an integer
	verseNumber, err := strconv.Atoi(verse)

	// if there is no verse query parameter write an error to the response body and end execution
	if len(verse) == 0 {
		http.Error(w, fmt.Sprintf(`{"error": query parameter 'verse' is required, maximum indexable range %d}`, len(responses)-1), 404)
		return
	}

	// if there is an error parsing the verse number write an the error to the response body and end execution
	if err != nil {
		http.Error(w, fmt.Sprintf(`{"error": %s}`, err.Error()), 404)
		return
	}

	// if the requested verse number exceeds the indexable length of the Slice write an error to the response body and end execution
	if verseNumber > len(responses)-1 {
		http.Error(w, fmt.Sprintf(`{"error": index out of range, recieved %d indexable range upper limit %d}`, verseNumber, len(responses)-1), 404)
		return

	}

	// write data to the response body
	w.Write(responses[verseNumber])

	// write response status code
	w.WriteHeader(http.StatusOK)

	fmt.Println("\nRequest URL: ", requestURL)

	fmt.Printf("\nHost: %v\n\nHost Name: %v\n\nPath: %v\n\nPort: %v\n\nQuery: %v",
		host,
		hostName,
		path,
		port,
		query,
	)

	// access request header
	requestHeader := r.Header

	fmt.Println("\nRequest Header: ", requestHeader)

	// Read request body
	requestBody := make([]byte, r.ContentLength)

	_, err = r.Body.Read(requestBody)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nRequest Body: ", requestBody)

	fmt.Println("\nRequest Body: ", r.Body)

}

// Default ServeMux

//   - a predefined http.ServeMux
//   - called by all http.Handle methods
