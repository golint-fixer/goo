package goo

var _ ChanSend = ChanSendUint64(nil)

// ChanSendUint64 is a send channel of uint64.
type ChanSendUint64 chan uint64

// Cap implements ChanSend.
func (c ChanSendUint64) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendUint64) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendUint64) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUint64) Make(cap int) Chan {
	return make(ChanUint64, cap)
}

// Send implements Chan.
func (c ChanSendUint64) Send(v interface{}) {
	c <- v.(uint64)
}
