package goo

var _ Map = MapInterfaceBool(nil)

// MapInterfaceBool is a map from interface{} to bool.
type MapInterfaceBool map[interface{}]bool

// Delete implements Map.
func (m MapInterfaceBool) Delete(k interface{}) {
	delete(m, k.(interface{}))
}

// Get implements Map.
func (m MapInterfaceBool) Get(k interface{}) interface{} {
	return m[k.(interface{})]
}

// GetCheck implements Map.
func (m MapInterfaceBool) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(interface{})]

	return v, ok
}

// KeyValues implements Map.
func (m MapInterfaceBool) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInterfaceBool) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInterfaceBool) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInterfaceBool) Make(c int) Map {
	return make(MapInterfaceBool, c)
}

// Set implements Map.
func (m MapInterfaceBool) Set(k, v interface{}) {
	m[k.(interface{})] = v.(bool)
}
