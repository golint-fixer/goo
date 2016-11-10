package goo

var _ Map = MapIntUint8(nil)

// MapIntUint8 is a map from int to uint8.
type MapIntUint8 map[int]uint8

// Delete implements Map.
func (m MapIntUint8) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntUint8) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntUint8) Equals(other Equatable) bool {
	var n = other.(MapIntUint8)

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
func (m MapIntUint8) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntUint8) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntUint8) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntUint8) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntUint8) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntUint8) Make(c int) Map {
	return make(MapIntUint8, c)
}

// NotEquals implements Map.
func (m MapIntUint8) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntUint8) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntUint8) Set(k, v interface{}) {
	m[k.(int)] = v.(uint8)
}
