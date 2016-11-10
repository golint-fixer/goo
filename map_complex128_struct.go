package goo

var _ Map = MapComplex128Struct(nil)

// MapComplex128Struct is a map from complex128 to struct{}.
type MapComplex128Struct map[complex128]struct{}

// Delete implements Map.
func (m MapComplex128Struct) Delete(k interface{}) {
	delete(m, k.(complex128))
}

// Dereference implements Map.
func (m *MapComplex128Struct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapComplex128Struct) Equals(other Equatable) bool {
	var n = other.(MapComplex128Struct)

	if len(n) != len(m) {
		return false
	}

	for k, v := range m {
		if nv, ok := n[k]; !ok {
			return false
		} else if nv != v {
			return false
		}
	}

	return true
}

// Get implements Map.
func (m MapComplex128Struct) Get(k interface{}) interface{} {
	return m[k.(complex128)]
}

// GetCheck implements Map.
func (m MapComplex128Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(complex128)]

	return v, ok
}

// KeyValues implements Map.
func (m MapComplex128Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapComplex128Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapComplex128Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapComplex128Struct) Make(c int) Map {
	return make(MapComplex128Struct, c)
}

// NotEquals implements Map.
func (m MapComplex128Struct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapComplex128Struct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapComplex128Struct) Set(k, v interface{}) {
	m[k.(complex128)] = v.(struct{})
}
