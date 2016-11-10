package goo

var _ Chan = ChanInt32(nil)

// ChanInt32 is a channel of int32.
type ChanInt32 chan int32

// Cap implements Chan.
func (c ChanInt32) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanInt32) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanInt32) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanInt32) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanInt32) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanInt32) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanInt32) Make(cap int) Chan {
	return make(ChanInt32, cap)
}

// Receive implements Chan.
func (c ChanInt32) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanInt32) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanInt32) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanInt32) Send(v interface{}) {
	c <- v.(int32)
}
