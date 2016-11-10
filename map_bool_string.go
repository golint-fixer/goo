package goo

var _ Map = MapBoolString(nil)

var _ Pointer = &MapBoolString{}

// MapBoolString is a map from bool to string.
type MapBoolString map[bool]string

// Delete implements Map.
func (m MapBoolString) Delete(k interface{}) {
	delete(m, k.(bool))
}

// Dereference implements Map.
func (m *MapBoolString) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapBoolString) Equals(other Equatable) bool {
	var n = other.(MapBoolString)

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
func (m MapBoolString) Get(k interface{}) interface{} {
	return m[k.(bool)]
}

// GetCheck implements Map.
func (m MapBoolString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(bool)]

	return v, ok
}

// KeyValues implements Map.
func (m MapBoolString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapBoolString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapBoolString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapBoolString) Make(c int) Map {
	return make(MapBoolString, c)
}

// NotEquals implements Map.
func (m MapBoolString) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapBoolString) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapBoolString) Set(k, v interface{}) {
	m[k.(bool)] = v.(string)
}
