package goo

var _ Map = MapIntInt(nil)

// MapIntInt is a map from int to int.
type MapIntInt map[int]int

// Delete implements Map.
func (m MapIntInt) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntInt) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntInt) Make(c int) Map {
	return make(MapIntInt, c)
}

// Set implements Map.
func (m MapIntInt) Set(k, v interface{}) {
	m[k.(int)] = v.(int)
}
