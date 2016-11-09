package goo

var _ Map = MapStringByte(nil)

// MapStringByte is a map from string to byte.
type MapStringByte map[string]byte

// Delete implements Map.
func (m MapStringByte) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringByte) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringByte) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringByte) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringByte) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringByte) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringByte) Make(c int) Map {
	return make(MapStringByte, c)
}

// Set implements Map.
func (m MapStringByte) Set(k, v interface{}) {
	m[k.(string)] = v.(byte)
}
