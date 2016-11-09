package goo

var _ ChanReceive = ChanReceiveInt8(nil)

// ChanReceiveInt8 is a receive channel of int8.
type ChanReceiveInt8 chan int8

// Cap implements ChanReceive.
func (c ChanReceiveInt8) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveInt8) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveInt8) Make(cap int) Chan {
	return make(ChanInt8, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveInt8) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveInt8) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
