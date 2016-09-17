package goo

type Map interface {
	Delete(k interface{})

	Get(k interface{}) interface{}

	GetCheck(k interface{}) (interface{}, bool)

	KeyValues() [][2]interface{}

	Keys() []interface{}

	Len() int

	Make(c int) Map

	Set(k, v interface{})
}

func KeyValuesIterator(m Map) Iterator {
	var kv = m.KeyValues()

	return &keyValuesIterator{n: len(kv), s: kv}
}

func KeysIterator(m Map) Iterator {
	var k = m.Keys()

	return &keyIterator{n: len(k), s: k}
}

type keyIterator struct {
	i int

	n int

	s []interface{}
}

func (i keyIterator) More() bool {
	return i.i < i.n
}

func (i *keyIterator) Next() interface{} {
	var v = i.s[i.i]

	i.i++

	return v
}

type keyValuesIterator struct {
	i int

	n int

	s [][2]interface{}
}

func (i keyValuesIterator) More() bool {
	return i.i < i.n
}

func (i *keyValuesIterator) Next() interface{} {
	var v = i.s[i.i]

	i.i++

	return v
}
