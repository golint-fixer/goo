package goo

var _ Map = MapComplex64Struct(nil)

// MapComplex64Struct is a map from complex64 to struct{}.
type MapComplex64Struct map[complex64]struct{}

// Delete implements Map.
func (m MapComplex64Struct) Delete(k interface{}) {
	delete(m, k.(complex64))
}

// Get implements Map.
func (m MapComplex64Struct) Get(k interface{}) interface{} {
	return m[k.(complex64)]
}

// GetCheck implements Map.
func (m MapComplex64Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(complex64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapComplex64Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapComplex64Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapComplex64Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapComplex64Struct) Make(c int) Map {
	return make(MapComplex64Struct, c)
}

// Set implements Map.
func (m MapComplex64Struct) Set(k, v interface{}) {
	m[k.(complex64)] = v.(struct{})
}
