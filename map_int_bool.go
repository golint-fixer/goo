package goo

var _ Map = MapIntBool(nil)

// MapIntBool is a map from int to bool.
type MapIntBool map[int]bool

// Delete implements Map.
func (m MapIntBool) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntBool) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntBool) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntBool) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntBool) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntBool) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntBool) Make(c int) Map {
	return make(MapIntBool, c)
}

// Set implements Map.
func (m MapIntBool) Set(k, v interface{}) {
	m[k.(int)] = v.(bool)
}
