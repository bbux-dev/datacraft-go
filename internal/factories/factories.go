package factories

import (
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

func Register(registry interfaces.Registry) {
	registry.RegisterSupplier("uuid", &UuidSupplierFactory{})
	registry.RegisterSupplier("rownum", &RowNumberSupplierFactory{})
}
