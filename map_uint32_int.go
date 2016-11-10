package goo

var _ Map = MapUint32Int(nil)

// MapUint32Int is a map from uint32 to int.
type MapUint32Int map[uint32]int

// Delete implements Map.
func (m MapUint32Int) Delete(k interface{}) {
	delete(m, k.(uint32))
}

// Dereference implements Map.
func (m *MapUint32Int) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUint32Int) Equals(other Equatable) bool {
	var n = other.(MapUint32Int)

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
func (m MapUint32Int) Get(k interface{}) interface{} {
	return m[k.(uint32)]
}

// GetCheck implements Map.
func (m MapUint32Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint32Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint32Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint32Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint32Int) Make(c int) Map {
	return make(MapUint32Int, c)
}

// NotEquals implements Map.
func (m MapUint32Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUint32Int) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUint32Int) Set(k, v interface{}) {
	m[k.(uint32)] = v.(int)
}
