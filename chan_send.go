package goo

type ChanSend interface {
	Cap() int

	Close()

	Len() int

	Make(c int) Chan

	Send(v interface{})
}
