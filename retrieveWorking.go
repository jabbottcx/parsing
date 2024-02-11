package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

// Config represents the structure of the configuration file
type Config struct {
	URL        string `json:"url"`
	PrivateKey string `json:"privateKey"`
}

func main() {
	// Load configuration from retrieveConfig.json
	var config Config
	configFile, err := os.ReadFile("retrieveConfig")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(configFile, &config)
	if err != nil {
		panic(err)
	}

	// The query payload as per the Direct Archive Query API documentation
	// Adjust the query according to your needs
	query := map[string]interface{}{
		"query": "source logs",
		"metadata": map[string]interface{}{
			"tier":                   "TIER_ARCHIVE",
			"syntax":                 "QUERY_SYNTAX_DATAPRIME",
			"limit":                  50000,
			"startDate":              "2024-02-01T00:00:00Z",
			"endDate":                "2024-02-01T23:59:59Z",
			"defaultSource":          "string",
			"strictFieldsValidation": true,
		},
	}

	// Convert the query to a JSON byte slice
	requestBody, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	// Create a new HTTP request using the URL from the configuration file
	req, err := http.NewRequest("POST", config.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}

	// Set the necessary headers, including the private key from the configuration file
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.PrivateKey)
	req.Header.Set("Accept", "application/json")

	// Make the HTTP request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Read the response body
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Print the response status and body to standard output
	fmt.Println("Response status:", resp.Status)
	fmt.Println("Response body:", string(responseBody))

	// Output the response body to a file named retrieveOutput.json
	file, err := os.Create("retrieveOutput.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(responseBody)
	if err != nil {
		panic(err)
	}

	fmt.Println("The response has been saved to retrieveOutput.json")
}
