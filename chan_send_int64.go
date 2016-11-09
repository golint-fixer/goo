package goo

var _ ChanSend = ChanSendInt64(nil)

// ChanSendInt64 is a send channel of int64.
type ChanSendInt64 chan int64

// Cap implements ChanSend.
func (c ChanSendInt64) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendInt64) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendInt64) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendInt64) Make(cap int) Chan {
	return make(ChanInt64, cap)
}

// Send implements Chan.
func (c ChanSendInt64) Send(v interface{}) {
	c <- v.(int64)
}
