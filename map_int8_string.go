package goo

var _ Map = MapInt8String(nil)

// MapInt8String is a map from int8 to string.
type MapInt8String map[int8]string

// Delete implements Map.
func (m MapInt8String) Delete(k interface{}) {
	delete(m, k.(int8))
}

// Dereference implements Map.
func (m *MapInt8String) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapInt8String) Equals(other Equatable) bool {
	var n = other.(MapInt8String)

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
func (m MapInt8String) Get(k interface{}) interface{} {
	return m[k.(int8)]
}

// GetCheck implements Map.
func (m MapInt8String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt8String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt8String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt8String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt8String) Make(c int) Map {
	return make(MapInt8String, c)
}

// NotEquals implements Map.
func (m MapInt8String) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapInt8String) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapInt8String) Set(k, v interface{}) {
	m[k.(int8)] = v.(string)
}
