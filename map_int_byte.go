package goo

var _ Map = MapIntByte(nil)

// MapIntByte is a map from int to byte.
type MapIntByte map[int]byte

// Delete implements Map.
func (m MapIntByte) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntByte) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntByte) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntByte) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntByte) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntByte) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntByte) Make(c int) Map {
	return make(MapIntByte, c)
}

// Set implements Map.
func (m MapIntByte) Set(k, v interface{}) {
	m[k.(int)] = v.(byte)
}
