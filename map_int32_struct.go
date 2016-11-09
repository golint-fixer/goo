package goo

var _ Map = MapInt32Struct(nil)

// MapInt32Struct is a map from int32 to struct{}.
type MapInt32Struct map[int32]struct{}

// Delete implements Map.
func (m MapInt32Struct) Delete(k interface{}) {
	delete(m, k.(int32))
}

// Equals implements Map.
func (m MapInt32Struct) Equals(other Equatable) bool {
	var n = other.(MapInt32Struct)

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
func (m MapInt32Struct) Get(k interface{}) interface{} {
	return m[k.(int32)]
}

// GetCheck implements Map.
func (m MapInt32Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt32Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt32Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt32Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt32Struct) Make(c int) Map {
	return make(MapInt32Struct, c)
}

// NotEquals implements Map.
func (m MapInt32Struct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapInt32Struct) Set(k, v interface{}) {
	m[k.(int32)] = v.(struct{})
}
