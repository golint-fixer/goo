package goo

var _ Map = MapIntFloat64(nil)

// MapIntFloat64 is a map from int to float64.
type MapIntFloat64 map[int]float64

// Delete implements Map.
func (m MapIntFloat64) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntFloat64) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntFloat64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntFloat64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntFloat64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntFloat64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntFloat64) Make(c int) Map {
	return make(MapIntFloat64, c)
}

// Set implements Map.
func (m MapIntFloat64) Set(k, v interface{}) {
	m[k.(int)] = v.(float64)
}
