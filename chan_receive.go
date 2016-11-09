package goo

// ChanReceive is a receive channel.
type ChanReceive interface {
	// Cap returns the capacity.
	Cap() int

	// Len returns the length.
	Len() int

	// Make returns a new chan with capacity c.
	Make(c int) Chan

	// Receive returns the next element.
	Receive() interface{}

	// ReceiveCheck returns the next element and whether there is a next element.
	ReceiveCheck() (interface{}, bool)
}

// TODO: Cycle, Iterate, Repeat

func ChanEqual(c, d ChanReceive) bool {
	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc != vd {
			return false
		}
	}

	return true
}

func ChanEqualDeep(c, d ChanReceive) bool {
	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc, ok := vc.(Equatable); ok {
			if vd, ok := vd.(Equatable); ok {
				if vc.NotEquals(vd) {
					return false
				}
			}
		}
	}

	return true
}

func ChanIterator(c ChanReceive) Iterator {
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
