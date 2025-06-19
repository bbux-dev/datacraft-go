package suppliers

import (
	"github.com/google/uuid"
)

// UuidSupplier implements ValueSupplier[string]
type UuidSupplier struct{}

func (s *UuidSupplier) Next(iteration int) (any, error) {
	return uuid.New().String(), nil
}

// RowNumberSupplier implements ValueSupplier[int]
type RowNumberSupplier struct{}

func (s *RowNumberSupplier) Next(iteration int) (any, error) {
	return iteration, nil
}
