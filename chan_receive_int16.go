package goo

var _ ChanReceive = ChanReceiveInt16(nil)

// ChanReceiveInt16 is a receive channel of int16.
type ChanReceiveInt16 chan int16

// Cap implements ChanReceive.
func (c ChanReceiveInt16) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveInt16) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveInt16) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveInt16) Make(cap int) Chan {
	return make(ChanInt16, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveInt16) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveInt16) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveInt16) Reference() Pointer {
	return &c
}
