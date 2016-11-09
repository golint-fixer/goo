package goo

var _ Map = MapIntUint64(nil)

// MapIntUint64 is a map from int to uint64.
type MapIntUint64 map[int]uint64

// Delete implements Map.
func (m MapIntUint64) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntUint64) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntUint64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntUint64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntUint64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntUint64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntUint64) Make(c int) Map {
	return make(MapIntUint64, c)
}

// Set implements Map.
func (m MapIntUint64) Set(k, v interface{}) {
	m[k.(int)] = v.(uint64)
}
