package goo

var _ ChanReceive = ChanReceiveInt(nil)

// ChanReceiveInt is a receive channel of int.
type ChanReceiveInt chan int

// Cap implements ChanReceive.
func (c ChanReceiveInt) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveInt) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveInt) Make(cap int) Chan {
	return make(ChanInt, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveInt) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveInt) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
