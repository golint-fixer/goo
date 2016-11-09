package goo

var _ ChanSend = ChanSendString(nil)

// ChanSendString is a send channel of string.
type ChanSendString chan string

// Cap implements ChanSend.
func (c ChanSendString) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendString) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendString) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendString) Make(cap int) Chan {
	return make(ChanString, cap)
}

// Send implements Chan.
func (c ChanSendString) Send(v interface{}) {
	c <- v.(string)
}
