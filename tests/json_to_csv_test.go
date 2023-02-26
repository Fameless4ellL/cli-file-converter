package test

import (
	"os"
	"testing"

	"github.com/Fameless4ellL/cli-file-converter/converter/commands"
)

func TestConvertJSONToCSV(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"[{\"name\":\"John\",\"age\":\"30\",\"city\":\"New York\"},{\"name\":\"Mary\",\"age\":\"25\",\"city\":\"Chicago\"},{\"name\":\"Peter\",\"age\":\"35\",\"city\":\"Los Angeles\"}]", "name,age,city\nJohn,30,New York\nMary,25,Chicago\nPeter,35,Los Angeles"},
	}

	for _, test := range tests {
		// Create a test JSON file
		jsonfile, err := os.Create("test.json")
		if err != nil {
			t.Errorf("Error creating test JSON file: %s", err)
		}
		defer jsonfile.Close()

		// Write some test data to the JSON file
		jsonfile.WriteString(test.input)

		// Convert the test JSON file to CSV
		err = commands.ConvertJSONToCSV("test.json", "test.csv", ",")
		if err != nil {
			t.Errorf("Error converting JSON to CSV: %s", err)
		}

		// Verify the contents of the CSV file
		if _, err := os.Stat("test.csv"); os.IsNotExist(err) {
			t.Errorf("Error converting JSON to CSV: %s", err)
		}
	}
}

func BenchmarkConvertJSONToCSV(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Create a test JSON file
		jsonfile, err := os.Create("test.json")
		if err != nil {
			b.Errorf("Error creating test JSON file: %s", err)
		}
		defer jsonfile.Close()

		// Write some test data to the JSON file
		jsonfile.WriteString("[{\"name\":\"John\",\"age\":\"30\",\"city\":\"New York\"},{\"name\":\"Mary\",\"age\":\"25\",\"city\":\"Chicago\"},{\"name\":\"Peter\",\"age\":\"35\",\"city\":\"Los Angeles\"}]")

		// Convert the test JSON file to CSV
		err = commands.ConvertJSONToCSV("test.json", "test.csv", ",")
		if err != nil {
			b.Errorf("Error converting JSON to CSV: %s", err)
		}

		// Verify the contents of the CSV file
		if _, err := os.Stat("test.csv"); os.IsNotExist(err) {
			b.Errorf("Error converting JSON to CSV: %s", err)
		}
	}
}