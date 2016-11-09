package goo

var _ ChanSend = ChanSendInterface(nil)

// ChanSendInterface is a send channel of interface{}.
type ChanSendInterface chan interface{}

// Cap implements ChanSend.
func (c ChanSendInterface) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendInterface) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendInterface) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInterface) Make(cap int) Chan {
	return make(ChanInterface, cap)
}

// Send implements Chan.
func (c ChanSendInterface) Send(v interface{}) {
	c <- v.(interface{})
}
