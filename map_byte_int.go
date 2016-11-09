package goo

var _ Map = MapByteInt(nil)

// MapByteInt is a map from byte to int.
type MapByteInt map[byte]int

// Delete implements Map.
func (m MapByteInt) Delete(k interface{}) {
	delete(m, k.(byte))
}

// Get implements Map.
func (m MapByteInt) Get(k interface{}) interface{} {
	return m[k.(byte)]
}

// GetCheck implements Map.
func (m MapByteInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(byte)]

	return v, ok
}

// KeyValues implements Map.
func (m MapByteInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapByteInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapByteInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapByteInt) Make(c int) Map {
	return make(MapByteInt, c)
}

// Set implements Map.
func (m MapByteInt) Set(k, v interface{}) {
	m[k.(byte)] = v.(int)
}
