package chans

type ChanReceive interface {
	Cap() int

	Len() int

	Make(c int) Chan

	Receive() interface{}

	ReceiveCheck() (interface{}, bool)
}
