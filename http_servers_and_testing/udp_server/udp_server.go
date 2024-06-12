package main

import (
	"fmt"
	"math/rand"
	"net"
	"os"
	"strconv"
	"strings"
)

// TODO: Review and explain code

var (
	PROTOCOL = "udp4"
)

func main() {

	arguments := os.Args

	if len(arguments) == 1 {
		fmt.Println("Provide a port number.")
		return
	}

	PORT := arguments[1]

	s, err := net.ResolveUDPAddr(PROTOCOL, PORT)

	if err != nil {
		fmt.Println(err)
		return
	}

	conn, err := net.ListenUDP(PROTOCOL, s)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer conn.Close()

	fmt.Println("UDP Server started listening on port", PORT)

	buffer := make([]byte, 2048)

	for {
		n, addr, err := conn.ReadFromUDP(buffer)

		if err != nil {
			fmt.Println(err)
			return
		}

		requestData := string(buffer[0:n])

		// n - 1 for new line?
		fmt.Print("-> ", requestData)

		if strings.TrimSpace(requestData) == "STOP" {
			fmt.Print("Exiting UDP Server")
			return
		}

		randomNumberResponse := strconv.Itoa(random(1, 1001))

		_, err = conn.WriteToUDP([]byte(randomNumberResponse), addr)

		if err != nil {
			fmt.Println(err)
			return
		}

	}

}

func random(min, max int) int {
	return rand.Intn(max-min) + min

}

// What is UDP?

//   - a connectionless network communication protocol between a client and a server

//   - data packets (datagrams) are sent individually with no guarantee of delivery

//   - unlike TCP their is no retransmission of lost data packets

//   - data packets are also unordered

//   - less reliable but faster than TCP as there is less setup under the hood and no error checking built-in

//   - there is no handshake or acknowledgement process

// When to Use a UDP Server over a TCP Server

//   - DNS lookup used UDP to resolve domain names to IP addresses

//   - Streaming services often use UDP for streaming video and audio

//   - Online gaming

//   - Live streaming

//   - Real time applications

//   - When you need speed over reliable order data transmission

// What is a UDP Server?

//   - a program or process that uses UDP to listen for and respond to requests
//     over a specified port from clients providing them some service

// How UDP Servers Work

//   - socket createion and port binding

//   - continously wait (infinite loop e.g. while loop) for incoming client UDP datagrams (packets)

//   - receive and process incoming datagram packets

//   - optionally send a response to the client

// Explain why Local and Remote Addresses are Handled Differently for UDP vs TCP

//   - UDP is a connectionless network communication protocol were datagrams (individual packets of data)
//     are sent and recieved over

//   - do to its connectionless nature UDP datagrams require both local and remote addresses
//     so the datagram knowns where it came from and who its going to

//   - with a TCP connection local and remote addresses are used to establish and maintain a connection
