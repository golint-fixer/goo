//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro chan_receives.go ../../chan_receive_bool.go -d {"Name":"Bool","Type":"bool"}
//go:generate goo macro chan_receives.go ../../chan_receive_byte.go -d {"Name":"Byte","Type":"byte"}
//go:generate goo macro chan_receives.go ../../chan_receive_complex64.go -d {"Name":"Complex64","Type":"complex64"}
//go:generate goo macro chan_receives.go ../../chan_receive_complex128.go -d {"Name":"Complex128","Type":"complex128"}
//go:generate goo macro chan_receives.go ../../chan_receive_float32.go -d {"Name":"Float32","Type":"float32"}
//go:generate goo macro chan_receives.go ../../chan_receive_float64.go -d {"Name":"Float64","Type":"float64"}
//go:generate goo macro chan_receives.go ../../chan_receive_int.go -d {"Name":"Int","Type":"int"}
//go:generate goo macro chan_receives.go ../../chan_receive_interface.go -d {"Name":"Interface","Type":"interface{}"}
//go:generate goo macro chan_receives.go ../../chan_receive_int8.go -d {"Name":"Int8","Type":"int8"}
//go:generate goo macro chan_receives.go ../../chan_receive_int16.go -d {"Name":"Int16","Type":"int16"}
//go:generate goo macro chan_receives.go ../../chan_receive_int32.go -d {"Name":"Int32","Type":"int32"}
//go:generate goo macro chan_receives.go ../../chan_receive_int64.go -d {"Name":"Int64","Type":"int64"}
//go:generate goo macro chan_receives.go ../../chan_receive_rune.go -d {"Name":"Rune","Type":"rune"}
//go:generate goo macro chan_receives.go ../../chan_receive_string.go -d {"Name":"String","Type":"string"}
//go:generate goo macro chan_receives.go ../../chan_receive_uint.go -d {"Name":"Uint","Type":"uint"}
//go:generate goo macro chan_receives.go ../../chan_receive_uintptr.go -d {"Name":"Uintptr","Type":"uintptr"}
//go:generate goo macro chan_receives.go ../../chan_receive_uint8.go -d {"Name":"Uint8","Type":"uint8"}
//go:generate goo macro chan_receives.go ../../chan_receive_uint16.go -d {"Name":"Uint16","Type":"uint16"}
//go:generate goo macro chan_receives.go ../../chan_receive_uint32.go -d {"Name":"Uint32","Type":"uint32"}
//go:generate goo macro chan_receives.go ../../chan_receive_uint64.go -d {"Name":"Uint64","Type":"uint64"}

package goo

/// {{if false}}
type __FIELD_Type__ struct{}
type Chan__FIELD_Name__ chan struct{}
type Chan interface{}
type ChanReceive interface{}
type Pointer interface{}
type Value interface{} /// {{end}}

var _ ChanReceive = ChanReceive__FIELD_Name__(nil)

// ChanReceive__FIELD_Name__ is a receive channel of __FIELD_Type__.
type ChanReceive__FIELD_Name__ chan __FIELD_Type__

// Cap implements ChanReceive.
func (c ChanReceive__FIELD_Name__) Cap() int {
	return cap(c)
}

// Dereference implements ChanReceive.
func (c *ChanReceive__FIELD_Name__) Dereference() Value {
	return *c
}

// Len implements ChanReceive.
func (c ChanReceive__FIELD_Name__) Len() int {
	return len(c)
}

// Make implements ChanReceive.
func (c ChanReceive__FIELD_Name__) Make(cap int) Chan {
	return make(Chan__FIELD_Name__, cap)
}

// Receive implements ChanReceive.
func (c ChanReceive__FIELD_Name__) Receive() interface{} {
	return <-c
}

// ReceiveCheck implements ChanReceive.
func (c ChanReceive__FIELD_Name__) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

// Reference implements ChanReceive.
func (c ChanReceive__FIELD_Name__) Reference() Pointer {
	return &c
}
