package goo

var _ ChanReceive = ChanReceiveFloat64(nil)

// ChanReceiveFloat64 is a receive channel of float64.
type ChanReceiveFloat64 chan float64

// Cap implements ChanReceive.
func (c ChanReceiveFloat64) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceiveFloat64) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceiveFloat64) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveFloat64) Make(cap int) Chan {
	return make(ChanFloat64, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveFloat64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveFloat64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceiveFloat64) Reference() Pointer {
	return &c
}
