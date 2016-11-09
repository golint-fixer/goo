package goo

var _ Map = MapStringInt(nil)

// MapStringInt is a map from string to int.
type MapStringInt map[string]int

// Delete implements Map.
func (m MapStringInt) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringInt) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInt) Make(c int) Map {
	return make(MapStringInt, c)
}

// Set implements Map.
func (m MapStringInt) Set(k, v interface{}) {
	m[k.(string)] = v.(int)
}
