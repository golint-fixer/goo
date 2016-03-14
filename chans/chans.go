package chans

import "github.com/willfaught/lang"

var (
	_ lang.Equatable = ChanBool(nil)
	_ lang.Equatable = ChanInt(nil)
	_ lang.Equatable = ChanInterface(nil)
	_ lang.Equatable = ChanRune(nil)
	_ lang.Equatable = ChanString(nil)
)

var (
	_ lang.EquatableDeep = ChanBool(nil)
	_ lang.EquatableDeep = ChanInt(nil)
	_ lang.EquatableDeep = ChanInterface(nil)
	_ lang.EquatableDeep = ChanRune(nil)
	_ lang.EquatableDeep = ChanString(nil)
)

type ChanBool chan bool

func (c ChanBool) Cap() int {
	return cap(c)
}

func (c ChanBool) ChanReceive() ChanReceive {
	return c
}

func (c ChanBool) ChanSend() ChanSend {
	return c
}

func (c ChanBool) Close() {
	close(c)
}

func (c ChanBool) Equals(v interface{}) bool {
	var d = v.(ChanBool)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc != vd {
			return false
		}
	}

	return true
}

func (c ChanBool) EqualsDeep(v interface{}) bool {
	var d = v.(ChanBool)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if !lang.Equal(vc, vd) {
			return false
		}
	}

	return true
}

func (c ChanBool) Len() int {
	return len(c)
}

func (c ChanBool) Make(cap int) Chan {
	return make(ChanBool, cap)
}

func (c ChanBool) Receive() interface{} {
	return <-c
}

func (c ChanBool) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanBool) Send(v interface{}) {
	c <- v.(bool)
}

type ChanInt chan int

func (c ChanInt) Cap() int {
	return cap(c)
}

func (c ChanInt) ChanReceive() ChanReceive {
	return c
}

func (c ChanInt) ChanSend() ChanSend {
	return c
}

func (c ChanInt) Close() {
	close(c)
}

func (c ChanInt) Equals(v interface{}) bool {
	var d = v.(ChanInt)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc != vd {
			return false
		}
	}

	return true
}

func (c ChanInt) EqualsDeep(v interface{}) bool {
	var d = v.(ChanInt)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if !lang.Equal(vc, vd) {
			return false
		}
	}

	return true
}

func (c ChanInt) Len() int {
	return len(c)
}

func (c ChanInt) Make(cap int) Chan {
	return make(ChanInt, cap)
}

func (c ChanInt) Receive() interface{} {
	return <-c
}

func (c ChanInt) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanInt) Send(v interface{}) {
	c <- v.(int)
}

type ChanInterface chan interface{}

func (c ChanInterface) Cap() int {
	return cap(c)
}

func (c ChanInterface) ChanReceive() ChanReceive {
	return c
}

func (c ChanInterface) ChanSend() ChanSend {
	return c
}

func (c ChanInterface) Close() {
	close(c)
}

func (c ChanInterface) Equals(v interface{}) bool {
	var d = v.(ChanInterface)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc != vd {
			return false
		}
	}

	return true
}

func (c ChanInterface) EqualsDeep(v interface{}) bool {
	var d = v.(ChanInterface)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if !lang.Equal(vc, vd) {
			return false
		}
	}

	return true
}

func (c ChanInterface) Len() int {
	return len(c)
}

func (c ChanInterface) Make(cap int) Chan {
	return make(ChanInterface, cap)
}

func (c ChanInterface) Receive() interface{} {
	return <-c
}

func (c ChanInterface) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanInterface) Send(v interface{}) {
	c <- v
}

type ChanRune chan rune

func (c ChanRune) Cap() int {
	return cap(c)
}

func (c ChanRune) ChanReceive() ChanReceive {
	return c
}

func (c ChanRune) ChanSend() ChanSend {
	return c
}

func (c ChanRune) Close() {
	close(c)
}

func (c ChanRune) Equals(v interface{}) bool {
	var d = v.(ChanRune)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc != vd {
			return false
		}
	}

	return true
}

func (c ChanRune) EqualsDeep(v interface{}) bool {
	var d = v.(ChanRune)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if !lang.Equal(vc, vd) {
			return false
		}
	}

	return true
}

func (c ChanRune) Len() int {
	return len(c)
}

func (c ChanRune) Make(cap int) Chan {
	return make(ChanRune, cap)
}

func (c ChanRune) Receive() interface{} {
	return <-c
}

func (c ChanRune) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanRune) Send(v interface{}) {
	c <- v.(rune)
}

type ChanString chan string

func (c ChanString) Cap() int {
	return cap(c)
}

func (c ChanString) ChanReceive() ChanReceive {
	return c
}

func (c ChanString) ChanSend() ChanSend {
	return c
}

func (c ChanString) Close() {
	close(c)
}

func (c ChanString) Equals(v interface{}) bool {
	var d = v.(ChanString)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if vc != vd {
			return false
		}
	}

	return true
}

func (c ChanString) EqualsDeep(v interface{}) bool {
	var d = v.(ChanString)

	for {
		var vc, okc = c.ReceiveCheck()
		var vd, okd = d.ReceiveCheck()

		if okc != okd {
			return false
		}

		if !okc {
			break
		}

		if !lang.Equal(vc, vd) {
			return false
		}
	}

	return true
}

func (c ChanString) Len() int {
	return len(c)
}

func (c ChanString) Make(cap int) Chan {
	return make(ChanString, cap)
}

func (c ChanString) Receive() interface{} {
	return <-c
}

func (c ChanString) ReceiveCheck() (interface{}, bool) {
	var v, ok = <-c

	return v, ok
}

func (c ChanString) Send(v interface{}) {
	c <- v.(string)
}
