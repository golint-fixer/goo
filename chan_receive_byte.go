package goo

var _ ChanReceive = ChanReceiveByte(nil)

// ChanReceiveByte is a receive channel of byte.
type ChanReceiveByte chan byte

// Cap implements ChanReceive.
func (c ChanReceiveByte) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveByte) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveByte) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveByte) Make(cap int) Chan {
	return make(ChanByte, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveByte) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveByte) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveByte) Reference() Pointer {
	return &c
}
