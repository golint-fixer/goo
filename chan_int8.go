package goo

var _ Chan = ChanInt8(nil)

// ChanInt8 is a channel of int8.
type ChanInt8 chan int8

// Cap implements Chan.
func (c ChanInt8) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanInt8) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanInt8) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanInt8) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanInt8) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanInt8) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanInt8) Make(cap int) Chan {
	return make(ChanInt8, cap)
}

// Receive implements Chan.
func (c ChanInt8) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanInt8) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanInt8) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanInt8) Send(v interface{}) {
	c <- v.(int8)
}
