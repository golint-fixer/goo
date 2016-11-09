package goo

var _ Map = MapStringUintptr(nil)

// MapStringUintptr is a map from string to uintptr.
type MapStringUintptr map[string]uintptr

// Delete implements Map.
func (m MapStringUintptr) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringUintptr) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringUintptr) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringUintptr) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringUintptr) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringUintptr) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringUintptr) Make(c int) Map {
	return make(MapStringUintptr, c)
}

// Set implements Map.
func (m MapStringUintptr) Set(k, v interface{}) {
	m[k.(string)] = v.(uintptr)
}
