package goo

var _ Map = MapStringInt16(nil)

// MapStringInt16 is a map from string to int16.
type MapStringInt16 map[string]int16

// Delete implements Map.
func (m MapStringInt16) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringInt16) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInt16) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInt16) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInt16) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInt16) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInt16) Make(c int) Map {
	return make(MapStringInt16, c)
}

// Set implements Map.
func (m MapStringInt16) Set(k, v interface{}) {
	m[k.(string)] = v.(int16)
}
