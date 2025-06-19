package registry

import (
	"fmt"

	. "github.com/bbux-dev/datacraft-go/pkg/interfaces"
)

type registry struct {
	suppliers map[string]SupplierFactory
}

func (r *registry) RegisterSupplier(name string, supplier SupplierFactory) {
	r.suppliers[name] = supplier
}

func (r *registry) GetSupplierFactory(name string) (SupplierFactory, error) {
	supplier, exists := r.suppliers[name]
	if !exists {
		return nil, fmt.Errorf("supplier factory '%s' not found", name)
	}
	return supplier, nil
}

func NewRegistry() Registry {
	return &registry{
		suppliers: make(map[string]SupplierFactory),
	}
}
