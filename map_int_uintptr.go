package goo

var _ Map = MapIntUintptr(nil)

var _ Pointer = &MapIntUintptr{}

// MapIntUintptr is a map from int to uintptr.
type MapIntUintptr map[int]uintptr

// Delete implements Map.
func (m MapIntUintptr) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntUintptr) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntUintptr) Equals(other Equatable) bool {
	var n = other.(MapIntUintptr)

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
func (m MapIntUintptr) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntUintptr) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntUintptr) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntUintptr) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntUintptr) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntUintptr) Make(c int) Map {
	return make(MapIntUintptr, c)
}

// NotEquals implements Map.
func (m MapIntUintptr) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntUintptr) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntUintptr) Set(k, v interface{}) {
	m[k.(int)] = v.(uintptr)
}
