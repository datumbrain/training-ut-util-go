package main

import (
	"bufio"
	"encoding/csv"
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	var operation string
	var inputPath string
	var outputPath string
	var format string

	flag.StringVar(&operation, "operation", "", "Specify the operation")
	flag.StringVar(&inputPath, "input-path", "", "Specify the Input Path")
	flag.StringVar(&outputPath, "output-path", "", "Specify the Output Path")
	flag.StringVar(&format, "format", "", "Specify the format")
	flag.Parse()

	if operation != "merge" {
		fmt.Println("Only merge operation is allowed")
		return
	}

	if format != "" && format != "csv" {
		fmt.Println("Only CSV format is allowed")
		return
	}

	files, err := filepath.Glob(filepath.Join(inputPath, "*.jsonl"))
	if err != nil {
		fmt.Printf("Error finding .jsonl files: %v\n", err)
		return
	}

	if format == "csv" {
		csvFilename := filepath.Join(outputPath, "output.csv")
		if err := convertToCSV(files, csvFilename); err != nil {
			fmt.Printf("Error converting to CSV: %v\n", err)
			return
		}
		fmt.Printf("CSV data saved to: %s\n", csvFilename)
	} else {
		outputFilename := filepath.Join(outputPath, "output.jsonl")
		outputFile, err := os.OpenFile(outputFilename, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
		if err != nil {
			fmt.Printf("Error opening output file: %v\n", err)
			return
		}
		defer outputFile.Close()
		writer := bufio.NewWriter(outputFile)

		for _, file := range files {
			fmt.Printf("Merging file: %v\n", file)

			f, err := os.Open(file)
			if err != nil {
				fmt.Printf("Error Opening File %s: %v\n", file, err)
				continue
			}
			defer f.Close()

			scanner := bufio.NewScanner(f)
			for scanner.Scan() {
				line := scanner.Text()
				if len(line) == 0 {
					continue // Skip empty lines
				}
				_, err := writer.WriteString(line + "\n")
				if err != nil {
					fmt.Printf("Error writing to output file: %v\n", err)
					return
				}
			}
			if err := scanner.Err(); err != nil {
				fmt.Printf("Error reading file %s: %v\n", file, err)
			}
		}

		writer.Flush()
		fmt.Printf("Merged data saved to: %s\n", outputFilename)
	}
}

func convertToCSV(jsonlFiles []string, csvFilename string) error {
	records := []map[string]interface{}{} //A slice of maps to store each JSON object read from the JSONL files
	fieldSet := map[string]bool{}         //A map to keep track of all unique fields(headers)

	// Reading and Parsin JSONL file
	for _, file := range jsonlFiles {
		jsonlFile, err := os.Open(file)
		if err != nil {
			return fmt.Errorf("error opening JSONL file %s: %v", file, err)
		}
		defer jsonlFile.Close()

		scanner := bufio.NewScanner(jsonlFile)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) == 0 {
				continue // Skip empty lines
			}
			record := map[string]interface{}{}
			err := json.Unmarshal([]byte(line), &record)
			if err != nil {
				return fmt.Errorf("error parsing JSON in file %s: %v", file, err)
			}
			records = append(records, record)
			for key := range record {
				fieldSet[key] = true
			}
		}
		if err := scanner.Err(); err != nil {
			return fmt.Errorf("error reading file %s: %v", file, err)
		}
	}

	fields := []string{}
	for field := range fieldSet {
		fields = append(fields, field)
	}

	csvFile, err := os.Create(csvFilename)
	if err != nil {
		return fmt.Errorf("error creating CSV file: %v", err)
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	if err := writer.Write(fields); err != nil {
		return fmt.Errorf("error writing header to CSV: %v", err)
	}

	// Writing records to the CSV file

	for _, record := range records {
		row := make([]string, len(fields))
		for i, field := range fields {
			if val, ok := record[field]; ok {
				row[i] = fmt.Sprintf("%v", val)
			} else {
				row[i] = ""
			}
		}
		if err := writer.Write(row); err != nil {
			return fmt.Errorf("error writing record to CSV: %v", err)
		}
	}

	return nil
}
