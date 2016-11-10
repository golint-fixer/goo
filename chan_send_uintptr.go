package goo

var _ ChanSend = ChanSendUintptr(nil)

var _ Pointer = (*ChanSendUintptr)(nil)

// ChanSendUintptr is a send channel of uintptr.
type ChanSendUintptr chan uintptr

// Cap implements ChanSend.
func (c ChanSendUintptr) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendUintptr) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendUintptr) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendUintptr) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendUintptr) Make(cap int) Chan {
	return make(ChanUintptr, cap)
}

// Reference implements ChanSend.
func (c ChanSendUintptr) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendUintptr) Send(v interface{}) {
	c <- v.(uintptr)
}
