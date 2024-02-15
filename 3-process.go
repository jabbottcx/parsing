package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	// Open the input file.
	inputFile, err := os.Open("singleLineOutput")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	// Create the output file.
	outputFile, err := os.Create("FinalOutput.json")
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

		// Unescape the userData JSON string. This automatically handles unicode characters.
		var userData interface{}
		if err := json.Unmarshal([]byte(userDataJSON), &userData); err != nil {
			fmt.Println("Error unescaping userData:", err)
			continue
		}

		// Use json.Encoder to write userData directly to the output file, disabling HTML escaping.
		encoder := json.NewEncoder(writer)
		encoder.SetEscapeHTML(false) // Disable HTML escaping
		if err := encoder.Encode(userData); err != nil {
			fmt.Println("Error encoding userData to JSON:", err)
			continue
		}
		// Note: Encoder.Encode adds its own newline, so you don't need to add another.
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading from input file:", err)
	}
}
