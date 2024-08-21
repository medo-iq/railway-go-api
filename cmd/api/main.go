package main

import (
	"encoding/json"
	"fmt"
	"log"
)

func main() {
	// Define the JSON data directly
	jsonData := `{
		"id": 1,
		"name": "John Doe",
		"email": "john.doe@example.com"
	}`

	// Unmarshal the JSON data into a map
	var data map[string]interface{}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		log.Fatalf("Error unmarshaling JSON: %v", err)
	}

	// Print the data
	fmt.Println("Decoded JSON Data:", data)
	
	// Logging information
	logFile, err := os.OpenFile("app.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		log.Fatalf("Error opening log file: %v", err)
	}
	defer logFile.Close()
	logger := log.New(logFile, "INFO: ", log.Ldate|log.Ltime|log.Lshortfile)
	logger.Println("JSON data processed successfully")
}
