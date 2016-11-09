package goo

var _ Map = MapUint32String(nil)

// MapUint32String is a map from uint32 to string.
type MapUint32String map[uint32]string

// Delete implements Map.
func (m MapUint32String) Delete(k interface{}) {
	delete(m, k.(uint32))
}

// Get implements Map.
func (m MapUint32String) Get(k interface{}) interface{} {
	return m[k.(uint32)]
}

// GetCheck implements Map.
func (m MapUint32String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint32String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint32String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint32String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint32String) Make(c int) Map {
	return make(MapUint32String, c)
}

// Set implements Map.
func (m MapUint32String) Set(k, v interface{}) {
	m[k.(uint32)] = v.(string)
}
