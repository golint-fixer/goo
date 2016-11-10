package goo

var _ Map = MapByteString(nil)

// MapByteString is a map from byte to string.
type MapByteString map[byte]string

// Delete implements Map.
func (m MapByteString) Delete(k interface{}) {
	delete(m, k.(byte))
}

// Dereference implements Map.
func (m *MapByteString) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapByteString) Equals(other Equatable) bool {
	var n = other.(MapByteString)

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
func (m MapByteString) Get(k interface{}) interface{} {
	return m[k.(byte)]
}

// GetCheck implements Map.
func (m MapByteString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(byte)]

	return v, ok
}

// KeyValues implements Map.
func (m MapByteString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapByteString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapByteString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapByteString) Make(c int) Map {
	return make(MapByteString, c)
}

// NotEquals implements Map.
func (m MapByteString) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapByteString) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapByteString) Set(k, v interface{}) {
	m[k.(byte)] = v.(string)
}
