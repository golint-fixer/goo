package goo

var _ Chan = ChanBool(nil)

// ChanBool is a channel of bool.
type ChanBool chan bool

// Cap implements Chan.
func (c ChanBool) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanBool) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanBool) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanBool) Close() {
	close(c)
}

// Dereference implements Chan.
func (c *ChanBool) Dereference() Value {
	return *c
}

// Len implements Chan.
func (c ChanBool) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanBool) Make(cap int) Chan {
	return make(ChanBool, cap)
}

// Receive implements Chan.
func (c ChanBool) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanBool) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements Chan.
func (c ChanBool) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanBool) Send(v interface{}) {
	c <- v.(bool)
}
