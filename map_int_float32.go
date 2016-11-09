package goo

var _ Map = MapIntFloat32(nil)

// MapIntFloat32 is a map from int to float32.
type MapIntFloat32 map[int]float32

// Delete implements Map.
func (m MapIntFloat32) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntFloat32) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntFloat32) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntFloat32) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntFloat32) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntFloat32) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntFloat32) Make(c int) Map {
	return make(MapIntFloat32, c)
}

// Set implements Map.
func (m MapIntFloat32) Set(k, v interface{}) {
	m[k.(int)] = v.(float32)
}
