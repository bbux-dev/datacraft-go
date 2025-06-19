package loader

import (
	"testing"

	"github.com/bbux-dev/datacraft-go/internal/factories"
	"github.com/bbux-dev/datacraft-go/pkg/registry"
)

func TestLoader_Get(t *testing.T) {
	// Create a test spec with both uuid and rownum fields
	spec := map[string]map[string]interface{}{
		"id": {
			"type": "uuid",
		},
		"row": {
			"type": "rownum",
		},
	}

	// Initialize registry and register built-in factories
	reg := registry.NewRegistry()
	factories.Register(reg)

	// Create loader with our test spec
	loader := NewLoader(reg, spec)

	// Test getting UUID supplier
	uuidSupplier, err := loader.Get("id")
	if err != nil {
		t.Errorf("Get() error for uuid = %v", err)
	}
	if uuidSupplier == nil {
		t.Error("Get() returned nil supplier for uuid")
	}

	// Test getting RowNumber supplier
	rowSupplier, err := loader.Get("row")
	if err != nil {
		t.Errorf("Get() error for row = %v", err)
	}
	if rowSupplier == nil {
		t.Error("Get() returned nil supplier for row")
	}

	// Test that suppliers are cached
	cachedUuidSupplier, err := loader.Get("id")
	if err != nil {
		t.Errorf("Get() error for cached uuid = %v", err)
	}
	if cachedUuidSupplier != uuidSupplier {
		t.Error("Get() did not return cached supplier for uuid")
	}

	// Test getting non-existent supplier type
	_, err = loader.Get("non-existent")
	if err == nil {
		t.Error("Get() expected error for non-existent supplier type")
	}
}

func TestLoader_GetWithInvalidSpec(t *testing.T) {
	// Create a test spec with an invalid field type
	spec := map[string]map[string]interface{}{
		"invalid": {
			"type": "non-existent",
		},
	}

	// Initialize registry and register built-in factories
	reg := registry.NewRegistry()
	factories.Register(reg)

	// Create loader with our test spec
	loader := NewLoader(reg, spec)

	// Test getting supplier for invalid type
	_, err := loader.Get("non-existent")
	if err == nil {
		t.Error("Get() expected error for invalid supplier type")
	}
}
