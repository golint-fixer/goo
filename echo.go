package goo

import "sync"

type Echo struct {
	m      *sync.Mutex
	pinged chan struct{}
	pings  int
}

func NewEcho() *Echo {
	return &Echo{m: &sync.Mutex{}, pinged: make(chan struct{}, 1)}
}

func (e *Echo) Ping(n int) {
	e.m.Lock()

	defer e.m.Unlock()

	if e.pings == 0 {
		close(e.pinged)
	}

	e.pings += n
}

func (e *Echo) Pong(n int) {
	for {
		e.m.Lock()

		if e.pings == 0 {
			var c = make(chan struct{})

			e.pinged = c
			e.m.Unlock()

			<-c

			continue
		}

		var rest = e.pings - n

		if rest < 0 {
			e.pings = 0
			n += rest
			e.m.Unlock()

			continue
		}

		e.pings = rest
		e.m.Unlock()

		break
	}
}
