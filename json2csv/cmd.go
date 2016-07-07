package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/ashwanthkumar/indix-api-tools/util"
	"github.com/spf13/cobra"
)

var rowDelimiter = "\n"
var fieldDelimiter = ","
var multiValueDelimiter = "|"

// JSON2Csv is used to convert bulk output JSON to CSV
var JSON2Csv = &cobra.Command{
	Use:   "json2csv",
	Short: "Used to convert Bulk Job's JSON output to CSV",
	Long:  `Used to convert Bulk Job's JSON output to CSV`,
	Run:   util.AttachHandler(doJSONToCsv),
}

func doJSONToCsv(args []string) error {
	log.Printf("Starting json2csv %v\n", Version)
	if len(args) != 2 {
		return fmt.Errorf(
			`missing input source file and output source file
Usage: json2csv /path/to/input.json /path/to/output.csv`)
	}

	input := args[0]
	output := args[1]

	log.Printf("Reading from %s and writing the output to %s\n", input, output)
	inputFile, err := os.Open(input)
	defer inputFile.Close()
	if err != nil {
		return err
	}
	outputFile, err := os.Create(output)
	defer outputFile.Close()
	if err != nil {
		return err
	}

	// Write Headers
	fmt.Fprintf(outputFile, "%s%s", BulkOutputHeader(fieldDelimiter), rowDelimiter)
	rowCounts := 0
	scanner := bufio.NewScanner(inputFile)
	for scanner.Scan() {
		byt := scanner.Bytes()
		var dat BulkOutput
		if err := json.Unmarshal(byt, &dat); err != nil {
			log.Printf("[ERROR] %s\n", err.Error())
		}
		// Write Row
		fmt.Fprintf(outputFile, "%s%s", dat.ToRow(multiValueDelimiter, fieldDelimiter), rowDelimiter)
		rowCounts++
		if 0 == rowCounts%1000 {
			log.Printf("Processed %d rows so far\n", rowCounts)
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	log.Printf("Processed %d rows in total\n", rowCounts)

	return nil
}

func init() {
	JSON2Csv.PersistentFlags().StringVar(&fieldDelimiter, "field-delimiter", ",", "Field delimiters for the output file")
	JSON2Csv.PersistentFlags().StringVar(&multiValueDelimiter, "multivalue-delimiter", "|", "Value delimiters for multivalued fields")
}
