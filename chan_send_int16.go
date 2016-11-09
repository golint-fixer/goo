package goo

var _ ChanSend = ChanSendInt16(nil)

// ChanSendInt16 is a send channel of int16.
type ChanSendInt16 chan int16

// Cap implements ChanSend.
func (c ChanSendInt16) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendInt16) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendInt16) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInt16) Make(cap int) Chan {
	return make(ChanInt16, cap)
}

// Send implements Chan.
func (c ChanSendInt16) Send(v interface{}) {
	c <- v.(int16)
}
