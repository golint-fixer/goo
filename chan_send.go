package goo

// ChanSend is a send channel.
type ChanSend interface {
	// Cap returns the capacity.
	Cap() int

	// Close closes.
	Close()

	// Len returns the length.
	Len() int

	// Make returns a new chan with capacity c.
	Make(c int) Chan

	// Send sends v.
	Send(v interface{})
}
