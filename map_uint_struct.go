package goo

var _ Map = MapUintStruct(nil)

// MapUintStruct is a map from uint to struct{}.
type MapUintStruct map[uint]struct{}

// Delete implements Map.
func (m MapUintStruct) Delete(k interface{}) {
	delete(m, k.(uint))
}

// Equals implements Map.
func (m MapUintStruct) Equals(other Equatable) bool {
	var n = other.(MapUintStruct)

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
func (m MapUintStruct) Get(k interface{}) interface{} {
	return m[k.(uint)]
}

// GetCheck implements Map.
func (m MapUintStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUintStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUintStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUintStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUintStruct) Make(c int) Map {
	return make(MapUintStruct, c)
}

// NotEquals implements Map.
func (m MapUintStruct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapUintStruct) Set(k, v interface{}) {
	m[k.(uint)] = v.(struct{})
}
