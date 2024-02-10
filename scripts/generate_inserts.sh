#!/bin/bash

# Define colors for output messages
GREEN="$(tput setaf 2)"
RED="$(tput setaf 1)"
NC="$(tput sgr0)" # No Color

# Prompt the user for the table name
read -p "Please enter the table name: " tableName

# Prompt the user for the Excel file name
read -p "Please enter the Excel file name (file should be in the 'files' directory of the project): " fileName

# Run the Go program with the provided table name and file path as arguments
if go run "../main.go" TABLE_NAME="$tableName" FILE_PATH="../files/$fileName"; then
    # If the Go program executed successfully
    echo -e "${GREEN}Insert statements for '$tableName' table were generated successfully.${NC}"
    echo -e "${GREEN}The result file was created in the 'files' directory.${NC}"
    say "Go program execution was finished successfully."
else
    # If there was an error executing the Go program
    echo -e "${RED}Failed to generate insert statements.${NC}"
    say "Go program execution failed."
fi
