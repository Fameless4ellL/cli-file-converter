package commands

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/Fameless4ellL/cli-file-converter/converter/utils"
)

func ConvertJSONToCSV(inputFile, outputFile, delimiter string) error {
	// Validate the input file extension.
	ext := utils.GetFileExtension(inputFile)
	if ext != "json" {
		fmt.Errorf("input file must be a JSON file")
	}
	// Open the input file.
	input, err := os.Open(inputFile)
	if err != nil {
		fmt.Errorf("error opening input file: %v", err)
	}
	defer input.Close()

	// Read the contents of the input file.
	jsonData := make([]map[string]string, 0)
	err = json.NewDecoder(input).Decode(&jsonData)
	if err != nil {
		fmt.Errorf("error reading input file: %v", err)
	}

	// Convert the JSON records to CSV.
	// The first record is assumed to be the header.
	// The header is used as the key for each field in the CSV record.
	// The CSV records are stored in a slice.
	csvData := make([][]string, 0)
	header := make([]string, 0)
	for key := range jsonData[0] {
		header = append(header, key)
	}
	csvData = append(csvData, header)
	for _, record := range jsonData {
		csvRecord := make([]string, 0)
		for _, key := range header {
			csvRecord = append(csvRecord, record[key])
		}
		csvData = append(csvData, csvRecord)
	}

	// Write the CSV data to the output file.
	output, err := os.Create(outputFile)
	if err != nil {
		fmt.Errorf("error creating output file: %v", err)
	}
	defer output.Close()
	writer := csv.NewWriter(output)
	writer.Comma = rune(delimiter[0])
	err = writer.WriteAll(csvData)
	if err != nil {
		fmt.Errorf("error writing output file: %v", err)
	}

	fmt.Println("JSON to CSV conversion successful.")
	return nil
}
