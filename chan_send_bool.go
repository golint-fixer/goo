package goo

var _ ChanSend = ChanSendBool(nil)

// ChanSendBool is a send channel of bool.
type ChanSendBool chan bool

// Cap implements ChanSend.
func (c ChanSendBool) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendBool) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendBool) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendBool) Make(cap int) Chan {
	return make(ChanBool, cap)
}

// Send implements Chan.
func (c ChanSendBool) Send(v interface{}) {
	c <- v.(bool)
}
