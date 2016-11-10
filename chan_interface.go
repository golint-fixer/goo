package goo

var _ Chan = ChanInterface(nil)

// ChanInterface is a channel of interface{}.
type ChanInterface chan interface{}

// Cap implements Chan.
func (c ChanInterface) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanInterface) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanInterface) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanInterface) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanInterface) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanInterface) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanInterface) Make(cap int) Chan {
	return make(ChanInterface, cap)
}

// Receive implements Chan.
func (c ChanInterface) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanInterface) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanInterface) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanInterface) Send(v interface{}) {
	c <- v.(interface{})
}
