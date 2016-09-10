package maps

type MapInterfaceBool map[interface{}]bool

func (m MapInterfaceBool) Delete(k interface{}) {
	delete(m, k)
}

func (m MapInterfaceBool) Get(k interface{}) interface{} {
	return m[k]
}

func (m MapInterfaceBool) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k]

	return v, ok
}

func (m MapInterfaceBool) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

func (m MapInterfaceBool) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

func (m MapInterfaceBool) Len() int {
	return len(m)
}

func (m MapInterfaceBool) Make(c int) Map {
	return make(MapInterfaceBool, c)
}

func (m MapInterfaceBool) Set(k, v interface{}) {
	m[k] = v.(bool)
}

type MapInterfaceInt map[interface{}]int

func (m MapInterfaceInt) Delete(k interface{}) {
	delete(m, k)
}

func (m MapInterfaceInt) Get(k interface{}) interface{} {
	return m[k]
}

func (m MapInterfaceInt) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k]

	return v, ok
}

func (m MapInterfaceInt) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

func (m MapInterfaceInt) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

func (m MapInterfaceInt) Len() int {
	return len(m)
}

func (m MapInterfaceInt) Make(c int) Map {
	return make(MapInterfaceInt, c)
}

func (m MapInterfaceInt) Set(k, v interface{}) {
	m[k] = v.(int)
}

type MapInterfaceInterface map[interface{}]interface{}

func (m MapInterfaceInterface) Delete(k interface{}) {
	delete(m, k)
}

func (m MapInterfaceInterface) Get(k interface{}) interface{} {
	return m[k]
}

func (m MapInterfaceInterface) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k]

	return v, ok
}

func (m MapInterfaceInterface) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

func (m MapInterfaceInterface) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

func (m MapInterfaceInterface) Len() int {
	return len(m)
}

func (m MapInterfaceInterface) Make(c int) Map {
	return make(MapInterfaceInterface, c)
}

func (m MapInterfaceInterface) Set(k, v interface{}) {
	m[k] = v.(int)
}

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
