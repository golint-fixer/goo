package goo

var _ Map = MapUintptrString(nil)

// MapUintptrString is a map from uintptr to string.
type MapUintptrString map[uintptr]string

// Delete implements Map.
func (m MapUintptrString) Delete(k interface{}) {
	delete(m, k.(uintptr))
}

// Get implements Map.
func (m MapUintptrString) Get(k interface{}) interface{} {
	return m[k.(uintptr)]
}

// GetCheck implements Map.
func (m MapUintptrString) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uintptr)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUintptrString) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUintptrString) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUintptrString) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUintptrString) Make(c int) Map {
	return make(MapUintptrString, c)
}

// Set implements Map.
func (m MapUintptrString) Set(k, v interface{}) {
	m[k.(uintptr)] = v.(string)
}
