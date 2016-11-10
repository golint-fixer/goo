package goo

var _ Map = MapComplex64Int(nil)

var _ Pointer = &MapComplex64Int{}

// MapComplex64Int is a map from complex64 to int.
type MapComplex64Int map[complex64]int

// Delete implements Map.
func (m MapComplex64Int) Delete(k interface{}) {
	delete(m, k.(complex64))
}

// Dereference implements Map.
func (m *MapComplex64Int) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapComplex64Int) Equals(other Equatable) bool {
	var n = other.(MapComplex64Int)

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
func (m MapComplex64Int) Get(k interface{}) interface{} {
	return m[k.(complex64)]
}

// GetCheck implements Map.
func (m MapComplex64Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(complex64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapComplex64Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapComplex64Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapComplex64Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapComplex64Int) Make(c int) Map {
	return make(MapComplex64Int, c)
}

// NotEquals implements Map.
func (m MapComplex64Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapComplex64Int) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapComplex64Int) Set(k, v interface{}) {
	m[k.(complex64)] = v.(int)
}
