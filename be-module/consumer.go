// be-module/consumer.go
package main

import (
	"fmt"

	"github.com/sandeep-musale/golangpoc/shared-module"
)

func startConsumer() {
	// Implementation to consume messages from ActiveMQ Artemis queue
}

func processQueueMessage(message []byte) {
	var event shared.SessionEvent
	// Unmarshal message to SessionEvent
	// ...

	// Write to Oracle DB
	writeToOracleDB(event)

	// Create XML file
	createXMLFile(event)
}

func writeToOracleDB(event shared.SessionEvent) {
	// Implementation to write to Oracle DB
}

func createXMLFile(event shared.SessionEvent) {
	// Implementation to create XML file
}

func main() {
	startConsumer()
}