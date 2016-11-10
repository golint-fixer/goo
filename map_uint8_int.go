package goo

var _ Map = MapUint8Int(nil)

// MapUint8Int is a map from uint8 to int.
type MapUint8Int map[uint8]int

// Delete implements Map.
func (m MapUint8Int) Delete(k interface{}) {
	delete(m, k.(uint8))
}

// Dereference implements Map.
func (m *MapUint8Int) Dereference() Value {
	return *m
}

// Equals implements Map.
func (m MapUint8Int) Equals(other Equatable) bool {
	var n = other.(MapUint8Int)

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
func (m MapUint8Int) Get(k interface{}) interface{} {
	return m[k.(uint8)]
}

// GetCheck implements Map.
func (m MapUint8Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint8)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint8Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint8Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint8Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint8Int) Make(c int) Map {
	return make(MapUint8Int, c)
}

// NotEquals implements Map.
func (m MapUint8Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Reference implements Map.
func (m MapUint8Int) Reference() Pointer {
	return &m
}

// Set implements Map.
func (m MapUint8Int) Set(k, v interface{}) {
	m[k.(uint8)] = v.(int)
}
