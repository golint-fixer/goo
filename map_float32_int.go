package goo

var _ Map = MapFloat32Int(nil)

// MapFloat32Int is a map from float32 to int.
type MapFloat32Int map[float32]int

// Delete implements Map.
func (m MapFloat32Int) Delete(k interface{}) {
	delete(m, k.(float32))
}

// Equals implements Map.
func (m MapFloat32Int) Equals(other Equatable) bool {
	var n = other.(MapFloat32Int)

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
func (m MapFloat32Int) Get(k interface{}) interface{} {
	return m[k.(float32)]
}

// GetCheck implements Map.
func (m MapFloat32Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(float32)]

	return v, ok
}

// KeyValues implements Map.
func (m MapFloat32Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapFloat32Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapFloat32Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapFloat32Int) Make(c int) Map {
	return make(MapFloat32Int, c)
}

// NotEquals implements Map.
func (m MapFloat32Int) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m MapFloat32Int) Set(k, v interface{}) {
	m[k.(float32)] = v.(int)
}
