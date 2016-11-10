package goo

var _ Chan = ChanUintptr(nil)

// ChanUintptr is a channel of uintptr.
type ChanUintptr chan uintptr

// Cap implements Chan.
func (c ChanUintptr) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanUintptr) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanUintptr) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanUintptr) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanUintptr) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanUintptr) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanUintptr) Make(cap int) Chan {
	return make(ChanUintptr, cap)
}

// Receive implements Chan.
func (c ChanUintptr) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanUintptr) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanUintptr) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanUintptr) Send(v interface{}) {
	c <- v.(uintptr)
}
