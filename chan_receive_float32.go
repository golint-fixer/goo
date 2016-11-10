package goo

var _ ChanReceive = ChanReceiveFloat32(nil)

var _ Pointer = (*ChanReceiveFloat32)(nil)

// ChanReceiveFloat32 is a receive channel of float32.
type ChanReceiveFloat32 chan float32

// Cap implements ChanReceive.
func (c ChanReceiveFloat32) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveFloat32) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveFloat32) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveFloat32) Make(cap int) Chan {
	return make(ChanFloat32, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveFloat32) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveFloat32) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveFloat32) Reference() Pointer {
	return &c
}
