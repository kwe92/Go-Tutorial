package main

import (
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	router.HandleFunc("/health", HealthCheckHandler)

	fmt.Println("Server Started!")

	log.Fatal(http.ListenAndServe(":8080", router))

}

func HealthCheckHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	w.WriteHeader(http.StatusOK)

	io.WriteString(w, `{"alive": true}`)
}
