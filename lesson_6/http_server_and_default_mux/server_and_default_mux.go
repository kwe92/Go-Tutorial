package main

import (
	"fmt"
	"net/http"
)

func main() {

	http.HandleFunc("/", homeRouteHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		panic(err)
	} else if err == nil {
		fmt.Println("Server has started successfully!")
	}

}

func homeRouteHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "text/plain")

	w.Header().Set("X-Custom-Header", "In the beginning was the Word.")

	var responseData []byte

	responseData = []byte("Why are ye fearful, O ye of little faith?\n\nThen he arose, and rebuked the winds and the sea; and there was a great calm.")

	w.Write(responseData)

	w.WriteHeader(http.StatusOK)

	host := r.URL.Host

	hostName := r.URL.Hostname()

	path := r.URL.Path

	port := r.URL.Port()

	query := r.URL.Query()

	fmt.Printf("\nHost: %v\n\nHost Name: %v\n\nPath: %v\n\nPort: %v\n\nQuery: %v",
		host,
		hostName,
		path,
		port,
		query,
	)

	requestHeader := r.Header

	fmt.Println("\nRequest Header: ", requestHeader)

	requestBody := make([]byte, r.ContentLength)

	_, err := r.Body.Read(requestBody)

	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("\nRequest Body: ", requestBody)

	fmt.Println("\nRequest Body: ", r.Body)

}
