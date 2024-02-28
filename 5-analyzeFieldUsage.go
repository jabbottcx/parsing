package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

func main() {
	filePath := "finalOutput.json"
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var totalRecords int
	fieldCounts := map[string]int{
		"message": 0,
		"body":    0,
		"text":    0,
		"Message": 0,
	}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var record map[string]interface{}
		err := json.Unmarshal(scanner.Bytes(), &record)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue // Skip records that cannot be parsed
		}
		totalRecords++

		for field := range fieldCounts {
			if _, ok := record[field]; ok {
				fieldCounts[field]++
			}
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	output := "Field Analysis Summary:\n"
	for field, count := range fieldCounts {
		percent := (float64(count) / float64(totalRecords)) * 100
		output += fmt.Sprintf("%s: %d records, %.2f%% of all records\n", field, count, percent)
	}

	err = os.WriteFile("finalOutputFieldAnalysis.txt", []byte(output), 0644)
	if err != nil {
		panic(err)
	}
}
