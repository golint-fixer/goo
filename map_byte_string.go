package goo

var _ Map = MapByteString(nil)

// MapByteString is a map from byte to string.
type MapByteString map[byte]string

// Delete implements Map.
func (m MapByteString) Delete(k interface{}) {
	delete(m, k.(byte))
}

// Get implements Map.
func (m MapByteString) Get(k interface{}) interface{} {
	return m[k.(byte)]
}

// GetCheck implements Map.
func (m MapByteString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(byte)]

	return v, ok
}

// KeyValues implements Map.
func (m MapByteString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapByteString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapByteString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapByteString) Make(c int) Map {
	return make(MapByteString, c)
}

// Set implements Map.
func (m MapByteString) Set(k, v interface{}) {
	m[k.(byte)] = v.(string)
}
