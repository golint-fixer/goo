package goo

var _ ChanSend = ChanSendString(nil)

// ChanSendString is a send channel of string.
type ChanSendString chan string

// Cap implements ChanSend.
func (c ChanSendString) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendString) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendString) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendString) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendString) Make(cap int) Chan {
	return make(ChanString, cap)
}

// Reference implements ChanSend.
func (c ChanSendString) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendString) Send(v interface{}) {
	c <- v.(string)
}
