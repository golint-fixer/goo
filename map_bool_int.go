package goo

var _ Map = MapBoolInt(nil)

// MapBoolInt is a map from bool to int.
type MapBoolInt map[bool]int

// Delete implements Map.
func (m MapBoolInt) Delete(k interface{}) {
	delete(m, k.(bool))
}

// Equals implements Map.
func (m MapBoolInt) Equals(other Equatable) bool {
	var n = other.(MapBoolInt)

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
func (m MapBoolInt) Get(k interface{}) interface{} {
	return m[k.(bool)]
}

// GetCheck implements Map.
func (m MapBoolInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(bool)]

	return v, ok
}

// KeyValues implements Map.
func (m MapBoolInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapBoolInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapBoolInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapBoolInt) Make(c int) Map {
	return make(MapBoolInt, c)
}

// NotEquals implements Map.
func (m MapBoolInt) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapBoolInt) Set(k, v interface{}) {
	m[k.(bool)] = v.(int)
}
