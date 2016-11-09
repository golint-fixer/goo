package goo

var _ Map = MapComplex128Int(nil)

// MapComplex128Int is a map from complex128 to int.
type MapComplex128Int map[complex128]int

// Delete implements Map.
func (m MapComplex128Int) Delete(k interface{}) {
	delete(m, k.(complex128))
}

// Equals implements Map.
func (m MapComplex128Int) Equals(other Equatable) bool {
	var n = other.(MapComplex128Int)

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
func (m MapComplex128Int) Get(k interface{}) interface{} {
	return m[k.(complex128)]
}

// GetCheck implements Map.
func (m MapComplex128Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(complex128)]

	return v, ok
}

// KeyValues implements Map.
func (m MapComplex128Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapComplex128Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapComplex128Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapComplex128Int) Make(c int) Map {
	return make(MapComplex128Int, c)
}

// NotEquals implements Map.
func (m MapComplex128Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapComplex128Int) Set(k, v interface{}) {
	m[k.(complex128)] = v.(int)
}
