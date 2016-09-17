package goo

import "sync"

type Count struct {
	m *sync.Mutex
	n int
}

func NewCount(n int) *Count {
	return &Count{m: &sync.Mutex{}, n: n}
}

func (c *Count) Add(n int) {
	c.m.Lock()

	defer c.m.Unlock()

	c.n += n
}

func (c *Count) Get() int {
	c.m.Lock()

	defer c.m.Unlock()

	return c.n
}

func (c *Count) Set(n int) {
	c.m.Lock()

	defer c.m.Unlock()

	c.n = n
}

func (c *Count) Sub(n int) {
	c.m.Lock()

	defer c.m.Unlock()

	c.n -= n
}
