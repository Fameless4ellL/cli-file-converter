package converter

import (
	"fmt"

	"github.com/Fameless4ellL/cli-file-converter/converter/commands"
	"github.com/Fameless4ellL/cli-file-converter/converter/utils"
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

	inputExt := utils.GetFileExtension(inputFile)
	outputExt := utils.GetFileExtension(outputFile)

	// create a interface to recognize the type of file to convert
	if inputExt == "csv" && outputExt == "json" {
		commands.ConvertCSVToJSON(inputFile, outputFile, delimiter)
	} else if inputExt == "json" && outputExt == "csv" {
		commands.ConvertJSONToCSV(inputFile, outputFile, delimiter)
	} else {
		fmt.Printf("unsupported file conversion: %s to %s", inputExt, outputExt)
	}
}