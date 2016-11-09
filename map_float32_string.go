package goo

var _ Map = MapFloat32String(nil)

// MapFloat32String is a map from float32 to string.
type MapFloat32String map[float32]string

// Delete implements Map.
func (m MapFloat32String) Delete(k interface{}) {
	delete(m, k.(float32))
}

// Get implements Map.
func (m MapFloat32String) Get(k interface{}) interface{} {
	return m[k.(float32)]
}

// GetCheck implements Map.
func (m MapFloat32String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(float32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapFloat32String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapFloat32String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapFloat32String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapFloat32String) Make(c int) Map {
	return make(MapFloat32String, c)
}

// Set implements Map.
func (m MapFloat32String) Set(k, v interface{}) {
	m[k.(float32)] = v.(string)
}
