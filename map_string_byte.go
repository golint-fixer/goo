package goo

var _ Map = MapStringByte(nil)

// MapStringByte is a map from string to byte.
type MapStringByte map[string]byte

// Delete implements Map.
func (m MapStringByte) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringByte) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringByte) Equals(other Equatable) bool {
	var n = other.(MapStringByte)

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
func (m MapStringByte) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringByte) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringByte) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringByte) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringByte) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringByte) Make(c int) Map {
	return make(MapStringByte, c)
}

// NotEquals implements Map.
func (m MapStringByte) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringByte) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringByte) Set(k, v interface{}) {
	m[k.(string)] = v.(byte)
}
