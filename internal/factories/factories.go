package factories

import (
	"math/rand/v2"

	. "github.com/bbux-dev/datacraft-go/internal/suppliers"
	"github.com/bbux-dev/datacraft-go/pkg/interfaces"
)

type UuidSupplierFactory struct{}

func (f *UuidSupplierFactory) Create(spec map[string]interface{}, loader interfaces.Loader) (interfaces.ValueSupplier, error) {
	return &UuidSupplier{}, nil
}

type RowNumberSupplierFactory struct{}

func (f *RowNumberSupplierFactory) Create(spec map[string]interface{}, loader interfaces.Loader) (interfaces.ValueSupplier, error) {
	return &RowNumberSupplier{}, nil
}

type IntegerSupplierFactory struct{}

func (f *IntegerSupplierFactory) Create(spec map[string]interface{}, loader interfaces.Loader) (interfaces.ValueSupplier, error) {
	const MaxInt = int64(^uint64(0) >> 1)
	const MinInt = -MaxInt - 1

	min, ok := spec["min"].(int64)
	if !ok {
		min = MinInt // Default value if not specified
	}
	max, ok := spec["max"].(int64)
	if !ok {
		max = MaxInt // Default value if not specified
	}

	seed, ok := spec["seed"].(uint64)
	if !ok {
		seed = uint64(1234567891011121314) // Default seed if not specified
	}
	gen := rand.New(rand.NewPCG(42, seed))

	return &IntegerSupplier{Min: min, Max: max, Random: gen}, nil
}

func Register(registry interfaces.Registry) {
	registry.RegisterSupplier("uuid", &UuidSupplierFactory{})
	registry.RegisterSupplier("rownum", &RowNumberSupplierFactory{})
	registry.RegisterSupplier("iteration", &RowNumberSupplierFactory{})
	registry.RegisterSupplier("integer", &IntegerSupplierFactory{})
}
