package goo

var _ Map = MapRuneString(nil)

// MapRuneString is a map from rune to string.
type MapRuneString map[rune]string

// Delete implements Map.
func (m MapRuneString) Delete(k interface{}) {
	delete(m, k.(rune))
}

// Get implements Map.
func (m MapRuneString) Get(k interface{}) interface{} {
	return m[k.(rune)]
}

// GetCheck implements Map.
func (m MapRuneString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(rune)]

	return v, ok
}

// KeyValues implements Map.
func (m MapRuneString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapRuneString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapRuneString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapRuneString) Make(c int) Map {
	return make(MapRuneString, c)
}

// Set implements Map.
func (m MapRuneString) Set(k, v interface{}) {
	m[k.(rune)] = v.(string)
}
