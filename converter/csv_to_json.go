package cmd

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var ConvertCmd = &cobra.Command{
	Use:   "convert",
	Short: "Converts a file from one format to another",
	Long:  `Converts a file from one format to another. Currently supports CSV to JSON`,
	Run:   convert,
}

func init() {
	ConvertCmd.Flags().StringP("input", "i", "", "Input file to convert")
	ConvertCmd.Flags().StringP("output", "o", "", "Output file to convert to")
	ConvertCmd.Flags().StringP("delimiter", "d", ",", "Delimiter used in the input file")
	ConvertCmd.MarkFlagRequired("input")
	ConvertCmd.MarkFlagRequired("output")

	viper.BindPFlag("input", ConvertCmd.Flags().Lookup("input"))
	viper.BindPFlag("output", ConvertCmd.Flags().Lookup("output"))
	viper.BindPFlag("delimiter", ConvertCmd.Flags().Lookup("delimiter"))
}

func convert(cmd *cobra.Command, args []string) {
	inputFile := viper.GetString("input")
	outputFile := viper.GetString("output")
	delimiter := viper.GetString("delimiter")


	inputExt := getFileExtension(inputFile)
	outputExt := getFileExtension(outputFile)

	if inputExt == "csv" && outputExt == "json" {
		convertCSVToJSON(inputFile, outputFile, delimiter)
	} else if inputExt == "json" && outputExt == "csv" {
		fmt.Printf("unsupported file conversion: %s to %s", inputExt, outputExt)
	} else {
		fmt.Printf("unsupported file conversion: %s to %s", inputExt, outputExt)
	}
}

func getFileExtension(filename string) string {
	ext := ""
	for i := len(filename) - 1; i >= 0; i-- {
		if filename[i] == '.' {
			break
		}
		ext = string(filename[i]) + ext
	}
	return ext
}

func convertCSVToJSON(inputFileName, outputFileName string, delimiter string) error  {
	// Validate the input file extension.
	ext := getFileExtension(inputFileName)
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