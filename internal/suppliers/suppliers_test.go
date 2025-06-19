package suppliers

import (
	"testing"
)

func TestUuidSupplier_Next(t *testing.T) {
	supplier := &UuidSupplier{}

	// Test first UUID generation
	uuid1, err := supplier.Next(0)
	if err != nil {
		t.Errorf("UuidSupplier.Next() error = %v", err)
	}
	if uuid1 == "" {
		t.Error("UuidSupplier.Next() returned empty string")
	}

	// Test second UUID generation to ensure it's different
	uuid2, err := supplier.Next(1)
	if err != nil {
		t.Errorf("UuidSupplier.Next() error = %v", err)
	}
	if uuid1 == uuid2 {
		t.Error("UuidSupplier.Next() returned same UUID for different iterations")
	}
}

func TestRowNumberSupplier_Next(t *testing.T) {
	supplier := &RowNumberSupplier{}

	// Test first row number
	row1, err := supplier.Next(0)
	if err != nil {
		t.Errorf("RowNumberSupplier.Next() error = %v", err)
	}
	if row1 != 0 {
		t.Errorf("RowNumberSupplier.Next() = %v, want %v", row1, 0)
	}

	// Test second row number
	row2, err := supplier.Next(1)
	if err != nil {
		t.Errorf("RowNumberSupplier.Next() error = %v", err)
	}
	if row2 != 1 {
		t.Errorf("RowNumberSupplier.Next() = %v, want %v", row2, 1)
	}
}
