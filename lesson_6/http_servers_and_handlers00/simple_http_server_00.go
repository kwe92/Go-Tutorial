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

	// declare and initialize a ServeMux instance
	mux := http.NewServeMux()

	// register handler to path
	mux.Handle("/", &home{})

	fmt.Println("Server has started!")

	// start HTTP server on port 8080 using the ServeMux instance as a handler
	log.Fatal(http.ListenAndServe(Address, mux))

}

type home struct {
}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Form is form, Emptiness is emptiness."))
}

// ServerMux

//   - a struct that is an HTTP multiplexer allowing the creation of HTTP servers
//   - ServerMux has methods to register handlers to paths `URLs` and
//     match incoming request URLs to paths with registered handlers
//   - if a request URL matches a path the registered handler is executed
//   - routes are matched to paths in O(n) `linear-time`

// http.NewServeMux

//   - allocates a new ServeMux struct in memory
//     as your own instance

// ServeMux.Handle(path, handler)

//   - a method that registers a route handler for a given path `patten / URL`

// ServeHTTP method

//   - required to implement the http.Handler interface
//   - ServeHTTP is responsible for writing headers and content to the http.ResponseWriter
//   - takes two arguments:
//       * http.ResponseWriter
//       * pointer to a http.Request struct

// http.ListenAndServe

//   - Starts an HTTP server that listens on a TCP network address with a specified port and
//     calls Serve on the handler to handle all incoming HTTP requests on that address
//   - http.ListenAndServe can be called multiple times with diffrent addresses and handlers
