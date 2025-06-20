package main

import (
	"testing"
)

func TestProcessSpec(t *testing.T) {
	spec := map[string]map[string]interface{}{
		"id": {
			"type": "uuid",
		},
		"row": {
			"type": "rownum",
		},
		"num": {
			"type": "integer",
		},
	}

	// Call the function with a single iteration for regression testing
	ProcessSpec(spec, 1)
}
