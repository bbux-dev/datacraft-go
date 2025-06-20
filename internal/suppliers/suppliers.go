package suppliers

import (
	"math/rand/v2"

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

// IntegerSupplier implements ValueSupplier[int64]
// It generates random integers within a specified range.
// The range is defined by min and max values.
type IntegerSupplier struct {
	Min    int64
	Max    int64
	Random *rand.Rand
}

func (s *IntegerSupplier) Next(iteration int) (any, error) {
	return s.Random.Int64N(s.Max-s.Min+1) + s.Min, nil
}
