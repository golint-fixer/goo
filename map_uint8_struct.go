package goo

var _ Map = MapUint8Struct(nil)

// MapUint8Struct is a map from uint8 to struct{}.
type MapUint8Struct map[uint8]struct{}

// Delete implements Map.
func (m MapUint8Struct) Delete(k interface{}) {
	delete(m, k.(uint8))
}

// Get implements Map.
func (m MapUint8Struct) Get(k interface{}) interface{} {
	return m[k.(uint8)]
}

// GetCheck implements Map.
func (m MapUint8Struct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint8Struct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint8Struct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint8Struct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint8Struct) Make(c int) Map {
	return make(MapUint8Struct, c)
}

// Set implements Map.
func (m MapUint8Struct) Set(k, v interface{}) {
	m[k.(uint8)] = v.(struct{})
}
