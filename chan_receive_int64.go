package goo

var _ ChanReceive = ChanReceiveInt64(nil)

// ChanReceiveInt64 is a receive channel of int64.
type ChanReceiveInt64 chan int64

// Cap implements ChanReceive.
func (c ChanReceiveInt64) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveInt64) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveInt64) Make(cap int) Chan {
	return make(ChanInt64, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveInt64) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveInt64) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
