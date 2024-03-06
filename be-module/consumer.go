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
	var event shared.sessionevent
	// Unmarshal message to sessionevent
	// ...

	// Write to Oracle DB
	writeToOracleDB(event)

	// Create XML file
	createXMLFile(event)
}

func writeToOracleDB(event shared.sessionevent) {
	// Implementation to write to Oracle DB
}

func createXMLFile(event shared.sessionevent) {
	// Implementation to create XML file
}

func main() {
	startConsumer()
}