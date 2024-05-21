package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// TODO: Review

// NON-CONCURRENT TCP SERVER

func main() {
	// command line arguments
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Provide a port number to listen over.")
		return
	}

	PORT := arguments[1]

	// start listening for incoming connection requests on the port specified over a network stream
	listener, err := net.Listen("tcp", PORT)

	if err != nil {
		fmt.Printf("There was an error when attempting to listen on port: %s | Error: %s", PORT, err.Error())
		return
	}

	fmt.Println("TCP server started on listening on:", PORT)

	// close the listener when the server is shutdown
	defer listener.Close()

	// blocks (pauses) program waiting for incoming client connections
	// accepted connections are handled outside of the for-loop because this server is not concurrent
	conn, err := listener.Accept()

	if err != nil {
		fmt.Println(err)
		return
	}

	// handle a successful client connection

	for {

		// read the request data
		clientData, err := bufio.NewReader(conn).ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		// stop the serber if the request data contains the word `STOP`
		if strings.TrimSpace(clientData) == "STOP" {
			fmt.Println("Exiting TCP Server!")
			return
		}

		// output the request data to the standard output
		fmt.Println("->: " + clientData)

		now := time.Now()

		formatedTime := now.Format(time.RFC3339) + "\n"

		// send the current time back to the client
		fmt.Fprint(conn, formatedTime)
	}
}
