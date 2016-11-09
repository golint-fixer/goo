package goo

var _ Map = MapStringUint64(nil)

// MapStringUint64 is a map from string to uint64.
type MapStringUint64 map[string]uint64

// Delete implements Map.
func (m MapStringUint64) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringUint64) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUint64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUint64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUint64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUint64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUint64) Make(c int) Map {
	return make(MapStringUint64, c)
}

// Set implements Map.
func (m MapStringUint64) Set(k, v interface{}) {
	m[k.(string)] = v.(uint64)
}
