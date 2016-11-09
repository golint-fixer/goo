package goo

var _ Map = MapUintString(nil)

// MapUintString is a map from uint to string.
type MapUintString map[uint]string

// Delete implements Map.
func (m MapUintString) Delete(k interface{}) {
	delete(m, k.(uint))
}

// Equals implements Map.
func (m MapUintString) Equals(other Equatable) bool {
	var n = other.(MapUintString)

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
func (m MapUintString) Get(k interface{}) interface{} {
	return m[k.(uint)]
}

// GetCheck implements Map.
func (m MapUintString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUintString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUintString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUintString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUintString) Make(c int) Map {
	return make(MapUintString, c)
}

// NotEquals implements Map.
func (m MapUintString) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapUintString) Set(k, v interface{}) {
	m[k.(uint)] = v.(string)
}
