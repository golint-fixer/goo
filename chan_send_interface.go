package goo

var _ ChanSend = ChanSendInterface(nil)

var _ Pointer = (*ChanSendInterface)(nil)

// ChanSendInterface is a send channel of interface{}.
type ChanSendInterface chan interface{}

// Cap implements ChanSend.
func (c ChanSendInterface) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendInterface) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendInterface) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendInterface) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInterface) Make(cap int) Chan {
	return make(ChanInterface, cap)
}

// Reference implements ChanSend.
func (c ChanSendInterface) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendInterface) Send(v interface{}) {
	c <- v.(interface{})
}
