package loader

import (
	"fmt"

	. "github.com/bbux-dev/datacraft-go/pkg/interfaces"
)

type loader struct {
	registry Registry
	spec     map[string]map[string]interface{}
	cache    map[string]ValueSupplier
}

func (l *loader) Get(spec string) (ValueSupplier, error) {
	if l.cache[spec] != nil {
		return l.cache[spec], nil
	}

	field_spec, ok := l.spec[spec]
	if !ok {
		return nil, fmt.Errorf("spec %s not found", spec)
	}
	field_type, ok := field_spec["type"].(string)
	if !ok {
		return nil, fmt.Errorf("spec %s does not have a valid type", spec)
	}

	factory, err := l.registry.GetSupplierFactory(field_type)
	if err != nil {
		return nil, err
	}
	supplier, err := factory.Create(field_spec, l)
	if err != nil {
		return nil, err
	}
	l.cache[spec] = supplier
	return supplier, nil
}

func NewLoader(registry Registry, spec map[string]map[string]interface{}) Loader {
	return &loader{
		registry: registry,
		spec:     spec,
		cache:    make(map[string]ValueSupplier),
	}
}
