package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

// TODO: review

// TODO: Add comments

var (
	PROTOCOL = "udp4"
)

func main() {
	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Provide a hostname:port.")
		return
	}

	ADDRESS := arguments[1]

	server, err := net.ResolveUDPAddr(PROTOCOL, ADDRESS)

	conn, err := net.DialUDP(PROTOCOL, nil, server)

	if err != nil {
		fmt.Println("there was in issue connecting to:", ADDRESS)
		return
	}

	fmt.Println("The UDP Server is:", conn.RemoteAddr().String())

	defer conn.Close()

	for {
		// create a new reader over the standard input
		reader := bufio.NewReader(os.Stdin)

		fmt.Print(">> ")

		userInput, err := reader.ReadString('\n')

		if err != nil {
			fmt.Println(err)
			return
		}

		_, err = conn.Write([]byte(userInput))

		if err != nil {
			fmt.Println(err)
			return
		}

		if strings.TrimSpace(userInput) == "STOP" {
			fmt.Print("Exiting UDP client...")
			return
		}

		buffer := make([]byte, 2048)

		n, _, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println("UDP Server reply:", string(buffer[0:n]))

	}

}

// TODO: What is a UDP Client?
