package goo

var _ ChanReceive = ChanReceiveString(nil)

// ChanReceiveString is a receive channel of string.
type ChanReceiveString chan string

// Cap implements ChanReceive.
func (c ChanReceiveString) Cap() int {
	return cap(c)
}

// Len implements ChanReceive.
func (c ChanReceiveString) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceiveString) Make(cap int) Chan {
	return make(ChanString, cap)
}

// Receive implements ChanReceive.
func (c ChanReceiveString) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceiveString) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}
