package goo

var _ ChanSend = ChanSendUintptr(nil)

// ChanSendUintptr is a send channel of uintptr.
type ChanSendUintptr chan uintptr

// Cap implements ChanSend.
func (c ChanSendUintptr) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendUintptr) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendUintptr) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUintptr) Make(cap int) Chan {
	return make(ChanUintptr, cap)
}

// Send implements Chan.
func (c ChanSendUintptr) Send(v interface{}) {
	c <- v.(uintptr)
}
