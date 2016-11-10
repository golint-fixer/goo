package goo

var _ Map = MapIntUint(nil)

// MapIntUint is a map from int to uint.
type MapIntUint map[int]uint

// Delete implements Map.
func (m MapIntUint) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntUint) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntUint) Equals(other Equatable) bool {
	var n = other.(MapIntUint)

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
func (m MapIntUint) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntUint) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntUint) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntUint) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntUint) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntUint) Make(c int) Map {
	return make(MapIntUint, c)
}

// NotEquals implements Map.
func (m MapIntUint) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntUint) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntUint) Set(k, v interface{}) {
	m[k.(int)] = v.(uint)
}
