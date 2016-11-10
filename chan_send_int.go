package goo

var _ ChanSend = ChanSendInt(nil)

var _ Pointer = (*ChanSendInt)(nil)

// ChanSendInt is a send channel of int.
type ChanSendInt chan int

// Cap implements ChanSend.
func (c ChanSendInt) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendInt) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendInt) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendInt) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInt) Make(cap int) Chan {
	return make(ChanInt, cap)
}

// Reference implements ChanSend.
func (c ChanSendInt) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendInt) Send(v interface{}) {
	c <- v.(int)
}
