package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"net/url"
	"strconv"
	constants "testing_handlers01/constants"
)

const (
	apiAddress = ":8081"
)

func main() {

	mux := setUpMux()

	fmt.Println("Server started successfully!")

	log.Fatal(http.ListenAndServe(apiAddress, mux))
}

// setUpMux registers handlers to patterns and returns a new multiplexer
func setUpMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc(constants.Endpoints.IsPrime, isPrimeHandler)

	return mux
}

// isPrimeHandler handles requests to `/check-is-prime`
func isPrimeHandler(w http.ResponseWriter, r *http.Request) {

	var queryParameters url.Values

	// parse query parameters of the request URL into a defined hash map url.Values
	queryParameters = r.URL.Query()

	fmt.Println("query parameters: ", queryParameters)

	// url.Values.Get returns an empty string if the key does not exist within the hash map
	numberParam := queryParameters.Get("number")

	// convert string representation of a number to int
	n, err := strconv.Atoi(numberParam)

	if err != nil {
		// reply to the request, writing the error string as the response content body
		http.Error(w, "invalid number", http.StatusBadRequest)
		return
	}

	fmt.Print(n)

	fmt.Fprint(w, strconv.FormatBool(isPrime(int64(n))))

}

// isPrime returns true or false if a number is prime
func isPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}
