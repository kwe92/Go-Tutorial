package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	Address = ":8080"
)

func main() {

	mux := setUpMux()

	// wrap your mux with the logger middleware
	wrappedMux := logger(mux)

	fmt.Println("Server has started successfully!")

	log.Fatal(http.ListenAndServe(Address, wrappedMux))

}

// logger: a middleware function that does request logging
func logger(h http.Handler) http.Handler {

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		// set logger prefix message
		log.SetPrefix("Request Metadata: ")

		log.Printf("\n%s %s", r.Method, r.URL.Path)

		h.ServeHTTP(w, r)
	})
}

var routeHandlers = map[string]func(w http.ResponseWriter, r *http.Request){
	"/":      homeHandler,
	"/about": aboutHandler,
}

// setUpMux: sets up the HTTP multiplexer
// and registers all handlers to patterns
func setUpMux() *http.ServeMux {

	mux := http.NewServeMux()

	for pattern, handler := range routeHandlers {
		mux.HandleFunc(pattern, handler)

	}

	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain")

	w.Write(
		[]byte("Trust in the Lord with all thine heart," +
			" and lean not unto thine own understanding." +
			" In all thy was acknowledge him, and he shall direct thy paths."))

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain")

	w.Write([]byte("Begin, to begin is half the work let half still remain." +
		"Again begin this and thou wilt have finished.",
	))

}

// Middleware Pattern

//   - a design pattern that enables pre- and/or post-processing of an HTTP request
//   - middleware sits between the GO HTTP server and the route handlers
//     adding functionality without modifying the servers core logic
//   - increasing modularity and assists in keeping code D.R.Y.
//   - can be implemented as a function or a struct
//      that wraps your handers or multiplexer, perferably the multiplexer
//   - useful for but not limited to:
//       ~ user authentication
//       ~ zipping `compressing` request or response body
//       ~ logging
//       ~ adding HTTP Headers

// Middleware Handlers in GO

//   - middleware implementations should:

//       - take an http.Handler as an argument
//       - return an http.Handler
//       - call the passed in handlers ServeHTTP method with the http.ResponseWriter
//         and *http.Request passed as arguments
