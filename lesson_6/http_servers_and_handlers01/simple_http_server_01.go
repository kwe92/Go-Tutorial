package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

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

//   - an adapter type whose input is a callback with the signature of func(w http.ResponseWriter, r *http.Request){...}
//   - functions that meet the signature criteria are transformed into http.Handler implementations
//   - http.HandlerFunc allows you to implement http.Handlers without defining a new struct
//     for each handler created to implement the ServeHTTP method
