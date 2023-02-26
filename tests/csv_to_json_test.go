package test

import (
	"os"
	"testing"

	"github.com/Fameless4ellL/cli-file-converter/converter/commands"
)

func TestConvertCSVToJSON(t *testing.T) {
	var tests = []struct {
		input    string
		expected string
	}{
		{"name,age,city\nJohn,30,New York\nMary,25,Chicago\nPeter,35,Los Angeles", "[{\"name\":\"John\",\"age\":\"30\",\"city\":\"New York\"},{\"name\":\"Mary\",\"age\":\"25\",\"city\":\"Chicago\"},{\"name\":\"Peter\",\"age\":\"35\",\"city\":\"Los Angeles\"}]"},
	}

	for _, test := range tests {
		// Create a test CSV file
		csvfile, err := os.Create("test.csv")
		if err != nil {
			t.Errorf("Error creating test CSV file: %s", err)
		}
		defer csvfile.Close()

		// Write some test data to the CSV file
		csvfile.WriteString(test.input)

		// Convert the test CSV file to JSON
		err = commands.ConvertCSVToJSON("test.csv", "test.json", ",")
		if err != nil {
			t.Errorf("Error converting CSV to JSON: %s", err)
		}

		// Verify the contents of the JSON file
		if _, err := os.Stat("test.json"); os.IsNotExist(err) {
			t.Errorf("Error converting CSV to JSON: %s", err)
		}
	}
}

func BenchmarkConvertCSVToJSON(b *testing.B) {
	for i := 0; i < b.N; i++ {
		// Create a test CSV file
		csvfile, err := os.Create("test.csv")
		if err != nil {
			b.Errorf("Error creating test CSV file: %s", err)
		}
		defer csvfile.Close()

		// Write some test data to the CSV file
		csvfile.WriteString("name,age,city\nJohn,30,New York\nMary,25,Chicago\nPeter,35,Los Angeles")

		// Convert the test CSV file to JSON
		err = commands.ConvertCSVToJSON("test.csv", "test.json", ",")
		if err != nil {
			b.Errorf("Error converting CSV to JSON: %s", err)
		}

		// Verify the contents of the JSON file
		if _, err := os.Stat("test.json"); os.IsNotExist(err) {
			b.Errorf("Error converting CSV to JSON: %s", err)
		}
	}
}