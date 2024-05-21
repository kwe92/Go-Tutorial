package main

import (
	"concurrent_tcp_server/model"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

// TODO: Review

func main() {
	// create new TCP server
	server, err := model.NewTcpServer(":8080")

	if err != nil {
		log.Fatalln(err.Error())
	}

	// start the server
	server.Start()

	// create a signal channel for graceful shutdown
	sigChan := make(chan os.Signal, 1)

	// listen and notify signal channel when an OS signal is triggered
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)

	// continuously wait (deadlock?) until an OS signal is read such as kill or 1
	<-sigChan

	fmt.Println("Shutting down server...")

	server.Stop()

	fmt.Println("Server stoped.")

}
