package goo

var _ ChanSend = ChanSendInt(nil)

// ChanSendInt is a send channel of int.
type ChanSendInt chan int

// Cap implements ChanSend.
func (c ChanSendInt) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendInt) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendInt) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInt) Make(cap int) Chan {
	return make(ChanInt, cap)
}

// Send implements Chan.
func (c ChanSendInt) Send(v interface{}) {
	c <- v.(int)
}
