package maps

type MapInterfaceStruct map[interface{}]struct{}

func (m MapInterfaceStruct) Delete(k interface{}) {
	delete(m, k)
}

func (m MapInterfaceStruct) Get(k interface{}) interface{} {
	return m[k]
}

func (m MapInterfaceStruct) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k]

	return v, ok
}

func (m MapInterfaceStruct) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

func (m MapInterfaceStruct) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

func (m MapInterfaceStruct) Len() int {
	return len(m)
}

func (m MapInterfaceStruct) Make(c int) Map {
	return make(MapInterfaceStruct, c)
}

func (m MapInterfaceStruct) Set(k, v interface{}) {
	m[k] = v.(struct{})
}
