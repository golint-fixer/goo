package goo

// Value can be referenced.
type Value interface {
	// Reference returns a reference.
	Reference() Pointer
}
