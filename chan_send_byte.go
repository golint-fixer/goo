package goo

var _ ChanSend = ChanSendByte(nil)

var _ Pointer = (*ChanSendByte)(nil)

// ChanSendByte is a send channel of byte.
type ChanSendByte chan byte

// Cap implements ChanSend.
func (c ChanSendByte) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendByte) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendByte) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendByte) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendByte) Make(cap int) Chan {
	return make(ChanByte, cap)
}

// Reference implements ChanSend.
func (c ChanSendByte) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendByte) Send(v interface{}) {
	c <- v.(byte)
}
