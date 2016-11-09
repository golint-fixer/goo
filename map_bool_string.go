package goo

var _ Map = MapBoolString(nil)

// MapBoolString is a map from bool to string.
type MapBoolString map[bool]string

// Delete implements Map.
func (m MapBoolString) Delete(k interface{}) {
	delete(m, k.(bool))
}

// Get implements Map.
func (m MapBoolString) Get(k interface{}) interface{} {
	return m[k.(bool)]
}

// GetCheck implements Map.
func (m MapBoolString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(bool)]

	return v, ok
}

// KeyValues implements Map.
func (m MapBoolString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapBoolString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapBoolString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapBoolString) Make(c int) Map {
	return make(MapBoolString, c)
}

// Set implements Map.
func (m MapBoolString) Set(k, v interface{}) {
	m[k.(bool)] = v.(string)
}
