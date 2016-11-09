package goo

var _ Map = MapStringUint(nil)

// MapStringUint is a map from string to uint.
type MapStringUint map[string]uint

// Delete implements Map.
func (m MapStringUint) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringUint) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUint) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUint) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUint) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUint) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUint) Make(c int) Map {
	return make(MapStringUint, c)
}

// Set implements Map.
func (m MapStringUint) Set(k, v interface{}) {
	m[k.(string)] = v.(uint)
}
