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
	inputFile, err := os.Open("retrieveOutput.json")
	if err != nil {
		fmt.Println("Error opening input file:", err)
		return
	}
	defer inputFile.Close()

	outputFile, err := os.Create("singleLineOutput.json")
	if err != nil {
		fmt.Println("Error creating output file:", err)
		return
	}
	defer outputFile.Close()

	reader := bufio.NewReader(inputFile)
	writer := bufio.NewWriter(outputFile)
	defer writer.Flush()

	for {
		line, err := reader.ReadBytes('\n')
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println("Error reading input file:", err)
			return
		}

		// Remove any trailing newline character from the line
		line = bytes.TrimRight(line, "\n")

		var data struct {
			Result struct {
				Results []json.RawMessage `json:"results"`
			} `json:"result"`
		}

		if err := json.Unmarshal(line, &data); err != nil {
			fmt.Println("Error parsing JSON:", err)
			return
		}

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
