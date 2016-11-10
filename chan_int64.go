package goo

var _ Chan = ChanInt64(nil)

var _ Pointer = (*ChanInt64)(nil)

// ChanInt64 is a channel of int64.
type ChanInt64 chan int64

// Cap implements Chan.
func (c ChanInt64) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanInt64) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanInt64) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanInt64) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanInt64) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanInt64) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanInt64) Make(cap int) Chan {
	return make(ChanInt64, cap)
}

// Receive implements Chan.
func (c ChanInt64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanInt64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanInt64) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanInt64) Send(v interface{}) {
	c <- v.(int64)
}
