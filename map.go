package goo

// Map is a map.
type Map interface {
	// Delete deletes k.
	Delete(k interface{})

	// Get returns the value for key k.
	Get(k interface{}) interface{}

	// GetCheck returns the value for key k and whether key k exists.
	GetCheck(k interface{}) (interface{}, bool)

	// KeyValues returns key-value pairs.
	KeyValues() [][2]interface{}

	// Keys returns the keys.
	Keys() []interface{}

	// Len returns the length.
	Len() int

	// Make returns a Map with capacity c.
	Make(c int) Map

	// Set sets the element for key k to v.
	Set(k, v interface{})
}
