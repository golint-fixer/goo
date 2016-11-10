package goo

var _ Chan = ChanUint64(nil)

// ChanUint64 is a channel of uint64.
type ChanUint64 chan uint64

// Cap implements Chan.
func (c ChanUint64) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanUint64) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanUint64) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanUint64) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanUint64) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanUint64) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanUint64) Make(cap int) Chan {
	return make(ChanUint64, cap)
}

// Receive implements Chan.
func (c ChanUint64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanUint64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanUint64) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanUint64) Send(v interface{}) {
	c <- v.(uint64)
}
