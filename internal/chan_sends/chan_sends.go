//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro chan_sends.go ../../chan_send_bool.go -d {"Name":"Bool","Type":"bool"}
//go:generate goo macro chan_sends.go ../../chan_send_byte.go -d {"Name":"Byte","Type":"byte"}
//go:generate goo macro chan_sends.go ../../chan_send_complex64.go -d {"Name":"Complex64","Type":"complex64"}
//go:generate goo macro chan_sends.go ../../chan_send_complex128.go -d {"Name":"Complex128","Type":"complex128"}
//go:generate goo macro chan_sends.go ../../chan_send_float32.go -d {"Name":"Float32","Type":"float32"}
//go:generate goo macro chan_sends.go ../../chan_send_float64.go -d {"Name":"Float64","Type":"float64"}
//go:generate goo macro chan_sends.go ../../chan_send_int.go -d {"Name":"Int","Type":"int"}
//go:generate goo macro chan_sends.go ../../chan_send_interface.go -d {"Name":"Interface","Type":"interface{}"}
//go:generate goo macro chan_sends.go ../../chan_send_int8.go -d {"Name":"Int8","Type":"int8"}
//go:generate goo macro chan_sends.go ../../chan_send_int16.go -d {"Name":"Int16","Type":"int16"}
//go:generate goo macro chan_sends.go ../../chan_send_int32.go -d {"Name":"Int32","Type":"int32"}
//go:generate goo macro chan_sends.go ../../chan_send_int64.go -d {"Name":"Int64","Type":"int64"}
//go:generate goo macro chan_sends.go ../../chan_send_rune.go -d {"Name":"Rune","Type":"rune"}
//go:generate goo macro chan_sends.go ../../chan_send_string.go -d {"Name":"String","Type":"string"}
//go:generate goo macro chan_sends.go ../../chan_send_uint.go -d {"Name":"Uint","Type":"uint"}
//go:generate goo macro chan_sends.go ../../chan_send_uintptr.go -d {"Name":"Uintptr","Type":"uintptr"}
//go:generate goo macro chan_sends.go ../../chan_send_uint8.go -d {"Name":"Uint8","Type":"uint8"}
//go:generate goo macro chan_sends.go ../../chan_send_uint16.go -d {"Name":"Uint16","Type":"uint16"}
//go:generate goo macro chan_sends.go ../../chan_send_uint32.go -d {"Name":"Uint32","Type":"uint32"}
//go:generate goo macro chan_sends.go ../../chan_send_uint64.go -d {"Name":"Uint64","Type":"uint64"}

package goo

/// {{if false}}
type __FIELD_Type__ struct{}
type Chan__FIELD_Name__ chan struct{}
type Chan interface{}
type ChanSend interface{} /// {{end}}

var _ ChanSend = ChanSend__FIELD_Name__(nil)

// ChanSend__FIELD_Name__ is a send channel of __FIELD_Type__.
type ChanSend__FIELD_Name__ chan __FIELD_Type__

// Cap implements ChanSend.
func (c ChanSend__FIELD_Name__) Cap() int {
	return cap(c)
}

// Close implements Chan.
func (c ChanSend__FIELD_Name__) Close() {
	close(c)
}

// Len implements ChanSend.
func (c ChanSend__FIELD_Name__) Len() int {
	return len(c)
}

// Make implements ChanSend.
func (c ChanSend__FIELD_Name__) Make(cap int) Chan {
	return make(Chan__FIELD_Name__, cap)
}

// Send implements Chan.
func (c ChanSend__FIELD_Name__) Send(v interface{}) {
	c <- v.(__FIELD_Type__)
}
