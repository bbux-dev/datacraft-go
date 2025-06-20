package suppliers

import (
	"math/rand/v2"
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

func TestIntegerSupplier_Next(t *testing.T) {
	supplier := &IntegerSupplier{
		Min:    1,
		Max:    10,
		Random: rand.New(rand.NewPCG(0, 0)), // Fixed seed for reproducibility
	}

	// Test first integer generation
	int1, err := supplier.Next(0)
	if err != nil {
		t.Errorf("IntegerSupplier.Next() error = %v", err)
	}
	if int1.(int64) < 1 || int1.(int64) > 10 {
		t.Errorf("IntegerSupplier.Next() = %v, want in range [1, 10]", int1)
	}

	// Test second integer generation
	int2, err := supplier.Next(1)
	if err != nil {
		t.Errorf("IntegerSupplier.Next() error = %v", err)
	}
	if int2.(int64) < 1 || int2.(int64) > 10 {
		t.Errorf("IntegerSupplier.Next() = %v, want in range [1, 10]", int2)
	}
}
