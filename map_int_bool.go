package goo

var _ Map = MapIntBool(nil)

var _ Pointer = &MapIntBool{}

// MapIntBool is a map from int to bool.
type MapIntBool map[int]bool

// Delete implements Map.
func (m MapIntBool) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntBool) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntBool) Equals(other Equatable) bool {
	var n = other.(MapIntBool)

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
func (m MapIntBool) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntBool) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntBool) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntBool) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntBool) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntBool) Make(c int) Map {
	return make(MapIntBool, c)
}

// NotEquals implements Map.
func (m MapIntBool) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntBool) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntBool) Set(k, v interface{}) {
	m[k.(int)] = v.(bool)
}
