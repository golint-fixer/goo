package goo

var _ Chan = ChanUint8(nil)

// ChanUint8 is a channel of uint8.
type ChanUint8 chan uint8

// Cap implements Chan.
func (c ChanUint8) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanUint8) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanUint8) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanUint8) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanUint8) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanUint8) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanUint8) Make(cap int) Chan {
	return make(ChanUint8, cap)
}

// Receive implements Chan.
func (c ChanUint8) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanUint8) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanUint8) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanUint8) Send(v interface{}) {
	c <- v.(uint8)
}
