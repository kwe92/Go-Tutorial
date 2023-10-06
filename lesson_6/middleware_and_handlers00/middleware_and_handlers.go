package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

const (
	Addr = ":8080"
)

func main() {

	mux := setUpMux()

	wrappedMux := setUpMiddleware(mux)

	log.Printf("server is listening at %s", Addr)

	log.Fatal(http.ListenAndServe(Addr, wrappedMux))
}

// setUpMux initalizes a new HTTP multiplexer, registers handlers to patterns and returns the mux
func setUpMux() *http.ServeMux {

	mux := http.NewServeMux()

	for pattern, handler := range routeHandlers {
		mux.HandleFunc(pattern, handler)
	}

	return mux
}

// a hash map of route handlers
var routeHandlers = map[string]func(w http.ResponseWriter, r *http.Request){
	"/":      homeHandler,
	"/about": aboutHandler,
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	homeResponseBody := []byte("The man who thinks he can and the man who thinks he can't are both right.")

	w.Write(homeResponseBody)

	fmt.Println("Response headers: ", w.Header())

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	aboutResponseBody := "And the light shineth in darkness; and the darkness comprehended it not."

	fmt.Fprintf(w, aboutResponseBody)

	fmt.Println("Response headers: ", w.Header())

}

// setUpMiddleware wraps the mux with middleware handlers
func setUpMiddleware(mux *http.ServeMux) http.Handler {
	return NewLogger(NewResponseHeader(mux, "Content-Type", "text/plain"))
}

// Logger is a middleware hanlder that does request logging
type Logger struct {
	Handler http.Handler
}

// ServeHTTP handler responding to requests and logging request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	l.Handler.ServeHTTP(w, r)

	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))

}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

// ResponseHeader is a middleware handler toadd a header to the response
type ResponseHeader struct {
	Handler     http.Handler
	HeaderName  string
	HeaderValue string
}

// ServeHTTP handles the request by adding the response header
func (resp *ResponseHeader) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	w.Header().Add(resp.HeaderName, resp.HeaderValue)

	resp.Handler.ServeHTTP(w, r)

}

// NewResponseHeader constructs a new ResponseHeader middleware handler
func NewResponseHeader(handler http.Handler, headerName string, headerValue string) *ResponseHeader {

	return &ResponseHeader{handler, headerName, headerValue}

}

// Chaining Middleware

//   - you can chain middleware with defined structs that implement ServeHTTP
//   - each middleware struct should have a constructor function to return a pointer to an instance of it
//   - there is also the option of using the adapter pattern

//
