package chans

type Chan interface {
	Cap() int

	ChanReceive() ChanReceive

	ChanSend() ChanSend

	Close()

	Len() int

	Make(c int) Chan

	Receive() interface{}

	ReceiveCheck() (interface{}, bool)

	Send(v interface{})
}

var (
	_ Chan = ChanBool(nil)
	_ Chan = ChanInt(nil)
	_ Chan = ChanInterface(nil)
	_ Chan = ChanRune(nil)
	_ Chan = ChanString(nil)
)

// TODO: Branch, Combine, Compose, Copy, CopyPrimSlice, Merge, Replay

func Unblock(c ChanSend) ChanSend {
	var unblocked = c.Make(0)
	var ss = make(chan []interface{})

	go func() {
		var s []interface{}

		for {
			var v, ok = unblocked.ReceiveCheck()

			if !ok {
				break
			}

			s = append(s, v)

			select {
			case ss <- s:
				s = nil

			default:
			}
		}

		close(ss)
	}()

	go func() {
		for s := range ss {
			for _, v := range s {
				c.Send(v)
			}
		}

		c.Close()
	}()

	return unblocked
}
