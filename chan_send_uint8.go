package goo

var _ ChanSend = ChanSendUint8(nil)

// ChanSendUint8 is a send channel of uint8.
type ChanSendUint8 chan uint8

// Cap implements ChanSend.
func (c ChanSendUint8) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendUint8) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendUint8) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUint8) Make(cap int) Chan {
	return make(ChanUint8, cap)
}

// Send implements Chan.
func (c ChanSendUint8) Send(v interface{}) {
	c <- v.(uint8)
}
