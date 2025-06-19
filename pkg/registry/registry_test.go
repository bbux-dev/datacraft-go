package registry

import (
	"testing"

	"github.com/bbux-dev/datacraft-go/internal/suppliers"
	"github.com/bbux-dev/datacraft-go/pkg/interfaces"
)

// Mock supplier factory for testing
type mockSupplierFactory struct{}

func (m *mockSupplierFactory) Create(spec map[string]interface{}, loader interfaces.Loader) (interfaces.ValueSupplier, error) {
	return &suppliers.RowNumberSupplier{}, nil
}

func TestRegistry_RegisterAndGetSupplier(t *testing.T) {
	reg := NewRegistry()
	mockFactory := &mockSupplierFactory{}

	// Test registering a supplier
	reg.RegisterSupplier("test-supplier", mockFactory)

	// Test getting the registered supplier
	factory, err := reg.GetSupplierFactory("test-supplier")
	if err != nil {
		t.Errorf("GetSupplierFactory() error = %v", err)
	}
	if factory == nil {
		t.Error("GetSupplierFactory() returned nil factory")
	}
}

func TestRegistry_GetNonExistentSupplier(t *testing.T) {
	reg := NewRegistry()

	// Test getting a non-existent supplier
	factory, err := reg.GetSupplierFactory("non-existent")
	if err == nil {
		t.Error("GetSupplierFactory() expected error for non-existent supplier")
	}
	if factory != nil {
		t.Error("GetSupplierFactory() returned non-nil factory for non-existent supplier")
	}
}

func TestRegistry_RegisterOverwrite(t *testing.T) {
	reg := NewRegistry()
	mockFactory1 := &mockSupplierFactory{}
	mockFactory2 := &mockSupplierFactory{}

	// Register first factory
	reg.RegisterSupplier("test-supplier", mockFactory1)

	// Register second factory with same name
	reg.RegisterSupplier("test-supplier", mockFactory2)

	// Get the factory and verify it's the second one
	factory, err := reg.GetSupplierFactory("test-supplier")
	if err != nil {
		t.Errorf("GetSupplierFactory() error = %v", err)
	}
	if factory != mockFactory2 {
		t.Error("GetSupplierFactory() did not return the most recently registered factory")
	}
}
