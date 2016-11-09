package goo

var _ Map = MapStringStruct(nil)

// MapStringStruct is a map from string to struct{}.
type MapStringStruct map[string]struct{}

// Delete implements Map.
func (m MapStringStruct) Delete(k interface{}) {
	delete(m, k.(string))
}

// Equals implements Map.
func (m MapStringStruct) Equals(other Equatable) bool {
	var n = other.(MapStringStruct)

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
func (m MapStringStruct) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringStruct) Make(c int) Map {
	return make(MapStringStruct, c)
}

// NotEquals implements Map.
func (m MapStringStruct) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapStringStruct) Set(k, v interface{}) {
	m[k.(string)] = v.(struct{})
}
