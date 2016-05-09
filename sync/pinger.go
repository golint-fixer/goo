package sync

import "sync"

type Pinger interface {
	Ping()

	Pong()
}

func NewPinger() Pinger {
	return &pinger{lock: &sync.Mutex{}, pinged: make(chan struct{}, 1)}
}

type pinger struct {
	lock   *sync.Mutex
	pinged chan struct{}
	pings  int
}

func (p *pinger) Ping() {
	p.lock.Lock()
	p.pings++

	if p.pings == 0 {
		p.pinged <- struct{}{}
	}

	p.lock.Unlock()
}

func (p *pinger) Pong() {
	p.lock.Lock()

	if p.pings == 0 {
		p.lock.Unlock()
		<-p.pinged
		p.lock.Lock()
	}

	p.pings--
	p.lock.Unlock()
}
