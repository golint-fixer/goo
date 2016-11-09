package goo

var _ Map = MapUintInt(nil)

// MapUintInt is a map from uint to int.
type MapUintInt map[uint]int

// Delete implements Map.
func (m MapUintInt) Delete(k interface{}) {
	delete(m, k.(uint))
}

// Equals implements Map.
func (m MapUintInt) Equals(other Equatable) bool {
	var n = other.(MapUintInt)

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
func (m MapUintInt) Get(k interface{}) interface{} {
	return m[k.(uint)]
}

// GetCheck implements Map.
func (m MapUintInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUintInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUintInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUintInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUintInt) Make(c int) Map {
	return make(MapUintInt, c)
}

// NotEquals implements Map.
func (m MapUintInt) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapUintInt) Set(k, v interface{}) {
	m[k.(uint)] = v.(int)
}
