package goo

var _ Chan = ChanRune(nil)

// ChanRune is a channel of rune.
type ChanRune chan rune

// Cap implements Chan.
func (c ChanRune) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanRune) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanRune) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanRune) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanRune) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanRune) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanRune) Make(cap int) Chan {
	return make(ChanRune, cap)
}

// Receive implements Chan.
func (c ChanRune) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanRune) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanRune) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanRune) Send(v interface{}) {
	c <- v.(rune)
}
