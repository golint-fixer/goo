package goo

var _ Chan = ChanInt16(nil)

// ChanInt16 is a channel of int16.
type ChanInt16 chan int16

// Cap implements Chan.
func (c ChanInt16) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanInt16) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanInt16) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanInt16) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanInt16) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanInt16) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanInt16) Make(cap int) Chan {
	return make(ChanInt16, cap)
}

// Receive implements Chan.
func (c ChanInt16) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanInt16) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanInt16) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanInt16) Send(v interface{}) {
	c <- v.(int16)
}
