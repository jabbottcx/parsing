package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"time"
)

// Config represents the expected structure of the retrieveConfig.json file
type Config struct {
	URL          string `json:"url"`
	PrivateKey   string `json:"privateKey"`
	CustomerName string `json:"customerName"`
	QuerySegment struct {
		Query     string `json:"query"`
		Tier      string `json:"tier"`
		Syntax    string `json:"syntax"`
		Limit     int    `json:"limit"`
		StartDate string `json:"startDate,omitempty"`
		EndDate   string `json:"endDate,omitempty"`
	} `json:"querySegment"`
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

	// Default startDate and endDate if not provided
	currentTime := time.Now().UTC()
	if config.QuerySegment.StartDate == "" {
		config.QuerySegment.StartDate = currentTime.Add(-12 * time.Minute).Format(time.RFC3339)
	}
	if config.QuerySegment.EndDate == "" {
		config.QuerySegment.EndDate = currentTime.Add(2 * time.Minute).Format(time.RFC3339)
	}

	query := map[string]interface{}{
		"query": config.QuerySegment.Query,
		"metadata": map[string]interface{}{
			"tier":                   config.QuerySegment.Tier,
			"syntax":                 config.QuerySegment.Syntax,
			"limit":                  config.QuerySegment.Limit,
			"startDate":              config.QuerySegment.StartDate,
			"endDate":                config.QuerySegment.EndDate,
			"defaultSource":          "string",
			"strictFieldsValidation": true,
		},
	}

	// Convert the query to a JSON byte slice
	requestBody, err := json.Marshal(query)
	if err != nil {
		panic(err)
	}

	// Print the query for inspection
	fmt.Printf("Query being sent to the API: %s\n", requestBody)

	// Create and send the HTTP request
	req, err := http.NewRequest("POST", config.URL, bytes.NewBuffer(requestBody))
	if err != nil {
		panic(err)
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Authorization", "Bearer "+config.PrivateKey)
	req.Header.Set("Accept", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	// Handle the response
	responseBody, err := io.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}

	// Save the response to a file
	file, err := os.Create("retrieveOutput") // Changed this line to use a static file name
	if err != nil {
		panic(err)
	}
	defer file.Close()

	written, err := file.Write(responseBody)
	if err != nil {
		panic(err)
	}

	fmt.Printf("The response has been saved to retrieveOutput.json. A total of %d bytes were written.\n", written)
}
