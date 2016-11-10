package goo

var _ Map = MapStringUint8(nil)

var _ Pointer = &MapStringUint8{}

// MapStringUint8 is a map from string to uint8.
type MapStringUint8 map[string]uint8

// Delete implements Map.
func (m MapStringUint8) Delete(k interface{}) {
	delete(m, k.(string))
}

// Dereference implements Map.
func (m *MapStringUint8) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapStringUint8) Equals(other Equatable) bool {
	var n = other.(MapStringUint8)

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
func (m MapStringUint8) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUint8) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUint8) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUint8) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUint8) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUint8) Make(c int) Map {
	return make(MapStringUint8, c)
}

// NotEquals implements Map.
func (m MapStringUint8) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapStringUint8) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapStringUint8) Set(k, v interface{}) {
	m[k.(string)] = v.(uint8)
}
