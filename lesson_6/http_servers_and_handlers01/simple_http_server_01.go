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

//   - an adapter whose input is a callback with the signature of func(w http.ResponseWriter, r *http.Request){...}
//   - if the function meets the signature criteria it is transformed into an http.Handler implementation
//   - this adapter pattern allows you to implement http.Handlers with less
//     boiler plate code
