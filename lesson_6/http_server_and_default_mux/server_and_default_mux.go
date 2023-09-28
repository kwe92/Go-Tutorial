// TODO: Add comments

package main

import (
	"fmt"
	"net/http"
)

func main() {

	// register handler to path
	http.HandleFunc("/", homeRouteHandler)

	// start HTTP server on port 8080, using the default ServeMux
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	} else if err == nil {
		fmt.Println("Server has started successfully!")
	}

}

// homeRouteHandler handles requests to the home path.
func homeRouteHandler(w http.ResponseWriter, r *http.Request) {

	// set response headers
	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set("X-Custom-Header", "In the beginning was the Word.")

	var responseData []byte

	responseData = []byte("Why are ye fearful, O ye of little faith?\n\nThen he arose, and rebuked the winds and the sea; and there was a great calm.")

	// write data to the response body
	w.Write(responseData)

	// write response status code
	w.WriteHeader(http.StatusOK)

	// access request URL
	requestURL := r.URL

	host := r.URL.Host

	hostName := r.URL.Hostname()

	path := r.URL.Path

	port := r.URL.Port()

	query := r.URL.Query()

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

	_, err := r.Body.Read(requestBody)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nRequest Body: ", requestBody)

	fmt.Println("\nRequest Body: ", r.Body)

}

// Default ServeMux

//   - a predefined http.ServeMux
//   - called by all http.Handle methods
