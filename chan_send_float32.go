package goo

var _ ChanSend = ChanSendFloat32(nil)

// ChanSendFloat32 is a send channel of float32.
type ChanSendFloat32 chan float32

// Cap implements ChanSend.
func (c ChanSendFloat32) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSendFloat32) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSendFloat32) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendFloat32) Make(cap int) Chan {
	return make(ChanFloat32, cap)
}

// Send implements Chan.
func (c ChanSendFloat32) Send(v interface{}) {
	c <- v.(float32)
}
