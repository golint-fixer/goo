package goo

var _ Map = MapFloat64String(nil)

var _ Pointer = &MapFloat64String{}

// MapFloat64String is a map from float64 to string.
type MapFloat64String map[float64]string

// Delete implements Map.
func (m MapFloat64String) Delete(k interface{}) {
	delete(m, k.(float64))
}

// Dereference implements Map.
func (m *MapFloat64String) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapFloat64String) Equals(other Equatable) bool {
	var n = other.(MapFloat64String)

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
func (m MapFloat64String) Get(k interface{}) interface{} {
	return m[k.(float64)]
}

// GetCheck implements Map.
func (m MapFloat64String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(float64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapFloat64String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapFloat64String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapFloat64String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapFloat64String) Make(c int) Map {
	return make(MapFloat64String, c)
}

// NotEquals implements Map.
func (m MapFloat64String) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapFloat64String) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapFloat64String) Set(k, v interface{}) {
	m[k.(float64)] = v.(string)
}
