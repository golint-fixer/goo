package goo

var _ Map = MapIntComplex64(nil)

// MapIntComplex64 is a map from int to complex64.
type MapIntComplex64 map[int]complex64

// Delete implements Map.
func (m MapIntComplex64) Delete(k interface{}) {
	delete(m, k.(int))
}

// Get implements Map.
func (m MapIntComplex64) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntComplex64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntComplex64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntComplex64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntComplex64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntComplex64) Make(c int) Map {
	return make(MapIntComplex64, c)
}

// Set implements Map.
func (m MapIntComplex64) Set(k, v interface{}) {
	m[k.(int)] = v.(complex64)
}
