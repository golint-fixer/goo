package goo

var _ Map = MapInt64String(nil)

// MapInt64String is a map from int64 to string.
type MapInt64String map[int64]string

// Delete implements Map.
func (m MapInt64String) Delete(k interface{}) {
	delete(m, k.(int64))
}

// Get implements Map.
func (m MapInt64String) Get(k interface{}) interface{} {
	return m[k.(int64)]
}

// GetCheck implements Map.
func (m MapInt64String) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt64String) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt64String) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt64String) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt64String) Make(c int) Map {
	return make(MapInt64String, c)
}

// Set implements Map.
func (m MapInt64String) Set(k, v interface{}) {
	m[k.(int64)] = v.(string)
}
