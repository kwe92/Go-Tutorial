package main

import (
	"encoding/json"
	"fmt"
	"log"

	"net/http"
)

func main() {

	var quotes []Quote

	const (
		url string = "https://zenquotes.io/api/random"
	)
	// instantiate HTTP Client

	client := http.Client{}

	// use Client instance to send a GET request to the ZenQuotes API

	resp, err := client.Get(url)

	// check for any errors returned when sending GET request

	checkError(err)

	// create new Decoder object that reads from the io.Reader argument passed in

	decoder := json.NewDecoder(resp.Body)

	// call Decode method of decoder reading and storing the JSON-encoded value into the value pointed to

	decoder.Decode(&quotes)

	// print the result to console

	fmt.Printf("quote: %+v\n", quotes)

}

type Quote struct {
	Quote  string `json:"q"`
	Author string `json:"a"`
}

func checkError(err error) {
	if err != nil {
		log.Fatalf("error: %s", err.Error())
	}
}
