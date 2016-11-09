package goo

var _ Map = MapIntInt64(nil)

// MapIntInt64 is a map from int to int64.
type MapIntInt64 map[int]int64

// Delete implements Map.
func (m MapIntInt64) Delete(k interface{}) {
	delete(m, k.(int))
}

// Equals implements Map.
func (m MapIntInt64) Equals(other Equatable) bool {
	var n = other.(MapIntInt64)

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
func (m MapIntInt64) Get(k interface{}) interface{} {
	return m[k.(int)]
}

// GetCheck implements Map.
func (m MapIntInt64) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int)]

	return v, ok
}

// KeyValues implements Map.
func (m MapIntInt64) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapIntInt64) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapIntInt64) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapIntInt64) Make(c int) Map {
	return make(MapIntInt64, c)
}

// NotEquals implements Map.
func (m MapIntInt64) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapIntInt64) Set(k, v interface{}) {
	m[k.(int)] = v.(int64)
}
