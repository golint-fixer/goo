package goo

var _ Map = MapRuneInt(nil)

// MapRuneInt is a map from rune to int.
type MapRuneInt map[rune]int

// Delete implements Map.
func (m MapRuneInt) Delete(k interface{}) {
	delete(m, k.(rune))
}

// Get implements Map.
func (m MapRuneInt) Get(k interface{}) interface{} {
	return m[k.(rune)]
}

// GetCheck implements Map.
func (m MapRuneInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(rune)]

	return v, ok
}

// KeyValues implements Map.
func (m MapRuneInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapRuneInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapRuneInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapRuneInt) Make(c int) Map {
	return make(MapRuneInt, c)
}

// Set implements Map.
func (m MapRuneInt) Set(k, v interface{}) {
	m[k.(rune)] = v.(int)
}
