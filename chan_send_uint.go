package goo

var _ ChanSend = ChanSendUint(nil)

// ChanSendUint is a send channel of uint.
type ChanSendUint chan uint

// Cap implements ChanSend.
func (c ChanSendUint) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendUint) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendUint) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUint) Make(cap int) Chan {
	return make(ChanUint, cap)
}

// Send implements Chan.
func (c ChanSendUint) Send(v interface{}) {
	c <- v.(uint)
}
