package goo

var _ Map = MapIntInterface(nil)

// MapIntInterface is a map from int to interface{}.
type MapIntInterface map[int]interface{}

// Delete implements Map.
func (m MapIntInterface) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntInterface) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntInterface) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntInterface) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntInterface) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntInterface) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntInterface) Make(c int) Map {
	return make(MapIntInterface, c)
}

// Set implements Map.
func (m MapIntInterface) Set(k, v interface{}) {
	m[k.(int)] = v.(interface{})
}
