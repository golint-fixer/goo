package goo

var _ Map = MapUint32Struct(nil)

var _ Pointer = &MapUint32Struct{}

// MapUint32Struct is a map from uint32 to struct{}.
type MapUint32Struct map[uint32]struct{}

// Delete implements Map.
func (m MapUint32Struct) Delete(k interface{}) {
	delete(m, k.(uint32))
}

// Dereference implements Map.
func (m *MapUint32Struct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUint32Struct) Equals(other Equatable) bool {
	var n = other.(MapUint32Struct)

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
func (m MapUint32Struct) Get(k interface{}) interface{} {
	return m[k.(uint32)]
}

// GetCheck implements Map.
func (m MapUint32Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint32Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint32Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint32Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint32Struct) Make(c int) Map {
	return make(MapUint32Struct, c)
}

// NotEquals implements Map.
func (m MapUint32Struct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUint32Struct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUint32Struct) Set(k, v interface{}) {
	m[k.(uint32)] = v.(struct{})
}
