package goo

var _ Map = MapInt16Struct(nil)

// MapInt16Struct is a map from int16 to struct{}.
type MapInt16Struct map[int16]struct{}

// Delete implements Map.
func (m MapInt16Struct) Delete(k interface{}) {
	delete(m, k.(int16))
}

// Get implements Map.
func (m MapInt16Struct) Get(k interface{}) interface{} {
	return m[k.(int16)]
}

// GetCheck implements Map.
func (m MapInt16Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int16)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt16Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt16Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt16Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt16Struct) Make(c int) Map {
	return make(MapInt16Struct, c)
}

// Set implements Map.
func (m MapInt16Struct) Set(k, v interface{}) {
	m[k.(int16)] = v.(struct{})
}
