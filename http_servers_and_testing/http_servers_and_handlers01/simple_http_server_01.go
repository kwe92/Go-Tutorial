package main

import (
	"fmt"
	"net/http"
)

func main() {

	// create a multiplexer instance
	mux := http.NewServeMux()

	// register the hander to thea given pattern
	mux.Handle("/", http.HandlerFunc(hanlder))

	fmt.Println("server has started.")

	http.ListenAndServe(":8082", mux)

}

func hanlder(w http.ResponseWriter, r *http.Request) {
	var responseData []byte

	responseData = []byte("In him was life; and the life was the light of men.\n\nAnd the light shineth in the darkness; and the darkness comprehended it not.")

	w.Write(responseData)
}

// http.HandlerFunc

//   - an adapter type that convert a function to an http.Handler implementation
//   - implement an http.Handler without defining a new struct
//     for each handler created to implement the ServeHTTP method explicitly
