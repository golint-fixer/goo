package goo

var _ ChanSend = ChanSendInt32(nil)

// ChanSendInt32 is a send channel of int32.
type ChanSendInt32 chan int32

// Cap implements ChanSend.
func (c ChanSendInt32) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendInt32) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendInt32) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInt32) Make(cap int) Chan {
	return make(ChanInt32, cap)
}

// Send implements Chan.
func (c ChanSendInt32) Send(v interface{}) {
	c <- v.(int32)
}
