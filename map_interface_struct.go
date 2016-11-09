package goo

var _ Map = MapInterfaceStruct(nil)

// MapInterfaceStruct is a map from interface{} to struct{}.
type MapInterfaceStruct map[interface{}]struct{}

// Delete implements Map.
func (m MapInterfaceStruct) Delete(k interface{}) {
	delete(m, k.(interface{}))
}

// Get implements Map.
func (m MapInterfaceStruct) Get(k interface{}) interface{} {
	return m[k.(interface{})]
}

// GetCheck implements Map.
func (m MapInterfaceStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(interface{})]

	return v, ok
}

// KeyValues implements Map.
func (m MapInterfaceStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInterfaceStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInterfaceStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInterfaceStruct) Make(c int) Map {
	return make(MapInterfaceStruct, c)
}

// Set implements Map.
func (m MapInterfaceStruct) Set(k, v interface{}) {
	m[k.(interface{})] = v.(struct{})
}
