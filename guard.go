package goo

import "sync"

type Guard struct {
	i interface{}
	m *sync.Mutex
}

func NewGuard(i interface{}) *Guard {
	return &Guard{i: i, m: &sync.Mutex{}}
}

func (g *Guard) Get() interface{} {
	g.m.Lock()

	defer g.m.Unlock()

	return g.i
}

func (g *Guard) Map(f func(interface{}) interface{}) {
	g.m.Lock()

	defer g.m.Unlock()

	g.i = f(g.i)
}

func (g *Guard) Set(i interface{}) {
	g.m.Lock()

	defer g.m.Unlock()

	g.i = i
}
