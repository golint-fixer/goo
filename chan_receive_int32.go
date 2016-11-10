package goo

var _ ChanReceive = ChanReceiveInt32(nil)

// ChanReceiveInt32 is a receive channel of int32.
type ChanReceiveInt32 chan int32

// Cap implements ChanReceive.
func (c ChanReceiveInt32) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveInt32) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveInt32) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveInt32) Make(cap int) Chan {
	return make(ChanInt32, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveInt32) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveInt32) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveInt32) Reference() Pointer {
	return &c
}
