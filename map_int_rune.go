package goo

var _ Map = MapIntRune(nil)

// MapIntRune is a map from int to rune.
type MapIntRune map[int]rune

// Delete implements Map.
func (m MapIntRune) Delete(k interface{}) {
	delete(m, k.(int))
}

// Dereference implements Map.
func (m *MapIntRune) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapIntRune) Equals(other Equatable) bool {
	var n = other.(MapIntRune)

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
func (m MapIntRune) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntRune) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntRune) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntRune) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntRune) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntRune) Make(c int) Map {
	return make(MapIntRune, c)
}

// NotEquals implements Map.
func (m MapIntRune) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapIntRune) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapIntRune) Set(k, v interface{}) {
	m[k.(int)] = v.(rune)
}
