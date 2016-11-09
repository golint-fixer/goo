package goo

var _ Map = MapRuneStruct(nil)

// MapRuneStruct is a map from rune to struct{}.
type MapRuneStruct map[rune]struct{}

// Delete implements Map.
func (m MapRuneStruct) Delete(k interface{}) {
	delete(m, k.(rune))
}

// Get implements Map.
func (m MapRuneStruct) Get(k interface{}) interface{} {
	return m[k.(rune)]
}

// GetCheck implements Map.
func (m MapRuneStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(rune)]

	return v, ok
}

// KeyValues implements Map.
func (m MapRuneStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapRuneStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapRuneStruct) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapRuneStruct) Make(c int) Map {
	return make(MapRuneStruct, c)
}

// Set implements Map.
func (m MapRuneStruct) Set(k, v interface{}) {
	m[k.(rune)] = v.(struct{})
}
