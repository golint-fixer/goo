package goo

var _ Map = MapFloat32Struct(nil)

// MapFloat32Struct is a map from float32 to struct{}.
type MapFloat32Struct map[float32]struct{}

// Delete implements Map.
func (m MapFloat32Struct) Delete(k interface{}) {
	delete(m, k.(float32))
}

// Dereference implements Map.
func (m *MapFloat32Struct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapFloat32Struct) Equals(other Equatable) bool {
	var n = other.(MapFloat32Struct)

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
func (m MapFloat32Struct) Get(k interface{}) interface{} {
	return m[k.(float32)]
}

// GetCheck implements Map.
func (m MapFloat32Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(float32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapFloat32Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapFloat32Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapFloat32Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapFloat32Struct) Make(c int) Map {
	return make(MapFloat32Struct, c)
}

// NotEquals implements Map.
func (m MapFloat32Struct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapFloat32Struct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapFloat32Struct) Set(k, v interface{}) {
	m[k.(float32)] = v.(struct{})
}
