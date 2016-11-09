package goo

var _ Map = MapStringComplex128(nil)

// MapStringComplex128 is a map from string to complex128.
type MapStringComplex128 map[string]complex128

// Delete implements Map.
func (m MapStringComplex128) Delete(k interface{}) {
	delete(m, k.(string))
}

// Equals implements Map.
func (m MapStringComplex128) Equals(other Equatable) bool {
	var n = other.(MapStringComplex128)

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
func (m MapStringComplex128) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringComplex128) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringComplex128) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringComplex128) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringComplex128) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringComplex128) Make(c int) Map {
	return make(MapStringComplex128, c)
}

// NotEquals implements Map.
func (m MapStringComplex128) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapStringComplex128) Set(k, v interface{}) {
	m[k.(string)] = v.(complex128)
}
