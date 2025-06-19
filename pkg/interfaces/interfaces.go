package interfaces

// ValueSupplier defines an interface for suppliers that can provide values.
// It has a Next method that returns the next value for a given iteration.
type ValueSupplier interface {
	Next(iteration int) (any, error)
}

// Loader is an interface for loading value suppliers based on a key name.
type Loader interface {
	Get(spec string) (ValueSupplier, error)
}

// SupplierFactory is a function that instantiates a ValueSupplier.
// It takes a map of parameters and a Loader, returning a ValueSupplier or an error.
type SupplierFactory interface {
	Create(map[string]interface{}, Loader) (ValueSupplier, error)
}

// Registry is an interface for managing suppliers.
// It allows registering suppliers by name and retrieving them by name.
type Registry interface {
	RegisterSupplier(name string, supplier SupplierFactory)
	GetSupplierFactory(name string) (SupplierFactory, error)
}
