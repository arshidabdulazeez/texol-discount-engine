package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Check for required arguments
	if len(os.Args) < 3 {
		fmt.Println("Usage: go-cli <order_total> <customer_type>")
		os.Exit(1)
	}

	// Parse inputs
	orderTotal := os.Args[1]
	customerType := os.Args[2]
	url := "http://localhost:8080/apply-discount"

	// Create the request payload
	payload := map[string]string{
		"order_total":   orderTotal,
		"customer_type": customerType,
	}
	payloadBytes, err := json.Marshal(payload)
	if err != nil {
		fmt.Printf("Error creating JSON payload: %v\n", err)
		os.Exit(1)
	}

	// Make the POST request
	resp, err := http.Post(url, "application/json", bytes.NewBuffer(payloadBytes))
	if err != nil {
		fmt.Printf("Error making POST request: %v\n", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	// Print the response
	fmt.Println("Response Status:", resp.Status)
	fmt.Println("Response Body:")
	_, err = fmt.Fprintln(os.Stdout, resp.Body)
	if err != nil {
		fmt.Printf("Error reading response: %v\n", err)
	}
}
