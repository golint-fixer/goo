package goo

var _ Map = MapStringInt64(nil)

// MapStringInt64 is a map from string to int64.
type MapStringInt64 map[string]int64

// Delete implements Map.
func (m MapStringInt64) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringInt64) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringInt64) Equals(other Equatable) bool {
	var n = other.(MapStringInt64)

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
func (m MapStringInt64) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInt64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInt64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInt64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInt64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInt64) Make(c int) Map {
	return make(MapStringInt64, c)
}

// NotEquals implements Map.
func (m MapStringInt64) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringInt64) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringInt64) Set(k, v interface{}) {
	m[k.(string)] = v.(int64)
}
