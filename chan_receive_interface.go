package goo

var _ ChanReceive = ChanReceiveInterface(nil)

var _ Pointer = (*ChanReceiveInterface)(nil)

// ChanReceiveInterface is a receive channel of interface{}.
type ChanReceiveInterface chan interface{}

// Cap implements ChanReceive.
func (c ChanReceiveInterface) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveInterface) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveInterface) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveInterface) Make(cap int) Chan {
	return make(ChanInterface, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveInterface) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveInterface) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveInterface) Reference() Pointer {
	return &c
}
