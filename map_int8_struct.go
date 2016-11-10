package goo

var _ Map = MapInt8Struct(nil)

var _ Pointer = &MapInt8Struct{}

// MapInt8Struct is a map from int8 to struct{}.
type MapInt8Struct map[int8]struct{}

// Delete implements Map.
func (m MapInt8Struct) Delete(k interface{}) {
	delete(m, k.(int8))
}

// Dereference implements Map.
func (m *MapInt8Struct) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapInt8Struct) Equals(other Equatable) bool {
	var n = other.(MapInt8Struct)

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
func (m MapInt8Struct) Get(k interface{}) interface{} {
	return m[k.(int8)]
}

// GetCheck implements Map.
func (m MapInt8Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt8Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt8Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt8Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt8Struct) Make(c int) Map {
	return make(MapInt8Struct, c)
}

// NotEquals implements Map.
func (m MapInt8Struct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapInt8Struct) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapInt8Struct) Set(k, v interface{}) {
	m[k.(int8)] = v.(struct{})
}
