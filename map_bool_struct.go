package goo

var _ Map = MapBoolStruct(nil)

// MapBoolStruct is a map from bool to struct{}.
type MapBoolStruct map[bool]struct{}

// Delete implements Map.
func (m MapBoolStruct) Delete(k interface{}) {
	delete(m, k.(bool))
}

// Get implements Map.
func (m MapBoolStruct) Get(k interface{}) interface{} {
	return m[k.(bool)]
}

// GetCheck implements Map.
func (m MapBoolStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(bool)]

	return v, ok
}

// KeyValues implements Map.
func (m MapBoolStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapBoolStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapBoolStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapBoolStruct) Make(c int) Map {
	return make(MapBoolStruct, c)
}

// Set implements Map.
func (m MapBoolStruct) Set(k, v interface{}) {
	m[k.(bool)] = v.(struct{})
}
