package goo

var _ Map = MapInterfaceInterface(nil)

// MapInterfaceInterface is a map from interface{} to interface{}.
type MapInterfaceInterface map[interface{}]interface{}

// Delete implements Map.
func (m MapInterfaceInterface) Delete(k interface{}) {
	delete(m, k.(interface{}))
}

// Dereference implements Map.
func (m *MapInterfaceInterface) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapInterfaceInterface) Equals(other Equatable) bool {
	var n = other.(MapInterfaceInterface)

	if len(n) != len(m) {
		return false
	}

	for k, v := range m {
		if nv, ok := n[k]; !ok {
			return false
		} else if nv != v {
			return false
		}
	}

	return true
}

// Get implements Map.
func (m MapInterfaceInterface) Get(k interface{}) interface{} {
	return m[k.(interface{})]
}

// GetCheck implements Map.
func (m MapInterfaceInterface) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(interface{})]

	return v, ok
}

// KeyValues implements Map.
func (m MapInterfaceInterface) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInterfaceInterface) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInterfaceInterface) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInterfaceInterface) Make(c int) Map {
	return make(MapInterfaceInterface, c)
}

// NotEquals implements Map.
func (m MapInterfaceInterface) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapInterfaceInterface) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapInterfaceInterface) Set(k, v interface{}) {
	m[k.(interface{})] = v.(interface{})
}
