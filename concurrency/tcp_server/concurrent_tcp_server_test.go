package main

import (
	"concurrent_tcp_server/model"
	"fmt"
	"net"
	"testing"
	"time"
)

func TestTcpServer(t *testing.T) {

	// create a new server
	testServer, err := model.NewTcpServer(":8080")

	clientMsg := "this is a test"

	if err != nil {
		t.Fatalf("error in TestTcpServer: %s", err.Error())
	}

	testServer.Start()

	testClient, err := net.Dial("tcp", ":8080")

	if err != nil {
		t.Fatalf("error in TestTcpServer: %s", err.Error())
	}
	defer testClient.Close()

	fmt.Fprint(testClient, clientMsg)

	buff := make([]byte, 1024)

	_, err = testClient.Read(buff)

	if err != nil {
		t.Fatalf("error in TestTcpServer: %s", err.Error())
	}

	expected := string(buff)

	actual := string(buff)

	if actual != expected {
		t.Errorf("expected: %s| but got: %s|", expected, actual)
	}

	time.Sleep(4 * time.Second)

	testServer.Stop()

	fmt.Println()

}
