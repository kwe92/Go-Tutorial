package main

import (
	"log"
	"net/http"
	"time"
)

const (
	Addr = ":8080"
)

func main() {

	mux := setUpMux()

	wrappedMux := NewLogger(mux)

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

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	aboutResponseBody := []byte("And the light shineth in darkness; and the darkness comprehended it not.")

	w.Write(aboutResponseBody)

}

// Logger is a middleware hanlder that does request logging
type Logger struct {
	Handler http.Handler
}

// ServeHTTP handlers responding to requests and logging request details
func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	start := time.Now()

	l.Handler.ServeHTTP(w, r)

	log.Printf("%s %s %v", r.Method, r.URL.Path, time.Since(start))

}

// NewLogger constructs a new Logger middleware handler
func NewLogger(handlerToWrap http.Handler) *Logger {
	return &Logger{handlerToWrap}
}

type AddHeaders struct {
	handler http.Handler
}
