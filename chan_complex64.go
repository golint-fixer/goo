package goo

var _ Chan = ChanComplex64(nil)

// ChanComplex64 is a channel of complex64.
type ChanComplex64 chan complex64

// Cap implements Chan.
func (c ChanComplex64) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanComplex64) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanComplex64) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanComplex64) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanComplex64) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanComplex64) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanComplex64) Make(cap int) Chan {
	return make(ChanComplex64, cap)
}

// Receive implements Chan.
func (c ChanComplex64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanComplex64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanComplex64) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanComplex64) Send(v interface{}) {
	c <- v.(complex64)
}
