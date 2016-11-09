package goo

var _ Chan = ChanByte(nil)

// ChanByte is a channel of byte.
type ChanByte chan byte

// Cap implements Chan.
func (c ChanByte) Cap() int {
	return cap(c)
}

// ChanReceive implements Chan.
func (c ChanByte) ChanReceive() ChanReceive {
	return c
}

// ChanSend implements Chan.
func (c ChanByte) ChanSend() ChanSend {
	return c
}

// Close implements Chan.
func (c ChanByte) Close() {
	close(c)
}

// Len implements Chan.
func (c ChanByte) Len() int {
	return len(c)
}

// Make implements Chan.
func (c ChanByte) Make(cap int) Chan {
	return make(ChanByte, cap)
}

// Receive implements Chan.
func (c ChanByte) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements Chan.
func (c ChanByte) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Send implements Chan.
func (c ChanByte) Send(v interface{}) {
	c <- v.(byte)
}
