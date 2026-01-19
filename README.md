# go-json-prettier

A lightweight CLI tool to format and prettify JSON files.

## Features

- Pretty-print JSON with proper indentation
- Process multiple files at once
- Write formatted output back to files in-place with the `-w` flag
- Error handling for malformed JSON

## Installation

```bash
go build -o json-prettier
```

## Usage

**Print formatted JSON to stdout:**

```bash
./json-prettier file.json
```

**Format multiple files:**

```bash
./json-prettier file1.json file2.json file3.json
```

**Write formatted JSON back to file (in-place):**

```bash
./json-prettier -w file.json
```

## Examples

```bash
# View prettified JSON
./json-prettier data.json

# Format multiple files and save them
./json-prettier -w config1.json config2.json
```

## How it Works

1. Reads the JSON file
2. Parses the JSON content
3. Re-formats with 2-space indentation
4. Either prints to stdout or writes back to the file (with `-w` flag)

## Error Handling

- Reports file not found errors
- Displays JSON parsing errors with context
- Handles file write permission issues
