package main

import (
	"fmt"
	"log"
	"math/big"
	"net/http"
	"strconv"
)

const (
	apiAddress = ":8081"
)

func main() {
	mux := setUpMux()

	log.Fatal(http.ListenAndServe(apiAddress, mux))
}

func setUpMux() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/check-is-prime", isPrimeHandler)

	return mux
}

func isPrimeHandler(w http.ResponseWriter, r *http.Request) {
	number := r.URL.Query().Get("number")
	n, err := strconv.Atoi(number)

	if err != nil {
		http.Error(w, err.Error(), 404)
		return
	}

	fmt.Print(n)

	fmt.Fprint(w, strconv.FormatBool(isPrime(int64(n))))

}

func isPrime(n int64) bool {
	return big.NewInt(n).ProbablyPrime(0)
}
