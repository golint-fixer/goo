package goo

// Chan is a channel.
type Chan interface {
	// Cap returns the capacity.
	Cap() int

	// ChanReceive returns the receive channel.
	ChanReceive() ChanReceive

	// ChanSend returns the send channel.
	ChanSend() ChanSend

	// Close closes.
	Close()

	// Len returns the length.
	Len() int

	// Make returns a new chan with capacity c.
	Make(c int) Chan

	// Receive returns the next element.
	Receive() interface{}

	// ReceiveCheck returns the next element and whether there is a next element.
	ReceiveCheck() (interface{}, bool)

	// Send sends v.
	Send(v interface{})
}

/*
var (
	_ Chan = ChanBool(nil)
	_ Chan = ChanInt(nil)
	_ Chan = ChanInterface(nil)
	_ Chan = ChanRune(nil)
	_ Chan = ChanString(nil)
)

func Branch(c ChanReceive, n int) []ChanReceive {
	var rs, ss = make([]ChanReceive, n), make([]ChanSend, n)

	for i := 0; i < n; i++ {
		var b = c.Make(c.Cap())

		rs[i], ss[i] = b, b
	}

	go func() {
		for i := 0; ; i++ {
			var v, ok = c.ReceiveCheck()

			if !ok {
				break
			}

			ss[i%n].Send(v)
		}

		for i := 0; i < n; i++ {
			ss[i].Close()
		}
	}()

	return rs
}

func Broadcast(c ...ChanSend) ChanSend {
	var broadcast = c[0].Make(c[0].Cap())

	go func() {
		for {
			var v, ok = broadcast.ReceiveCheck()

			if !ok {
				break
			}

			for _, c := range c {
				c.Send(v)
			}
		}

		for _, c := range c {
			c.Close()
		}
	}()

	return broadcast
}

func Compose(r ChanReceive, s ChanSend) Chan {
	var c = r.Make(s.Cap())

	go func() {
		var v, ok = r.ReceiveCheck()

		if !ok {
			c.Close()

			return
		}

		c.Send(v)
	}()

	go func() {
		var v, ok = c.ReceiveCheck()

		if !ok {
			s.Close()

			return
		}

		s.Send(v)
	}()

	return c
}

func Copy(c ChanReceive, n int) []ChanReceive {
	var rs []ChanReceive
	var ss []ChanSend
	var ns []chan struct{}
	var vs []interface{}
	var done bool
	var vsm, donem sync.Mutex

	for i := 0; i < n; i++ {
		var copy = c.Make(c.Cap())

		rs = append(rs, copy)
		ss = append(ss, copy)
		ns = append(ns, make(chan struct{}, 1))

		go func() {
			var j int

			for {
				donem.Lock()
				var d = done
				donem.Unlock()

				if d {
					break
				}

				<-ns[i]

				vsm.Lock()
				var s = vs[j:len(vs)]
				vsm.Unlock()

				for _, v := range s {
					ss[i].Send(v)
				}

				j += len(s)
			}

			ss[i].Close()
		}()
	}

	go func() {
		for {
			var v, ok = c.ReceiveCheck()

			if !ok {
				break
			}

			vsm.Lock()
			vs = append(vs, v)
			vsm.Unlock()

			for _, n := range ns {
				if len(n) == 0 {
					n <- struct{}{}
				}
			}
		}

		donem.Lock()
		done = true
		donem.Unlock()
	}()

	return rs
}

func Cycle(c ChanReceive) ChanReceive {
	var vs []interface{}
	var cycle = make(ChanInterface, c.Cap())

	go func() {
		for {
			var v, ok = c.ReceiveCheck()

			if !ok {
				break
			}

			vs = append(vs, v)
			cycle.Send(v)
		}

		for {
			for _, v := range vs {
				cycle.Send(v)
			}
		}
	}()

	return cycle
}

func Merge(c ...ChanReceive) ChanReceive {
	var merged = c[0].Make(0)
	var w sync.WaitGroup

	w.Add(len(c))

	for i, l := 0, len(c); i < l; i++ {
		go func(r ChanReceive) {
			for {
				var v, ok = r.ReceiveCheck()

				if !ok {
					break
				}

				merged.Send(v)
			}

			w.Done()
		}(c[i])
	}

	go func() {
		w.Wait()
		merged.Close()
	}()

	return merged
}

func Repeat(v interface{}) ChanReceive {
	var c = make(ChanInterface)

	go func() {
		for {
			c.Send(v)
		}
	}()

	return c
}

func Replicate(v interface{}, n int) ChanReceive {
	var c = make(ChanInterface)

	go func() {
		for i := 0; i < n; i++ {
			c.Send(v)
		}

		c.Close()
	}()

	return c
}

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
*/
