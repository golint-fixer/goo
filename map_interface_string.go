package goo

var _ Map = MapInterfaceString(nil)

// MapInterfaceString is a map from interface{} to string.
type MapInterfaceString map[interface{}]string

// Delete implements Map.
func (m MapInterfaceString) Delete(k interface{}) {
	delete(m, k.(interface{}))
}

// Equals implements Map.
func (m MapInterfaceString) Equals(other Equatable) bool {
	var n = other.(MapInterfaceString)

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
func (m MapInterfaceString) Get(k interface{}) interface{} {
	return m[k.(interface{})]
}

// GetCheck implements Map.
func (m MapInterfaceString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(interface{})]

	return v, ok
}

// KeyValues implements Map.
func (m MapInterfaceString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInterfaceString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInterfaceString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInterfaceString) Make(c int) Map {
	return make(MapInterfaceString, c)
}

// NotEquals implements Map.
func (m MapInterfaceString) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapInterfaceString) Set(k, v interface{}) {
	m[k.(interface{})] = v.(string)
}
