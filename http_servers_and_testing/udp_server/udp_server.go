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

// TODO: What is UDP?

// TODO: What is a UDP Server?

// TODO: Explain why UDP methods like conn.WriteToUDP are used instead of conn.Write

// TODO: What is ResolveUDPAddr and why do we need it to connect and to listen? TCP only has one mechnism it seems like

// TODO: explain why Local Address and Remote Address are important for UDP
