package goo

var _ Map = MapByteStruct(nil)

// MapByteStruct is a map from byte to struct{}.
type MapByteStruct map[byte]struct{}

// Delete implements Map.
func (m MapByteStruct) Delete(k interface{}) {
	delete(m, k.(byte))
}

// Equals implements Map.
func (m MapByteStruct) Equals(other Equatable) bool {
	var n = other.(MapByteStruct)

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
func (m MapByteStruct) Get(k interface{}) interface{} {
	return m[k.(byte)]
}

// GetCheck implements Map.
func (m MapByteStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(byte)]

	return v, ok
}

// KeyValues implements Map.
func (m MapByteStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapByteStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapByteStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapByteStruct) Make(c int) Map {
	return make(MapByteStruct, c)
}

// NotEquals implements Map.
func (m MapByteStruct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapByteStruct) Set(k, v interface{}) {
	m[k.(byte)] = v.(struct{})
}
