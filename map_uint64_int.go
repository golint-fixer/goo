package goo

var _ Map = MapUint64Int(nil)

// MapUint64Int is a map from uint64 to int.
type MapUint64Int map[uint64]int

// Delete implements Map.
func (m MapUint64Int) Delete(k interface{}) {
	delete(m, k.(uint64))
}

// Dereference implements Map.
func (m *MapUint64Int) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUint64Int) Equals(other Equatable) bool {
	var n = other.(MapUint64Int)

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
func (m MapUint64Int) Get(k interface{}) interface{} {
	return m[k.(uint64)]
}

// GetCheck implements Map.
func (m MapUint64Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint64Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint64Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint64Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint64Int) Make(c int) Map {
	return make(MapUint64Int, c)
}

// NotEquals implements Map.
func (m MapUint64Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUint64Int) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUint64Int) Set(k, v interface{}) {
	m[k.(uint64)] = v.(int)
}
