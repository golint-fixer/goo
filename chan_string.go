package goo

var _ Chan = ChanString(nil)

// ChanString is a channel of string.
type ChanString chan string

// Cap implements Chan.
func (c ChanString) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanString) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanString) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanString) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanString) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanString) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanString) Make(cap int) Chan {
	return make(ChanString, cap)
}

// Receive implements Chan.
func (c ChanString) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanString) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanString) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanString) Send(v interface{}) {
	c <- v.(string)
}
