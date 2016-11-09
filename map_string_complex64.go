package goo

var _ Map = MapStringComplex64(nil)

// MapStringComplex64 is a map from string to complex64.
type MapStringComplex64 map[string]complex64

// Delete implements Map.
func (m MapStringComplex64) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringComplex64) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringComplex64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringComplex64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringComplex64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringComplex64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringComplex64) Make(c int) Map {
	return make(MapStringComplex64, c)
}

// Set implements Map.
func (m MapStringComplex64) Set(k, v interface{}) {
	m[k.(string)] = v.(complex64)
}
