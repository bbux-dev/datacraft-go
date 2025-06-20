package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/bbux-dev/datacraft-go/internal/factories"
	"github.com/bbux-dev/datacraft-go/pkg/loader"
	"github.com/bbux-dev/datacraft-go/pkg/registry"
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
	ProcessSpec(spec, *iterations)
}

func ProcessSpec(spec map[string]map[string]interface{}, iterations int) {

	// Initialize registry and loader
	registry := registry.NewRegistry()
	factories.Register(registry) // Register built-in factories
	loader := loader.NewLoader(registry, spec)

	for i := range iterations {
		for field, fieldSpec := range spec {
			fmt.Printf("Generating value for field '%s':\n", field)
			fmt.Printf("Spec: %+v\n", fieldSpec)
			supplier, err := loader.Get(field)
			if err != nil {
				fmt.Printf("Error getting supplier for field '%s': %v\n", field, err)
				continue
			}
			value, err := supplier.Next(i)
			if err != nil {
				fmt.Printf("Error generating value for field '%s': %v\n", field, err)
				continue
			}
			fmt.Printf("Generated value for field '%s': %v\n", field, value)
			fmt.Println()
		}
	}

}

func loadSpec(path string) (map[string]map[string]interface{}, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("failed to read spec file: %v", err)
	}

	var spec map[string]map[string]interface{}
	if err := json.Unmarshal(content, &spec); err != nil {
		return nil, fmt.Errorf("failed to parse spec file: %v", err)
	}

	return spec, nil
}
