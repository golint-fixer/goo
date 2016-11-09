package goo

var _ ChanReceive = ChanReceiveInterface(nil)

// ChanReceiveInterface is a receive channel of interface{}.
type ChanReceiveInterface chan interface{}

// Cap implements ChanReceive.
func (c ChanReceiveInterface) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveInterface) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveInterface) Make(cap int) Chan {
	return make(ChanInterface, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveInterface) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveInterface) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
