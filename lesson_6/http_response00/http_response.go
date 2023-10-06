package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
)

const (
	Address = ":8080"
)

func main() {

	mux := setUpMux()

	fmt.Println("Server has started successfully!")

	log.Fatal(http.ListenAndServe(Address, mux))

}

func setUpMux() *http.ServeMux {

	mux := http.NewServeMux()

	mux.Handle("/", &home{})

	return mux
}

type home struct{}

func (h *home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	resp := &http.Response{
		StatusCode: http.StatusOK,
		Proto:      "HTTP/1.0",
		Header:     make(http.Header),
		Body:       io.NopCloser(strings.NewReader("form is form, emptiness is emptiness.")),
		// Body:       io.NopCloser(bytes.NewBufferString("form is form, emptiness is emptiness.")),
	}

	resp.Header.Add("Content-Type", "text/plain")

	resp.Write(w)
}

// http.Response | struct | http package

//   - A struct that represents an HTTP response to an HTTP request
//   - Containts import fields of information such as:
//       - StatusCode | Header | Body | Request | TLS `Transport Layer Security`
//   - Implements important methods such as:
//     - Read | Write `the Write implementation writes the response to a ResponseWriter`

// io.NopCloser | Exported Function | io package

//   - a convenience function that turns an implementation of io.Reader into
//     an io.ReadCloser
//   - one use case depicted above is to transform an io.Reader
//     into an io.ReadCloser to assign it to an http.ResponseBody

// strings.NewReader | Exported Function | strings package

// NewReader returns a new Reader reading from the string passed in as an argument
// similar to bytes.NewBufferString with better efficiency
