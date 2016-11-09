package goo

var _ Map = MapIntInt32(nil)

// MapIntInt32 is a map from int to int32.
type MapIntInt32 map[int]int32

// Delete implements Map.
func (m MapIntInt32) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntInt32) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntInt32) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntInt32) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntInt32) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntInt32) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntInt32) Make(c int) Map {
	return make(MapIntInt32, c)
}

// Set implements Map.
func (m MapIntInt32) Set(k, v interface{}) {
	m[k.(int)] = v.(int32)
}
