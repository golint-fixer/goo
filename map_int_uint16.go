package goo

var _ Map = MapIntUint16(nil)

// MapIntUint16 is a map from int to uint16.
type MapIntUint16 map[int]uint16

// Delete implements Map.
func (m MapIntUint16) Delete(k interface{}) {
	delete(m, k.(int))
}

// Equals implements Map.
func (m MapIntUint16) Equals(other Equatable) bool {
	var n = other.(MapIntUint16)

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
func (m MapIntUint16) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntUint16) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntUint16) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntUint16) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntUint16) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntUint16) Make(c int) Map {
	return make(MapIntUint16, c)
}

// NotEquals implements Map.
func (m MapIntUint16) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapIntUint16) Set(k, v interface{}) {
	m[k.(int)] = v.(uint16)
}
