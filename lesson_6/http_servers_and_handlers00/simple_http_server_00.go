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

// ServeHTTP handles the request and responds by writing to the http.ResponseWriter
func (h *homeHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	// add headers to the http.Header map
	w.Header().Add("Content-Type", "text/plain")

	// write the content of the header to the response with a status code
	w.WriteHeader(200)

	// write the response body
	w.Write([]byte("Form is form, Emptiness is emptiness."))

}

// HTTP Server | Concept

//   - a service that listens for incoming HTTP Requests and returns an HTTP Response
//   - requests are routed to different registered handlers `functions` depending on the request URL

// ServeMux | Struct

//   - a struct representing an HTTP multiplexer allowing the creation of HTTP web servers
//   - ServeMux has methods to:
//       ~ register a handler `function` to a pattern `URL`
//  	 ~ match an incoming request URL to a pattern with a registered handler
//   - if a request URL matches a pattern that paths registered handler is executed
//   - handlers are matched to patterns in O(n) `linear-time`

// http.NewServeMux | Exported Function

//   - allocates a new ServeMux struct in memory
//     which can be used as an instance instead of using the DefaultServeMux

// Handle | Method | ServeMux

//   - Handle registers a route handler to a given pattern `path / URL`
//   - the handler must be an implementation of http.Handler

// http.Handler | Interface | http package

//   - an interface that defines a handler object
//   - has one method signature ServeHTTP
//   - used to create custom handlers

// ServeHTTP | Method | http.Handler

//   - required to implement the http.Handler interface
//   - ServeHTTP serves the purpose of handling the request by:
//       ~ writing headers
//       ~ writing content to http.ResponseWriter `response object`
//       ~ checking the request
//       ~ parsing query parameters
//   - takes two arguments:
//       ~ http.ResponseWriter
//       ~ pointer to a http.Request struct

// http.ListenAndServe | Exported Function

//   - starts an HTTP server that listens on a TCP network address
//   - calls Serve on the handler to handle all incoming HTTP requests at that address
//   - http.ListenAndServe can be called multiple times with diffrent addresses and handlers
