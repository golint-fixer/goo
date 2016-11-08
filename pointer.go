package goo

// Pointer references a value.
type Pointer interface {
	// Dereference returns the referenced value.
	Dereference() Value
}
