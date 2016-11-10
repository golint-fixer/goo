package goo

var _ Map = MapUint8Struct(nil)

var _ Pointer = &MapUint8Struct{}

// MapUint8Struct is a map from uint8 to struct{}.
type MapUint8Struct map[uint8]struct{}

// Delete implements Map.
func (m MapUint8Struct) Delete(k interface{}) {
	delete(m, k.(uint8))
}

// Dereference implements Map.
func (m *MapUint8Struct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUint8Struct) Equals(other Equatable) bool {
	var n = other.(MapUint8Struct)

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
func (m MapUint8Struct) Get(k interface{}) interface{} {
	return m[k.(uint8)]
}

// GetCheck implements Map.
func (m MapUint8Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint8Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint8Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint8Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint8Struct) Make(c int) Map {
	return make(MapUint8Struct, c)
}

// NotEquals implements Map.
func (m MapUint8Struct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUint8Struct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUint8Struct) Set(k, v interface{}) {
	m[k.(uint8)] = v.(struct{})
}
