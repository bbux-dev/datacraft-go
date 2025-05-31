package main

import (
	"flag"
	"fmt"
	"os"
	"encoding/json"
)


func main() {
	specFile := flag.String("spec", "", "Path to data spec file")
	iterations := flag.Int("iterations", 3, "Number of records to generate")

	// Add short versions of flags
	flag.StringVar(specFile, "s", "", "Path to the specification file (JSON/YAML) (shorthand)")
	flag.IntVar(iterations, "i", 1, "Number of records to generate (shorthand)")
	
	// Parse command line flags
	flag.Parse()

	if *specFile == "" {
		fmt.Println("Error: --spec (-s) flag is required")
		flag.Usage()
		os.Exit(1)
	}

	if *iterations < 1 {
		fmt.Println("Error: --iterations (-i) must be greater than 0")
		flag.Usage()
		os.Exit(1)
	}

	// Load data spec
	spec, err := loadSpec(*specFile)
	if err != nil {
		fmt.Printf("Error loading spec file: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Loaded spec: %v\n", spec)
	
}

func loadSpec(path string) (*map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %v", err)
	}

	var spec map[string]interface{}
	if err := json.Unmarshal(content, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse spec file: %v", err)
	}

	return &spec, nil
}
