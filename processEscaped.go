package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open the input file.
	inputFile, err := os.Open("singleLineOutput.json")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file.
	outputFile, err := os.Create("finalOutput.json")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	scanner := bufio.NewScanner(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for scanner.Scan() {
		var logEntry map[string]interface{}

		// Parse the JSON line to a map to easily extract userData.
		line := scanner.Text()
		if err := json.Unmarshal([]byte(line), &logEntry); err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		// Extract userData as a JSON string.
		userDataJSON, ok := logEntry["userData"].(string)
		if !ok {
			fmt.Println("Error extracting userData or userData is not a string")
			continue
		}

		// Unescape the userData JSON string.
		var userData interface{}
		if err := json.Unmarshal([]byte(userDataJSON), &userData); err != nil {
			fmt.Println("Error unescaping userData:", err)
			continue
		}

		// Marshal the userData back to JSON to normalize it.
		userDataBytes, err := json.Marshal(userData)
		if err != nil {
			fmt.Println("Error marshaling userData to JSON:", err)
			continue
		}

		// Write the unescaped and normalized userData JSON to the output file.
		writer.Write(userDataBytes)
		writer.WriteString("\n") // Add a newline after each record for readability.
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input file:", err)
	}
}
