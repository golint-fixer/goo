package work

import "sync"

type Signal struct {
	m       *sync.Mutex
	off, on chan struct{}
	state   bool
}

func NewSignal(state bool) *Signal {
	var s = &Signal{off: make(chan struct{}), on: make(chan struct{}), state: state}
	var c = s.off

	if state {
		c = s.on
	}

	close(c)

	return s
}

func (s *Signal) Off() {
	s.m.Lock()

	defer s.m.Unlock()

	if !s.state {
		return
	}

	s.state = false
	s.on = make(chan struct{})
	close(s.off)
}

func (s *Signal) On() {
	s.m.Lock()

	defer s.m.Unlock()

	if s.state {
		return
	}

	s.state = true
	s.off = make(chan struct{})
	close(s.on)
}

func (s *Signal) WaitOff() {
	s.m.Lock()

	var c = s.off

	s.m.Unlock()

	<-c
}

func (s *Signal) WaitOn() {
	s.m.Lock()

	var c = s.on

	s.m.Unlock()

	<-c
}
