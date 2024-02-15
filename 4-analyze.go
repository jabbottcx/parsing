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

// PatternCount holds each pattern and its count.
type PatternCount struct {
	Pattern string `json:"pattern"`
	Count   int    `json:"count"`
}

func main() {
	deleteFileIfExists("singleLineOutput")
	deleteFileIfExists("retrieveOutput")

	fileName := "FinalOutput.json"
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

	var patterns []PatternCount
	for pattern, count := range patternCounts {
		patterns = append(patterns, PatternCount{Pattern: pattern, Count: count})
	}

	sort.Slice(patterns, func(i, j int) bool {
		return patterns[i].Count > patterns[j].Count
	})

	finalData, err := json.MarshalIndent(patterns, "", "  ")
	if err != nil {
		fmt.Println("Error marshaling JSON:", err)
		return
	}

	err = os.WriteFile("FinalAnalysis.json", finalData, 0644)
	if err != nil {
		fmt.Println("Error writing to file:", err)
		return
	}
}

func deleteFileIfExists(filename string) {
	if _, err := os.Stat(filename); err == nil {
		err := os.Remove(filename)
		if err != nil {
			fmt.Println("Error deleting file:", err)
		}
	}
}

func generatePattern(data interface{}) string {
	if data == nil {
		return "nil"
	}

	switch reflect.TypeOf(data).Kind() {
	case reflect.Map:
		mapData := data.(map[string]interface{})
		var keys []string
		for key := range mapData {
			keys = append(keys, key)
		}
		sort.Strings(keys)
		pattern := "{"
		for _, key := range keys {
			pattern += fmt.Sprintf("%s:%s,", key, generatePattern(mapData[key]))
		}
		return trimSuffix(pattern, ",") + "}"
	case reflect.Slice:
		sliceData := data.([]interface{})
		if len(sliceData) == 0 {
			return "[]"
		}
		pattern := "["
		for _, elem := range sliceData {
			pattern += generatePattern(elem) + ","
		}
		return trimSuffix(pattern, ",") + "]"
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
