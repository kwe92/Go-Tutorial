package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TODO: Review

func main() {
	// received command line arguments
	arguments := os.Args

	// ensure there was a host and port number passed in
	if len(arguments) == 1 {
		fmt.Println("Provide a host and port number.")
		return
	}

	// assign the hostname:portnumber to the constant variable ADDRESS
	ADDRESS := arguments[1]

	// send a request to connect to the server
	conn, err := net.Dial("tcp", ADDRESS)

	if err != nil {
		fmt.Printf("there was an error when attempting to connect to :%s", ADDRESS)
		return
	}

	for {
		// read user input from standard input
		reader := bufio.NewReader(os.Stdin)

		fmt.Println(">> ")

		// user input to string from standard input
		text, _ := reader.ReadString('\n')

		// write user input to the TCP server connected to
		fmt.Fprint(conn, text+"\n")

		// read server response data
		message, _ := bufio.NewReader(conn).ReadString('\n')

		fmt.Print("->: " + message)

		// terminate the reading of user input from the standard input if the user types `STOP`
		if strings.TrimSpace(text) == "STOP" {
			fmt.Println("TCP client exiting...")
			return
		}
	}

}
