package goo

var _ Map = MapComplex128String(nil)

// MapComplex128String is a map from complex128 to string.
type MapComplex128String map[complex128]string

// Delete implements Map.
func (m MapComplex128String) Delete(k interface{}) {
	delete(m, k.(complex128))
}

// Equals implements Map.
func (m MapComplex128String) Equals(other Equatable) bool {
	var n = other.(MapComplex128String)

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
func (m MapComplex128String) Get(k interface{}) interface{} {
	return m[k.(complex128)]
}

// GetCheck implements Map.
func (m MapComplex128String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(complex128)]

	return v, ok
}

// KeyValues implements Map.
func (m MapComplex128String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapComplex128String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapComplex128String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapComplex128String) Make(c int) Map {
	return make(MapComplex128String, c)
}

// NotEquals implements Map.
func (m MapComplex128String) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapComplex128String) Set(k, v interface{}) {
	m[k.(complex128)] = v.(string)
}
