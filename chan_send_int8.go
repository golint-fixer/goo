package goo

var _ ChanSend = ChanSendInt8(nil)

// ChanSendInt8 is a send channel of int8.
type ChanSendInt8 chan int8

// Cap implements ChanSend.
func (c ChanSendInt8) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendInt8) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendInt8) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendInt8) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInt8) Make(cap int) Chan {
	return make(ChanInt8, cap)
}

// Reference implements ChanSend.
func (c ChanSendInt8) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendInt8) Send(v interface{}) {
	c <- v.(int8)
}
