package goo

var _ ChanSend = ChanSendBool(nil)

// ChanSendBool is a send channel of bool.
type ChanSendBool chan bool

// Cap implements ChanSend.
func (c ChanSendBool) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendBool) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendBool) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendBool) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendBool) Make(cap int) Chan {
	return make(ChanBool, cap)
}

// Reference implements ChanSend.
func (c ChanSendBool) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendBool) Send(v interface{}) {
	c <- v.(bool)
}
