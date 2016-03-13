package chans

import "github.com/willfaught/lang"

type ChanReceive interface {
	Cap() int

	Len() int

	Make(c int) Chan

	Receive() interface{}

	ReceiveCheck() (interface{}, bool)
}

func Iterator(c ChanReceive) lang.Iterator {
	return &iterator{c: c}
}

type iterator struct {
	c ChanReceive

	checked bool

	ok bool

	v interface{}
}

func (i *iterator) More() bool {
	if i.checked {
		return i.ok
	}

	i.v, i.ok = i.c.ReceiveCheck()
	i.checked = true

	return i.ok
}

func (i *iterator) Next() interface{} {
	if !i.checked {
		i.More()
	}

	i.checked = false

	return i.v
}
