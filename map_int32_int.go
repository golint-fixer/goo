package goo

var _ Map = MapInt32Int(nil)

// MapInt32Int is a map from int32 to int.
type MapInt32Int map[int32]int

// Delete implements Map.
func (m MapInt32Int) Delete(k interface{}) {
	delete(m, k.(int32))
}

// Equals implements Map.
func (m MapInt32Int) Equals(other Equatable) bool {
	var n = other.(MapInt32Int)

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
func (m MapInt32Int) Get(k interface{}) interface{} {
	return m[k.(int32)]
}

// GetCheck implements Map.
func (m MapInt32Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt32Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt32Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt32Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt32Int) Make(c int) Map {
	return make(MapInt32Int, c)
}

// NotEquals implements Map.
func (m MapInt32Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapInt32Int) Set(k, v interface{}) {
	m[k.(int32)] = v.(int)
}
