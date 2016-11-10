package goo

var _ Chan = ChanInt(nil)

// ChanInt is a channel of int.
type ChanInt chan int

// Cap implements Chan.
func (c ChanInt) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanInt) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanInt) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanInt) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanInt) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanInt) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanInt) Make(cap int) Chan {
	return make(ChanInt, cap)
}

// Receive implements Chan.
func (c ChanInt) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanInt) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanInt) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanInt) Send(v interface{}) {
	c <- v.(int)
}
