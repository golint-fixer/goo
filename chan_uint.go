package goo

var _ Chan = ChanUint(nil)

var _ Pointer = (*ChanUint)(nil)

// ChanUint is a channel of uint.
type ChanUint chan uint

// Cap implements Chan.
func (c ChanUint) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanUint) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanUint) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanUint) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanUint) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanUint) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanUint) Make(cap int) Chan {
	return make(ChanUint, cap)
}

// Receive implements Chan.
func (c ChanUint) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanUint) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanUint) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanUint) Send(v interface{}) {
	c <- v.(uint)
}
