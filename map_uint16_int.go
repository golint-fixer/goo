package goo

var _ Map = MapUint16Int(nil)

// MapUint16Int is a map from uint16 to int.
type MapUint16Int map[uint16]int

// Delete implements Map.
func (m MapUint16Int) Delete(k interface{}) {
	delete(m, k.(uint16))
}

// Get implements Map.
func (m MapUint16Int) Get(k interface{}) interface{} {
	return m[k.(uint16)]
}

// GetCheck implements Map.
func (m MapUint16Int) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(uint16)]

	return v, ok
}

// KeyValues implements Map.
func (m MapUint16Int) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m MapUint16Int) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m MapUint16Int) Len() int {
	return len(m)
}

// Make implements Map.
func (m MapUint16Int) Make(c int) Map {
	return make(MapUint16Int, c)
}

// Set implements Map.
func (m MapUint16Int) Set(k, v interface{}) {
	m[k.(uint16)] = v.(int)
}
