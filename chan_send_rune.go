package goo

var _ ChanSend = ChanSendRune(nil)

var _ Pointer = (*ChanSendRune)(nil)

// ChanSendRune is a send channel of rune.
type ChanSendRune chan rune

// Cap implements ChanSend.
func (c ChanSendRune) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendRune) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendRune) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendRune) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendRune) Make(cap int) Chan {
	return make(ChanRune, cap)
}

// Reference implements ChanSend.
func (c ChanSendRune) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendRune) Send(v interface{}) {
	c <- v.(rune)
}
