package goo

var _ Map = MapInt32String(nil)

// MapInt32String is a map from int32 to string.
type MapInt32String map[int32]string

// Delete implements Map.
func (m MapInt32String) Delete(k interface{}) {
	delete(m, k.(int32))
}

// Get implements Map.
func (m MapInt32String) Get(k interface{}) interface{} {
	return m[k.(int32)]
}

// GetCheck implements Map.
func (m MapInt32String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt32String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt32String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt32String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt32String) Make(c int) Map {
	return make(MapInt32String, c)
}

// Set implements Map.
func (m MapInt32String) Set(k, v interface{}) {
	m[k.(int32)] = v.(string)
}
