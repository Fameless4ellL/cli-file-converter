# CLI File Converter
![ci](https://github.com/Fameless4ellL/cli-file-converter/actions/workflows/ci.yml/badge.svg) [![Go Report Card](https://goreportcard.com/badge/github.com/Fameless4ellL/cli-file-converter)](https://goreportcard.com/report/github.com/Fameless4ellL/cli-file-converter)

This is a command-line interface (CLI) tool for converting files between CSV and JSON formats. The tool is written in Go and uses the `cobra` and `viper` packages for command-line parsing and configuration management.

## Installation

To use the CLI file converter, you'll need to have Go installed on your machine. You can then install the tool using the following command:

`go install github.com/Fameless4ellL/cli-file-converter`


## Usage

To convert a file from CSV to JSON, use the `convert` command:

`cli-file-converter convert -i input.csv -o output.json`


By default, the converter assumes that CSV files use a comma (`,`) as the delimiter. You can specify a different delimiter using the `-d` or `--delimiter` flag:

`cli-file-converter convert -i input.csv -o output.json -d ";"`


For a full list of available commands and options, use the `--help` flag:

`cli-file-converter convert --help`


