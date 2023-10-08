package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	// create an HTTP multiplexer instance
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	fmt.Println("Server started successfully!")

	log.Fatal(http.ListenAndServe(":3000", mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	var responseData []byte

	responseData = []byte("Every kingdom divided against itself is brought to desolation, and every city or house divided against itself shall not stand.")

	w.Write(responseData)
}

// HandleFunc | Method | ServeMux

//   - HandleFunc registers a route handler to a given pattern `path / URL`
//   - requires the least amount of code to register route handlers to patterns
//   - the HandleFunc implementation calls http.HandlerFunc on the callback passed
//     as an argument and transforms the function into an http.Handler
