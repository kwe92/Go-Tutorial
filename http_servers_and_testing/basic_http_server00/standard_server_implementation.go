package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var (
	ADDRESS = "127.0.0.1:8080"
)

func main() {

	fs := http.FileServer(http.Dir("./static"))

	standardServer := NewStandardServer(ADDRESS)

	// home path
	http.Handle("/", fs)

	// myHandler path
	http.Handle("/myHandler", &myHandler{})

	log.Printf("Started the server successfully on: %s\n", ADDRESS)

	log.Fatalf("error when attampting to start the server: %s\n", standardServer.ListenAndServe().Error())

}

func NewStandardServer(address string) *http.Server {

	server := &http.Server{
		Addr: address,
		// Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	return server

}

type myHandler struct{}

func (h *myHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "This is just a standard server")
}
