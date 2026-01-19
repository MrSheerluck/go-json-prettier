package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
)

func main() {
	writeFile := flag.Bool("w", false, "write formatted JSON back to file")
	flag.Parse()
	files := flag.Args()
	for i := range files {
		content, err := os.ReadFile(files[i])
		if err != nil {
			os.Exit(1)
		}
		var data interface{}
		err = json.Unmarshal(content, &data)
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s", err)
		}
		prettyJSON, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
				fmt.Fprintf(os.Stderr, "Error formatting: %v\n", err)
				continue
		}
		if *writeFile {
			err = os.WriteFile(files[i], prettyJSON, 0644)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Error writing file: %v\n", err)
				}
		} else {
				fmt.Printf("%s\n", prettyJSON)
		}
	}
}