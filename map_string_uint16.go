package goo

var _ Map = MapStringUint16(nil)

// MapStringUint16 is a map from string to uint16.
type MapStringUint16 map[string]uint16

// Delete implements Map.
func (m MapStringUint16) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringUint16) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUint16) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUint16) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUint16) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUint16) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUint16) Make(c int) Map {
	return make(MapStringUint16, c)
}

// Set implements Map.
func (m MapStringUint16) Set(k, v interface{}) {
	m[k.(string)] = v.(uint16)
}
