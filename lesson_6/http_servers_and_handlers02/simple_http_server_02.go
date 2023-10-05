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

// ServeMux{}.HandleFunc()

//   - registers a pattern `path` to a route handler
//   - requires the least amount of boiler plate code to
//     register route handlers to patterns
//   - http.HandlerFunc is called inside ServeMux{}.HandleFunc()
