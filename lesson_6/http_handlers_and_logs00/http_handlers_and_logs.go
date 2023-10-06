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

	mux := setUpMux()

	fmt.Println("Server has started successfully!")

	log.Fatal(http.ListenAndServe(Address, mux))

}

func logger(h func(http.ResponseWriter, *http.Request)) func(http.ResponseWriter, *http.Request) {

	handler := http.HandlerFunc(h)

	return func(w http.ResponseWriter, r *http.Request) {

		log.SetPrefix("Request Metadata: ")

		log.Printf("\n%s %s", r.Method, r.URL.Path)

		handler.ServeHTTP(w, r)
	}
}

var routeHandlers = map[string]func(w http.ResponseWriter, r *http.Request){
	"/":      homeHandler,
	"/about": aboutHandler,
}

func setUpMux() *http.ServeMux {

	mux := http.NewServeMux()

	for pattern, handler := range routeHandlers {
		mux.HandleFunc(pattern, logger(handler))

	}

	return mux
}

func homeHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain")

	w.Write(
		[]byte("Trust in the Lord with all thine heart," +
			" and lean not unto thine own understanding." +
			" In all thy was acknowledge him, and he shall direct thy paths."))

}

func aboutHandler(w http.ResponseWriter, r *http.Request) {

	w.Header().Add("Content-Type", "text/plain")

	w.Write([]byte("Begin, to begin is half the work let half still remain." +
		"Again begin this and thou wilt have finished.",
	))

}
