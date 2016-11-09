package goo

var _ Map = MapStringUint32(nil)

// MapStringUint32 is a map from string to uint32.
type MapStringUint32 map[string]uint32

// Delete implements Map.
func (m MapStringUint32) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringUint32) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUint32) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUint32) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUint32) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUint32) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUint32) Make(c int) Map {
	return make(MapStringUint32, c)
}

// Set implements Map.
func (m MapStringUint32) Set(k, v interface{}) {
	m[k.(string)] = v.(uint32)
}
