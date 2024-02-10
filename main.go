package main

import (
	"fmt"
	"log"
	"os"
	"strings"
	"time"

	"github.com/tealeg/xlsx"
)

// Environment variables to be passed as command line arguments
var (
	TableName = "TABLE_NAME="
	FilePath  = "FILE_PATH="
)

func main() {
	// Extract environment variables from command line arguments
	tableName := strings.TrimPrefix(os.Args[1], TableName)
	filePath := strings.TrimPrefix(os.Args[2], FilePath)

	// Generate a unique filename for the result file
	resultFileName := generateResultFileName(tableName)

	// Create a new file to write the insert statements
	resultFile, err := os.Create("../files/" + resultFileName)
	if err != nil {
		log.Fatalf("Error creating result file: %s", err)
	}
	defer resultFile.Close()

	// Create a buffer to store insert statements
	var insertStatements strings.Builder

	// Open the Excel file
	xlFile, err := xlsx.OpenFile(filePath)
	if err != nil {
		log.Fatalf("Error opening Excel file: %s", err)
	}

	// Extract the first sheet from the Excel file
	sheet := xlFile.Sheets[0]

	// Extract headers from the first row
	headers := extractRecord(sheet.Rows[0])

	// Iterate through rows and generate INSERT statements
	for _, row := range sheet.Rows[1:] {
		record := extractRecord(row)
		insertStatement := generateInsertStatement(tableName, headers, record)

		// Write insert statement to the buffer
		insertStatements.WriteString(insertStatement + "\n")
	}

	// Write the insert statements from the buffer to the result file
	if _, err = resultFile.WriteString(insertStatements.String()); err != nil {
		log.Fatalf("Error writing to result file: %s", err)
	}
}

func generateResultFileName(tableName string) string {
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	return fmt.Sprintf("result_inserts_%s_%s.sql", tableName, timestamp)
}

func extractRecord(row *xlsx.Row) []string {
	var record []string
	for _, cell := range row.Cells {
		record = append(record, cell.String())
	}

	return record
}

func generateInsertStatement(table string, headers, record []string) string {
	columns := strings.Join(headers, ", ")
	values := strings.Join(record, ", ")

	return fmt.Sprintf("INSERT INTO %s (%s) VALUES (%s);", table, columns, values)
}
