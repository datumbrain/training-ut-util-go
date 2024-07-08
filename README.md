# Utility

This repository contains a Go CLI that performs file merging operations. The CLI merges all JSONL files in a specified directory into a single JSONL file and stores it in an output directory. Additionally, if the --format csv flag is used, the output is a CSV file.

## Task Description

**Objective:** Create a Go CLI tool named `util` that performs file merging operations. The tool should merge all JSONL files in a specified directory into a single JSONL file and store it in an output directory. Additionally, if the `--format csv` flag is used, the output should be a CSV file.

## How to run the Tool
   - cd scripts
   - chmod +x build.sh
   - ./build.sh --operation merge --input-path [your input-path] --output-path [your output-path]
      Then follow the terminal guidelines 

## Requirements

1. **CLI Flags:**

   - `--operation merge`: Specifies the operation to be performed.
   - `--input-path ./<any_path>`: Path to the directory containing JSONL files to be merged.
   - `--output-path ./<any_path>`: Path to the directory where the merged file will be stored.
   - `--format csv`: (Optional) If provided, the output should be a CSV file instead of a JSONL file.

2. **Functionality:**

   - Read all JSONL files from the input directory.
   - Merge the content of all JSONL files into a single JSONL file.
   - If the `--format csv` flag is detected, convert the merged JSONL content to CSV format.
   - Save the merged file (in JSONL or CSV format) to the specified output directory.

3. **Bonus:** Implement the `--format csv` flag to output a CSV file.

## Steps to Complete the Task

1. **Setup Go Environment:**

   - Ensure Go is installed on your machine.
   - Create a new Go module.

2. **Implement the CLI Tool:**

   - Use a CLI package like `flag` or `cobra` for parsing command-line arguments.
   - Implement functions to read and merge JSONL files.
   - Implement functions to convert JSONL to CSV format (if required).
   - Implement the main function to handle CLI arguments and execute the appropriate functions.

3. **Testing:**

   - Create a few sample JSONL files for testing.
   - Test the tool by running the command with different flag combinations.

4. **Documentation:**
   - Provide clear instructions on how to use the tool, including examples.

## Example Commands

1. Merge JSONL files into a single JSONL file:

   ```bash
   ./bin/util --operation merge --input-path ./data/jsonl_files/ --output-path ./output_data
   ```

2. Merge JSONL files and output as a CSV file:

   ```bash
   ./bin/util --operation merge --input-path ./data/jsonl_files/ --output-path ./output_data --format csv
   ```
## Submission

Submit the following:

1. The complete Go code.
2. Sample input JSONL files for testing.
3. A README file with instructions on how to build and run the tool, including example commands.

Good luck with the task!
