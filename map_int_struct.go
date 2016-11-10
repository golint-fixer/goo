package goo

var _ Map = MapIntStruct(nil)

// MapIntStruct is a map from int to struct{}.
type MapIntStruct map[int]struct{}

// Delete implements Map.
func (m MapIntStruct) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntStruct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntStruct) Equals(other Equatable) bool {
	var n = other.(MapIntStruct)

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
func (m MapIntStruct) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntStruct) Make(c int) Map {
	return make(MapIntStruct, c)
}

// NotEquals implements Map.
func (m MapIntStruct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntStruct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntStruct) Set(k, v interface{}) {
	m[k.(int)] = v.(struct{})
}
