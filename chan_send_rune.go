package goo

var _ ChanSend = ChanSendRune(nil)

// ChanSendRune is a send channel of rune.
type ChanSendRune chan rune

// Cap implements ChanSend.
func (c ChanSendRune) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendRune) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendRune) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendRune) Make(cap int) Chan {
	return make(ChanRune, cap)
}

// Send implements Chan.
func (c ChanSendRune) Send(v interface{}) {
	c <- v.(rune)
}
