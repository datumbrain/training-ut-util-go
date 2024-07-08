#!/bin/bash

usage(){
    echo -e "Press 1 to merge the JSONL files into a separate file\n"
    echo -e "Press 2 to merge the JSONL files into a CSV\n"
}
docker-compose up --build
operation=""
input_path=""
output_path=""
format="jsonl"

# Parse command-line arguments
while [[ $# -gt 0 ]]; do
    case "$1" in
        --operation)
            operation="$2"
            shift 2  # Move past the argument and its value
            ;;
        --input-path)
            input_path="$2"
            shift 2
            ;;
        --output-path)
            output_path="$2"
            shift 2
            ;;
        --format)
            format="$2"
            shift 2
            ;;
        *)
            echo "Invalid option: $1"
            ;;
    esac
done

if [ "$1" == "-b" ]; then
    cd ..
    mkdir -p bin
    go build -o bin/util ./util.go
else
    usage
    read -r input
    if [ "$input" == 1 ]; then
        cd ..
        go run util.go --operation merge --input-path "$input_path" --output-path "$output_path"
    elif [ "$input" == 2 ]; then
        cd ..
        go run util.go --operation merge --input-path "$input_path" --output-path "$output_path" --format csv
    fi
fi
