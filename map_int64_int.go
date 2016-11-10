package goo

var _ Map = MapInt64Int(nil)

// MapInt64Int is a map from int64 to int.
type MapInt64Int map[int64]int

// Delete implements Map.
func (m MapInt64Int) Delete(k interface{}) {
	delete(m, k.(int64))
}

// Dereference implements Map.
func (m *MapInt64Int) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapInt64Int) Equals(other Equatable) bool {
	var n = other.(MapInt64Int)

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
func (m MapInt64Int) Get(k interface{}) interface{} {
	return m[k.(int64)]
}

// GetCheck implements Map.
func (m MapInt64Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int64)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt64Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt64Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt64Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt64Int) Make(c int) Map {
	return make(MapInt64Int, c)
}

// NotEquals implements Map.
func (m MapInt64Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapInt64Int) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapInt64Int) Set(k, v interface{}) {
	m[k.(int64)] = v.(int)
}
