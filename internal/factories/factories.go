package factories

import (
	"encoding/json"
	"fmt"
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

func (f *IntegerSupplierFactory) Create(
	spec map[string]any,
	loader interfaces.Loader,
) (interfaces.ValueSupplier, error) {

	// ---------------------------------------------------------------------
	// 1) Pull min / max from the spec ­– accept int, float64, json.Number…
	// ---------------------------------------------------------------------
	min := int64(-1_000_000_000) // default = -1 billion
	if v, ok := spec["min"]; ok {
		m, err := asInt64(v)
		if err != nil {
			return nil, fmt.Errorf("min: %w", err)
		}
		min = m
	}

	max := int64(1_000_000_000) // default = 1 billion
	if v, ok := spec["max"]; ok {
		m, err := asInt64(v)
		if err != nil {
			return nil, fmt.Errorf("max: %w", err)
		}
		max = m
	}

	if min >= max {
		return nil, fmt.Errorf("min (%d) must be < max (%d)", min, max)
	}

	// ---------------------------------------------------------------------
	// 2) Optional seed value
	// ---------------------------------------------------------------------
	seed := uint64(rand.Uint64()) // default = nondeterministic
	if v, ok := spec["seed"]; ok {
		s, err := asUint64(v)
		if err != nil {
			return nil, fmt.Errorf("seed: %w", err)
		}
		seed = s
	}

	// PCG needs two 64-bit seeds; a different second seed gives a
	// different *stream* without changing the initial state.
	prng := rand.New(rand.NewPCG(seed, seed^0xdeadbeefcafebabe))

	return &IntegerSupplier{
		Min:    min,
		Max:    max,
		Random: prng,
	}, nil
}

// --------------------------------- helper utilities ----------------------

func asInt64(v any) (int64, error) {
	switch n := v.(type) {
	case int64:
		return n, nil
	case int:
		return int64(n), nil
	case float64:
		return int64(n), nil
	case json.Number: // if spec came from encoding/json with UseNumber
		return n.Int64()
	default:
		return 0, fmt.Errorf("unsupported numeric type %T", v)
	}
}

func asUint64(v any) (uint64, error) {
	switch n := v.(type) {
	case uint64:
		return n, nil
	case uint:
		return uint64(n), nil
	case int64:
		if n < 0 {
			return 0, fmt.Errorf("negative value not allowed")
		}
		return uint64(n), nil
	case float64:
		if n < 0 {
			return 0, fmt.Errorf("negative value not allowed")
		}
		return uint64(n), nil
	case json.Number:
		i, err := n.Int64()
		if err != nil {
			return 0, err
		}
		if i < 0 {
			return 0, fmt.Errorf("negative value not allowed")
		}
		return uint64(i), nil
	default:
		return 0, fmt.Errorf("unsupported numeric type %T", v)
	}
}

func Register(registry interfaces.Registry) {
	registry.RegisterSupplier("uuid", &UuidSupplierFactory{})
	registry.RegisterSupplier("rownum", &RowNumberSupplierFactory{})
	registry.RegisterSupplier("iteration", &RowNumberSupplierFactory{})
	registry.RegisterSupplier("integer", &IntegerSupplierFactory{})
}
