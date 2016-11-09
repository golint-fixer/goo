package goo

var _ Map = MapIntUint32(nil)

// MapIntUint32 is a map from int to uint32.
type MapIntUint32 map[int]uint32

// Delete implements Map.
func (m MapIntUint32) Delete(k interface{}) {
	delete(m, k.(int))
}

// Equals implements Map.
func (m MapIntUint32) Equals(other Equatable) bool {
	var n = other.(MapIntUint32)

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
func (m MapIntUint32) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntUint32) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntUint32) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntUint32) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntUint32) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntUint32) Make(c int) Map {
	return make(MapIntUint32, c)
}

// NotEquals implements Map.
func (m MapIntUint32) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapIntUint32) Set(k, v interface{}) {
	m[k.(int)] = v.(uint32)
}
