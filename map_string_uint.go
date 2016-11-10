package goo

var _ Map = MapStringUint(nil)

// MapStringUint is a map from string to uint.
type MapStringUint map[string]uint

// Delete implements Map.
func (m MapStringUint) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringUint) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringUint) Equals(other Equatable) bool {
	var n = other.(MapStringUint)

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
func (m MapStringUint) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUint) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUint) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUint) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUint) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUint) Make(c int) Map {
	return make(MapStringUint, c)
}

// NotEquals implements Map.
func (m MapStringUint) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringUint) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringUint) Set(k, v interface{}) {
	m[k.(string)] = v.(uint)
}
