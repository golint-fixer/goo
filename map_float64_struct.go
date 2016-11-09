package goo

var _ Map = MapFloat64Struct(nil)

// MapFloat64Struct is a map from float64 to struct{}.
type MapFloat64Struct map[float64]struct{}

// Delete implements Map.
func (m MapFloat64Struct) Delete(k interface{}) {
	delete(m, k.(float64))
}

// Get implements Map.
func (m MapFloat64Struct) Get(k interface{}) interface{} {
	return m[k.(float64)]
}

// GetCheck implements Map.
func (m MapFloat64Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(float64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapFloat64Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapFloat64Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapFloat64Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapFloat64Struct) Make(c int) Map {
	return make(MapFloat64Struct, c)
}

// Set implements Map.
func (m MapFloat64Struct) Set(k, v interface{}) {
	m[k.(float64)] = v.(struct{})
}
