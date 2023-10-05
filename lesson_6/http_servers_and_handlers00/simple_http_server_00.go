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

	// declare and initialize ServeMux instance
	mux := http.NewServeMux()

	// register handler to pattern
	mux.Handle("/", &homeHandler{})

	fmt.Println("Server has started!")

	// start HTTP server on port 8080
	log.Fatal(http.ListenAndServe(Address, mux))

}

// homeHandler implements http.Handler interface
type homeHandler struct {
}

func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain")

	w.WriteHeader(200)

	w.Write([]byte("Form is form, Emptiness is emptiness."))

}

// HTTP Server | Concept

//   - a service that listens for incoming HTTP Requests and returns an HTTP Response

// ServeMux | Struct

//   - a struct that is an HTTP multiplexer allowing the creation of HTTP servers
//   - ServerMux has methods to register handlers to paths `URLs` and
//     match incoming request URLs to paths with registered handlers
//   - if a request URL matches a path the registered handler is executed
//   - routes are matched to paths in O(n) `linear-time`

// http.NewServeMux | Exported Function

//   - allocates a new ServeMux struct in memory
//     which can be used as an instance instead of using the defaultServeMux

// Handle | Method | ServeMux

//   - a method that registers a route handler to a given pattern `path / URL`

// http.Handler | Interface | http package

//   - an interface that defines a handler object
//   - has one method signature ServeHTTP
//   - used to create custom handlers

// ServeHTTP | Method | http.Handler

//   - required to implement the http.Handler interface
//   - ServeHTTP writes headers and content to the http.ResponseWriter
//   - takes two arguments:
//       * http.ResponseWriter
//       * pointer to a http.Request struct

// http.ListenAndServe | Exported Function

//   - Starts an HTTP server that listens on a TCP network address with a specified port and
//     calls Serve on the handler to handle all incoming HTTP requests on that address
//   - http.ListenAndServe can be called multiple times with diffrent addresses and handlers
