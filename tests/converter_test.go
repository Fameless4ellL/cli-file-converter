package test

import (
	"os"
	"testing"

	"github.com/Fameless4ellL/cli-file-converter/converter"
)

func TestConvertCSVToJSON(t *testing.T) {
	// Create a test CSV file
	csvfile, err := os.Create("test.csv")
	if err != nil {
		t.Errorf("Error creating test CSV file: %s", err)
	}
	defer csvfile.Close()

	// Write some test data to the CSV file
	csvfile.WriteString("name,age,city\n")
	csvfile.WriteString("John,30,New York\n")
	csvfile.WriteString("Mary,25,Chicago\n")
	csvfile.WriteString("Peter,35,Los Angeles")

	// Convert the test CSV file to JSON
	err = converter.ConvertCSVToJSON("test.csv", "test.json", ",")
	if err != nil {
		t.Errorf("Error converting CSV to JSON: %s", err)
	}

	// Verify the contents of the JSON file
	if _, err := os.Stat("test.json"); os.IsNotExist(err) {
		t.Errorf("Error converting CSV to JSON: %s", err)
	}
}
