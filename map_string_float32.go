package goo

var _ Map = MapStringFloat32(nil)

// MapStringFloat32 is a map from string to float32.
type MapStringFloat32 map[string]float32

// Delete implements Map.
func (m MapStringFloat32) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringFloat32) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringFloat32) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringFloat32) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringFloat32) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringFloat32) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringFloat32) Make(c int) Map {
	return make(MapStringFloat32, c)
}

// Set implements Map.
func (m MapStringFloat32) Set(k, v interface{}) {
	m[k.(string)] = v.(float32)
}
