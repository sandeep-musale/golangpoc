// be-module/filewriter.go
package main

import (
	"encoding/xml"
	"fmt"
	"os"
	"time"
	
	"github.com/sandeep-musale/golangpoc/shared-module"
)

func writeXMLFile(filename string, event shared.SessionEvent) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	encoder := xml.NewEncoder(file)
	encoder.Indent("", "  ")
	if err := encoder.Encode(event); err != nil {
		return err
	}

	fmt.Println("XML file created:", filename)
	return nil
}

func main() {
	// Example usage
	event := shared.SessionEvent{
		liid:      "123",
		cin:       "456",
		startdate: time.Now().UTC(),
	}

	err := writeXMLFile("output.xml", event)
	if err != nil {
		fmt.Println("Error creating XML file:", err)
	}
}