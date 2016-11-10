package goo

var _ Map = MapUintptrStruct(nil)

var _ Pointer = &MapUintptrStruct{}

// MapUintptrStruct is a map from uintptr to struct{}.
type MapUintptrStruct map[uintptr]struct{}

// Delete implements Map.
func (m MapUintptrStruct) Delete(k interface{}) {
	delete(m, k.(uintptr))
}

// Dereference implements Map.
func (m *MapUintptrStruct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUintptrStruct) Equals(other Equatable) bool {
	var n = other.(MapUintptrStruct)

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
func (m MapUintptrStruct) Get(k interface{}) interface{} {
	return m[k.(uintptr)]
}

// GetCheck implements Map.
func (m MapUintptrStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uintptr)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUintptrStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUintptrStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUintptrStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUintptrStruct) Make(c int) Map {
	return make(MapUintptrStruct, c)
}

// NotEquals implements Map.
func (m MapUintptrStruct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUintptrStruct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUintptrStruct) Set(k, v interface{}) {
	m[k.(uintptr)] = v.(struct{})
}
