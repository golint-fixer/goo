package goo

var _ Map = MapIntString(nil)

// MapIntString is a map from int to string.
type MapIntString map[int]string

// Delete implements Map.
func (m MapIntString) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntString) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntString) Equals(other Equatable) bool {
	var n = other.(MapIntString)

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
func (m MapIntString) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntString) Make(c int) Map {
	return make(MapIntString, c)
}

// NotEquals implements Map.
func (m MapIntString) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntString) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntString) Set(k, v interface{}) {
	m[k.(int)] = v.(string)
}
