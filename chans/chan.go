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

// TODO: Branch, Combine, Compose, Copy, Merge, Replay, Unblock
