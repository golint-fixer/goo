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
