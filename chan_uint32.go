package goo

var _ Chan = ChanUint32(nil)

// ChanUint32 is a channel of uint32.
type ChanUint32 chan uint32

// Cap implements Chan.
func (c ChanUint32) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanUint32) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanUint32) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanUint32) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanUint32) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanUint32) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanUint32) Make(cap int) Chan {
	return make(ChanUint32, cap)
}

// Receive implements Chan.
func (c ChanUint32) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanUint32) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanUint32) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanUint32) Send(v interface{}) {
	c <- v.(uint32)
}
