# :mechanical_arm: Excel to SQL Insert Generator (Written in Go and Bash)

## :robot: Overview

The Excel to SQL Insert Generator is a program written in Go and Bash that converts data from an Excel spreadsheet into SQL insert statements.
It automates the process of generating SQL insert statements, making it easier to populate a database from Excel data.
The program includes a Bash script that simplifies execution and provides a convenient user interface, prompting users to enter necessary data.

## :pill: Real-Life Use Case

Many companies face the challenge of migrating data from one system to another or integrating data from external sources into their existing systems.
Excel files are commonly used to store and organize data, making them a convenient format for data transfer.
The Excel to SQL Insert Generator streamlines this process by generating SQL insert statements directly from Excel files.
This allows for seamless migration or integration of data into databases, saving time and effort for developers and data analysts.

## :sparkles: Features

:old_key: Converts Excel data into SQL insert statements using the Go program. <br>
:old_key: Simplifies usage with a Bash (sh) script that prompts users for necessary data and executes the Go program. <br>
:old_key: Supports custom table names and file paths provided by the user. <br>
:old_key: Automatically generates a unique SQL file for the result, which includes the table name and timestamp of the execution. <br>
:old_key: The SQL file is saved in the "files" directory of the project for easy access. <br>

## :hammer_and_wrench: Installation and Usage

To use the Excel to SQL Insert Generator, follow these steps:

1. Clone the repository:
```
git clone https://github.com/ArtiomStartev/Excel-to-SQL-Insert-Generator.git
```
2. Navigate to the project directory:
```
cd Excel-to-SQL-Insert-Generator
```
3. Open the "scripts" folder:
```
cd scripts
```
4. Run the Bash script generate_inserts.sh using /bin/bash:
```
/bin/bash generate_inserts.sh
```
5. The script will prompt you to enter the table name and the path to the Excel file.
6. After providing the necessary data, the script will execute the Go program automatically.
7. The Go program will generate a unique SQL file in the "files" directory of the project.
