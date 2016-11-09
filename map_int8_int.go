package goo

var _ Map = MapInt8Int(nil)

// MapInt8Int is a map from int8 to int.
type MapInt8Int map[int8]int

// Delete implements Map.
func (m MapInt8Int) Delete(k interface{}) {
	delete(m, k.(int8))
}

// Equals implements Map.
func (m MapInt8Int) Equals(other Equatable) bool {
	var n = other.(MapInt8Int)

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
func (m MapInt8Int) Get(k interface{}) interface{} {
	return m[k.(int8)]
}

// GetCheck implements Map.
func (m MapInt8Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(int8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapInt8Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapInt8Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapInt8Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapInt8Int) Make(c int) Map {
	return make(MapInt8Int, c)
}

// NotEquals implements Map.
func (m MapInt8Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapInt8Int) Set(k, v interface{}) {
	m[k.(int8)] = v.(int)
}
