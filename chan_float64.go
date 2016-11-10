package goo

var _ Chan = ChanFloat64(nil)

// ChanFloat64 is a channel of float64.
type ChanFloat64 chan float64

// Cap implements Chan.
func (c ChanFloat64) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanFloat64) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanFloat64) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanFloat64) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanFloat64) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanFloat64) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanFloat64) Make(cap int) Chan {
	return make(ChanFloat64, cap)
}

// Receive implements Chan.
func (c ChanFloat64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanFloat64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanFloat64) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanFloat64) Send(v interface{}) {
	c <- v.(float64)
}
