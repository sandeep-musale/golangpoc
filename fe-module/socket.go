// fe-module/socket.go
package main

import (
	"encoding/json"
	"fmt"
	"net"

	"github.com/sandeep-musale/golangpoc/shared-module"
)

const (
	// SimulatorSocketPort is the port for simulator communication
	SimulatorSocketPort = "8080"
)

func startSocketServer() {
	listener, err := net.Listen("tcp", ":"+SimulatorSocketPort)
	if err != nil {
		fmt.Println("Error starting socket server:", err)
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}

		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	for {
		n, err := conn.Read(buffer)
		if err != nil {
			fmt.Println("Error reading from socket:", err)
			return
		}

		data := buffer[:n]
		go processSocketData(data)
	}
}

func processSocketData(data []byte) {
	var event shared.SessionEvent
	err := json.Unmarshal(data, &event)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Send to ActiveMQ Artemis
	sendToQueue(event)
}

func sendToQueue(event shared.SessionEvent) {
	// Implementation to send event to ActiveMQ Artemis queue
}

func main() {
	startSocketServer()
}