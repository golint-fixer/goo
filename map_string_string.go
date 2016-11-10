package goo

var _ Map = MapStringString(nil)

var _ Pointer = &MapStringString{}

// MapStringString is a map from string to string.
type MapStringString map[string]string

// Delete implements Map.
func (m MapStringString) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringString) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringString) Equals(other Equatable) bool {
	var n = other.(MapStringString)

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
func (m MapStringString) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringString) Make(c int) Map {
	return make(MapStringString, c)
}

// NotEquals implements Map.
func (m MapStringString) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringString) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringString) Set(k, v interface{}) {
	m[k.(string)] = v.(string)
}
