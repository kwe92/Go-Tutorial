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

//   - an adapter that defines a function type with a ServeHTTP method
//     which takes as an argument a function that implements ServeHTTP
//     transforms the function into a http.Handler implementation
//   - this adapter pattern allows you to implement http.Handlers with less
//     boiler plate code
