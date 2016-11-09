package goo

var _ ChanReceive = ChanReceiveUintptr(nil)

// ChanReceiveUintptr is a receive channel of uintptr.
type ChanReceiveUintptr chan uintptr

// Cap implements ChanReceive.
func (c ChanReceiveUintptr) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveUintptr) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveUintptr) Make(cap int) Chan {
	return make(ChanUintptr, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveUintptr) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveUintptr) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
