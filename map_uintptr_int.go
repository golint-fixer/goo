package goo

var _ Map = MapUintptrInt(nil)

// MapUintptrInt is a map from uintptr to int.
type MapUintptrInt map[uintptr]int

// Delete implements Map.
func (m MapUintptrInt) Delete(k interface{}) {
	delete(m, k.(uintptr))
}

// Get implements Map.
func (m MapUintptrInt) Get(k interface{}) interface{} {
	return m[k.(uintptr)]
}

// GetCheck implements Map.
func (m MapUintptrInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uintptr)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUintptrInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUintptrInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUintptrInt) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUintptrInt) Make(c int) Map {
	return make(MapUintptrInt, c)
}

// Set implements Map.
func (m MapUintptrInt) Set(k, v interface{}) {
	m[k.(uintptr)] = v.(int)
}
