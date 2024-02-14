package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"reflect"
	"sort"
	"strings"
)

// Define a struct to hold each pattern and its count.
type PatternCount struct {
	Pattern string `json:"pattern"`
	Count   int    `json:"count"`
}

func main() {
	fileName := "finalOutput.json"
	file, err := os.Open(fileName)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	patternCounts := make(map[string]int)

	for scanner.Scan() {
		var record map[string]interface{}
		err := json.Unmarshal(scanner.Bytes(), &record)
		if err != nil {
			fmt.Println("Error parsing JSON:", err)
			continue
		}

		pattern := generatePattern(record)
		patternCounts[pattern]++
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Convert the map to a slice of PatternCount for sorting and JSON output.
	var patterns []PatternCount
	for pattern, count := range patternCounts {
		patterns = append(patterns, PatternCount{Pattern: pattern, Count: count})
	}

	// Sort the slice for consistent output.
	sort.Slice(patterns, func(i, j int) bool {
		return patterns[i].Count > patterns[j].Count // Sort by count, descending
	})

	// Marshal the slice into JSON.
	finalData, err := json.MarshalIndent(patterns, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	// Write the JSON to 'FinalAnalysis.json'.
	err = os.WriteFile("FinalAnalysis.json", finalData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func generatePattern(data interface{}) string {
	if data == nil {
		return "nil"
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		mapData, ok := data.(map[string]interface{})
		if !ok {
			return "invalid-map"
		}
		var keys []string
		for key := range mapData {
			keys = append(keys, key)
		}
		sort.Strings(keys) // Ensure consistent ordering
		pattern := "{"
		for _, key := range keys {
			pattern += fmt.Sprintf("%s:%s,", key, generatePattern(mapData[key]))
		}
		return trimSuffix(pattern, ",") + "}"
	case reflect.Slice:
		sliceData, ok := data.([]interface{})
		if !ok || len(sliceData) == 0 {
			return "[]"
		}
		return "[" + generatePattern(sliceData[0]) + "]"
	default:
		return "value"
	}
}

func trimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		return s[:len(s)-len(suffix)]
	}
	return s
}
