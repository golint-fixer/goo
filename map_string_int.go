package goo

var _ Map = MapStringInt(nil)

var _ Pointer = &MapStringInt{}

// MapStringInt is a map from string to int.
type MapStringInt map[string]int

// Delete implements Map.
func (m MapStringInt) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringInt) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringInt) Equals(other Equatable) bool {
	var n = other.(MapStringInt)

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
func (m MapStringInt) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInt) Make(c int) Map {
	return make(MapStringInt, c)
}

// NotEquals implements Map.
func (m MapStringInt) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringInt) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringInt) Set(k, v interface{}) {
	m[k.(string)] = v.(int)
}
