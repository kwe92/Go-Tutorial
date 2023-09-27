package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.Handle("/", &home{})

	http.ListenAndServe(":8080", mux)

	fmt.Println("Server has started")
}

type home struct {
}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form is form, Emptiness is emptiness."))
}

// http.NewServerMux

//   - allocates a ServerMux struct in memory
//     as your own instance

// ServerMux

//   - an HTTP multiplexer
//   - register a collection of paths to handlers and matches them in O(n)
//   - incoming request URLs are matched to registered paths
//     and their associated route handler executes the route handler

// http.NewServeMux().Handle()

//   - register a pattern (URL / path) to a route handler (http.Handler implementation)

// ServeHTTP

//   - a method required to implement a http.Handler
//   - takes two arguments:
//       * http.ResponseWriter
//       * http.Request

// http.ListenAndServe

//   - Starts an HTTP server that listens on a TCP network address with a specified port and
//     calls Serve on the handler to handle all incoming HTTP requests on that address
//   - http.ListenAndServe can be called multiple times with diffrent addresses and handlers
