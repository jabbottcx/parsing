package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"os"
)

func main() {
	// The input file name is "retrieveOutput.json"
	inputFileName := "retrieveOutput"
	// The output file name is set to "singleLineOutput.json"
	outputFileName := "singleLineOutput"

	// Open the input file
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file
	outputFile, err := os.Create(outputFileName)
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	// Create a buffered reader and writer
	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for {
		// Read each line from the input file
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}

		// Trim any trailing newline character from the line
		line = bytes.TrimRight(line, "\n")

		// Define the structure to unmarshal each JSON line into
		var data struct {
			Result struct {
				Results []json.RawMessage `json:"results"`
			} `json:"result"`
		}

		// Unmarshal the JSON from the line
		if err := json.Unmarshal(line, &data); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

		// Write each JSON record to the output file, followed by a newline
		for _, record := range data.Result.Results {
			if _, err := writer.Write(record); err != nil {
				fmt.Println("Error writing to output file:", err)
				return
			}
			if err := writer.WriteByte('\n'); err != nil {
				fmt.Println("Error writing newline to output file:", err)
				return
			}
		}
	}
}
