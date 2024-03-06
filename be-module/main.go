package main

import (
	"fmt"
	"encoding/xml"
	"encoding/json"
	"os"
		
	"github.com/sandeep-musale/golangpoc/shared-module"
)

const (
	// OracleDBConnectionString is the connection string for Oracle DB
	OracleDBConnectionString = "MC:oracle@tcp(172.18.8.55:1521)/mcngdb"
)


func startConsumer() {
	// Implementation to consume messages from ActiveMQ Artemis queue
	// ...

	// Example: Dummy data for testing
	message := []byte(`{"LIID": "123", "CIN": "456", "StartDate": "2022-03-07T12:34:56Z"}`)
	processQueueMessage(message)
}

func processQueueMessage(message []byte) {
	var event shared.SessionEvent
	// Unmarshal message to SessionEvent
	err := json.Unmarshal(message, &event)
	if err != nil {
		fmt.Println("Error unmarshalling JSON:", err)
		return
	}

	// Write to Oracle DB
	// writeToOracleDB(event)

	// Create XML file
	createXMLFile("output.xml", event)
}

func createXMLFile(filename string, event shared.SessionEvent) {
	// Example: Print the received event
	fmt.Printf("Creating XML File - LIID: %s, CIN: %s, StartDate: %s\n", event.LIID, event.CIN, event.StartDate)
	
	file, err := os.Create(filename)
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(event); err != nil {
		fmt.Println("Error encoding event:", err)
		return
	}

	fmt.Println("XML file created:", filename)
	return
}

func main() {
	startConsumer()
}