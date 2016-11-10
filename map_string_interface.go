package goo

var _ Map = MapStringInterface(nil)

var _ Pointer = &MapStringInterface{}

// MapStringInterface is a map from string to interface{}.
type MapStringInterface map[string]interface{}

// Delete implements Map.
func (m MapStringInterface) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringInterface) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringInterface) Equals(other Equatable) bool {
	var n = other.(MapStringInterface)

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
func (m MapStringInterface) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInterface) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInterface) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInterface) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInterface) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInterface) Make(c int) Map {
	return make(MapStringInterface, c)
}

// NotEquals implements Map.
func (m MapStringInterface) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringInterface) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringInterface) Set(k, v interface{}) {
	m[k.(string)] = v.(interface{})
}
