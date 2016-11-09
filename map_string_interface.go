package goo

var _ Map = MapStringInterface(nil)

// MapStringInterface is a map from string to interface{}.
type MapStringInterface map[string]interface{}

// Delete implements Map.
func (m MapStringInterface) Delete(k interface{}) {
	delete(m, k.(string))
}

// Get implements Map.
func (m MapStringInterface) Get(k interface{}) interface{} {
	return m[k.(string)]
}

// GetCheck implements Map.
func (m MapStringInterface) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(string)]

	return v, ok
}

// KeyValues implements Map.
func (m MapStringInterface) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapStringInterface) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapStringInterface) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapStringInterface) Make(c int) Map {
	return make(MapStringInterface, c)
}

// Set implements Map.
func (m MapStringInterface) Set(k, v interface{}) {
	m[k.(string)] = v.(interface{})
}
