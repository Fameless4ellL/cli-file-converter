package commands

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Fameless4ellL/cli-file-converter/converter/utils"
)

// ConvertCSVToJSON converts a CSV file to a JSON file.
func ConvertCSVToJSON(inputFileName, outputFileName string, delimiter string) error {
	// Validate the input file extension.
	ext := utils.GetFileExtension(inputFileName)
	if ext != "csv" {
		return fmt.Errorf("input file must be a CSV file")
	}
	// Open the input file.
	inputFile, err := os.Open(inputFileName)
	if err != nil {
		return fmt.Errorf("error opening input file: %v", err)
	}
	defer inputFile.Close()

	// Read the contents of the input file.
	reader := csv.NewReader(inputFile)
	reader.Comma = rune(delimiter[0])
	records, err := reader.ReadAll()
	if err != nil {
		return fmt.Errorf("error reading input file: %v", err)
	}

	// Convert the CSV records to JSON.
	// The first record is assumed to be the header.
	// The header is used as the key for each field in the JSON record.
	// The JSON records are stored in a slice.
	jsonData := make([]map[string]string, 0)
	header := records[0]
	for _, record := range records[1:] {
		jsonRecord := make(map[string]string)
		for i, field := range record {
			jsonRecord[header[i]] = field
		}
		jsonData = append(jsonData, jsonRecord)
	}
	fmt.Println(jsonData)

	// Write the JSON data to the output file.
	jsonBytes, err := json.MarshalIndent(jsonData, "", "  ")
	if err != nil {
		return fmt.Errorf("error marshalling JSON data: %v", err)
	}
	err = os.WriteFile(outputFileName, jsonBytes, 0644)
	if err != nil {
		return fmt.Errorf("error writing output file: %v", err)
	}

	fmt.Println("CSV to JSON conversion successful.")
	return nil
}
