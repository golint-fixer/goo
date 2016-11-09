package goo

var _ ChanSend = ChanSendFloat64(nil)

// ChanSendFloat64 is a send channel of float64.
type ChanSendFloat64 chan float64

// Cap implements ChanSend.
func (c ChanSendFloat64) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendFloat64) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendFloat64) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendFloat64) Make(cap int) Chan {
	return make(ChanFloat64, cap)
}

// Send implements Chan.
func (c ChanSendFloat64) Send(v interface{}) {
	c <- v.(float64)
}
