package goo

/*
import "sync"

const (
	CapacitySmall  = 1 << 3
	CapacityMedium = 1 << 6
	CapacityLarge  = 1 << 9
	CapacityHuge   = 1 << 12
)

const (
	GrowSlow   = 1.25
	GrowMedium = 2.0
	GrowFast   = 4.0
)

const (
	QueueSmall  = 1 << 7
	QueueMedium = 1 << 10
	QueueLarge  = 1 << 13
	QueueHuge   = 1 << 16
)

const (
	ShrinkSlow   = 0.75
	ShrinkMedium = 0.5
	ShrinkFast   = 0.25
)

const (
	ThresholdLow    = 0.25
	ThresholdMedium = 0.5
	ThresholdHigh   = 0.75
)

type Strategy func(used, total int) int

var (
	StrategyCapacitySmall   Strategy = Combine(capMinSmall, Max(CapacityMedium))
	StrategyCapacityMedium  Strategy = Combine(Min(CapacityMedium), Max(CapacityLarge))
	StrategyCapacityLarge   Strategy = Combine(Min(CapacityLarge), capMaxHuge)
	StrategyCapacityElastic Strategy = Combine(capMinSmall, capMaxHuge)
)

var (
	StrategySlow   Strategy = Combine(shrinkLowSlow, growHighSlow)
	StrategyFast   Strategy = Combine(shrinkLowFast, growHighFast)
	StrategyLight  Strategy = Combine(shrinkLowSlow, growHighFast)
	StrategyHeavy  Strategy = Combine(shrinkLowFast, growHighSlow)
	StrategySteady Strategy = Combine(Shrink(ThresholdLow, ShrinkMedium), Grow(ThresholdHigh, GrowMedium))
)

var StrategyCapacityDefault Strategy = Combine(Combine(Min(CapacitySmall), Max(CapacityHuge)), StrategySteady)

var (
	capMinSmall = Min(CapacitySmall)
	capMaxHuge  = Max(CapacityHuge)

	shrinkLowSlow = Shrink(ThresholdLow, ShrinkSlow)
	shrinkLowFast = Shrink(ThresholdLow, ShrinkFast)

	growHighSlow = Grow(ThresholdHigh, GrowSlow)
	growHighFast = Grow(ThresholdHigh, GrowFast)
)

func Combine(s ...Strategy) Strategy {
	var c = s[0]

	for i, l := 1, len(s); i < l; i++ {
		c = func(used, total int) int {
			return c(used, s[i](used, total))
		}
	}

	return c
}

func Constant(n int) Strategy {
	return func(used, total int) int {
		return n
	}
}

func Grow(threshold, grow float32) Strategy {
	return func(used, total int) int {
		if float32(used)/float32(total) < threshold {
			return int(float32(total) * grow)
		}

		return total
	}
}

func Max(max int) Strategy {
	return func(used, total int) int {
		if total > max {
			return max
		}

		return total
	}
}

func Min(min int) Strategy {
	return func(used, total int) int {
		if total < min {
			return min
		}

		return total
	}
}

func Shrink(threshold, shrink float32) Strategy {
	return func(used, total int) int {
		if float32(used)/float32(total) > threshold {
			return int(float32(total) * shrink)
		}

		return total
	}
}

type Concurrent interface {
	Add(work func())
	Wait()
}

var _ Concurrent = &pool{}

func NewConcurrent(capacity, s Strategy) Concurrent {
	return newPool(s)
}

type guard struct {
	i interface{}
	m *sync.Mutex
}

func newGuard(i interface{}) *guard {
	return &guard{i: i, m: &sync.Mutex{}}
}

func (g *guard) get() interface{} {
	g.m.Lock()

	defer g.m.Unlock()

	return g.i
}

func (g *guard) set(i interface{}) {
	g.m.Lock()

	defer g.m.Unlock()

	g.i = i
}

type pool struct {
	strategy Strategy
	wait     *guard
	works    *concQueue
	workers  *concQueue
}

func newPool(s Strategy) *pool {
	var p = &pool{strategy: s, wait: newGuard(make(chan struct{})), works: newQueue(), workers: newQueue()}

	go p.run()

	for i, n := 0, s(0, 0); i < n; i++ {
		var w = newConcurrentWorker(p)

		go w.run()

		p.workers.in <- w
	}

	return p
}

func (p *pool) Add(work func()) {
	p.works.in <- work
	p.resize()
}

func (p *pool) Wait() {
	p.waitm.Lock()

	var c = p.wait

	p.waitm.Unlock()
	<-p.wait
}

func (p *pool) resize() {
	var newTotal = p.strategy(len(p.workerChan), cap(p.workers))

	if cap(p.workers) < newTotal {
		var c = make(chan worker, newTotal)

		p.workerChanMutex.Lock()

		for w := range p.workerChan {
			c <- w
		}

		close(p.workerChan)
		p.workerChan = c
		p.workerChanMutex.Unlock()

		for i := 0; i < newTotal-cap(p.workers); i++ {
			p.workerChan <- worker{pool: p, works: make(chan Work)}
		}
	} else if cap(p.workers) > newTotal {

	}
}

func (p *pool) run() {
	for work := range p.works.out {
		(<-p.workers).works <- work
		p.resize()
	}
}

type concQueue struct {
	cap, len int
	done     chan struct{}
	in, out  chan interface{}
	m        *sync.Mutex
	w        *sync.WaitGroup
}

func newQueue(cap int) *concQueue {
	return &concQueue{done: make(chan struct{}), in: make(chan interface{}, 1), out: make(chan interface{}, 1)}
}

func (q *concQueue) main(capacity int) {
	var ins, outs []interface{}
	var n int

	if capacity > 0 {
		ins = make([]interface{}, 0, capacity)
	}

	for {
		if capacity > 0 && len(ins) == cap(ins) {
			select {
			case q.out <- outs[i]:

			case <-q.done:
				return
			}
		} else if false {
			select {
			case in := <-q.in:
				ins = append(ins, in)

			case <-q.done:
				return
			}
		} else {
			select {
			case in := <-q.in:
				ins = append(ins, in)

			case q.out <- outs[i]:

			case <-q.done:
				return
			}
		}
	}
}

func (q *concQueue) getCap() int {
	q.m.Lock()

	defer q.m.Unlock()

	return q.cap
}

func (q *concQueue) close() {
	q.done <- struct{}{}
	close(q.done)
	close(q.in)
	close(q.out)
}

func (q *concQueue) getLen() int {
	q.m.Lock()

	defer q.m.Unlock()

	return q.len
}

func (q *concQueue) receive(in, out chan []interface{}) {
	for _, items := range in {
	Loop:
		for _, item := range q.in {
			items = append(items, item)
			q.m.Lock()
			q.len++
			q.m.Unlock()

			select {
			case out <- items:
				break Loop

			default:
			}
		}
	}
}

func (q *concQueue) send(in, out chan []interface{}) {
	for _, items := range in {
		for _, item := range items {
			q.out <- item
			q.m.Lock()
			q.len--
			q.m.Unlock()

			select {
			case q.sent <- struct{}{}:

			default:
			}
		}

		out <- items[:]
	}
}

func (q *concQueue) run() {
	var a, b = make(chan []interface{}, 1), make(chan []interface{}, 1)

	a <- make([]interface{}, 0, 1)
	b <- make([]interface{}, 0, 1)

	go q.receive(a, b)
	go q.send(b, a)

	<-q.done
	close(a)
	close(b)
}

type concurrentWorker struct {
	pool  *pool
	works chan func()
}

func newConcurrentWorker(p *pool) *concurrentWorker {
	return &concurrentWorker{pool: p, works: make(chan func())}
}

func (c *concurrentWorker) run() {
	for work := range w.works {
		work()
		c.pool.workers.in <- w
		c.pool.resize()
	}
}
*/
