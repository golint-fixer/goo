package goo

var _ Map = MapComplex64String(nil)

// MapComplex64String is a map from complex64 to string.
type MapComplex64String map[complex64]string

// Delete implements Map.
func (m MapComplex64String) Delete(k interface{}) {
	delete(m, k.(complex64))
}

// Get implements Map.
func (m MapComplex64String) Get(k interface{}) interface{} {
	return m[k.(complex64)]
}

// GetCheck implements Map.
func (m MapComplex64String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(complex64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapComplex64String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapComplex64String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapComplex64String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapComplex64String) Make(c int) Map {
	return make(MapComplex64String, c)
}

// Set implements Map.
func (m MapComplex64String) Set(k, v interface{}) {
	m[k.(complex64)] = v.(string)
}
