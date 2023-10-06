package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", handler)

	fmt.Println("Server started successfully!")

	http.ListenAndServe(":3000", mux)
}

func handler(w http.ResponseWriter, r *http.Request) {
	var responseData []byte

	responseData = []byte("Every kingdom divided against itself is brought to desolation, and every city or house divided against itself shall not stand.")

	w.Write(responseData)
}

// HandleFunc | Method | ServeMux

//   - a method that registers a route handler to a given pattern `path / URL`
//   - requires the least amount of code to register route handlers to patterns
//   - calls http.HandlerFunc within its implementation to
//     transform callbacks with the proper signature into http.Handlers
