package goo

var _ ChanSend = ChanSendFloat32(nil)

// ChanSendFloat32 is a send channel of float32.
type ChanSendFloat32 chan float32

// Cap implements ChanSend.
func (c ChanSendFloat32) Cap() int {
	return cap(c)
}

// Close implements ChanSend.
func (c ChanSendFloat32) Close() {
	close(c)
}

// Dereference implements ChanSend.
func (c *ChanSendFloat32) Dereference() Value {
	return *c
}

// Len implements ChanSend.
func (c ChanSendFloat32) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSendFloat32) Make(cap int) Chan {
	return make(ChanFloat32, cap)
}

// Reference implements ChanSend.
func (c ChanSendFloat32) Reference() Pointer {
	return &c
}

// Send implements Chan.
func (c ChanSendFloat32) Send(v interface{}) {
	c <- v.(float32)
}
