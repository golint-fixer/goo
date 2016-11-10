package goo

var _ Map = MapIntComplex128(nil)

// MapIntComplex128 is a map from int to complex128.
type MapIntComplex128 map[int]complex128

// Delete implements Map.
func (m MapIntComplex128) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntComplex128) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntComplex128) Equals(other Equatable) bool {
	var n = other.(MapIntComplex128)

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
func (m MapIntComplex128) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntComplex128) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntComplex128) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntComplex128) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntComplex128) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntComplex128) Make(c int) Map {
	return make(MapIntComplex128, c)
}

// NotEquals implements Map.
func (m MapIntComplex128) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntComplex128) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntComplex128) Set(k, v interface{}) {
	m[k.(int)] = v.(complex128)
}
