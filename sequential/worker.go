package work

type Worker struct {
	done chan struct{}

	in []<-chan interface{}

	out []chan<- interface{}

	work func([]interface{}) (interface{}, error)
}

func NewWorker(work func([]interface{}) (interface{}, error)) Worker {
	return Worker{in: []<-chan interface{}{}, out: []chan<- interface{}{}, work: work}
}

func (w *Worker) Input(ws ...*Worker) {
	for _, in := range ws {
		var c = make(chan interface{})

		w.in = append(w.in, (<-chan interface{})(c))
		in.out = append(in.out, (chan<- interface{})(c))
	}
}

func (w *Worker) Output(ws ...*Worker) {
	for _, out := range ws {
		var c = make(chan interface{})

		out.in = append(out.in, (<-chan interface{})(c))
		w.out = append(w.out, (chan<- interface{})(c))
	}
}

func (w *Worker) Start() {
	w.done = make(chan struct{})

	go func() {
		for {
			select {
			case <-w.done:
				return

			default:
				if a, err := w.arguments(); err.e == nil {
					if v, err := w.work(a); err == nil {
						w.result(v)
					} else {
						w.result(workerError{e: err})
					}
				} else {
					w.result(err)
				}
			}
		}
	}()
}

func (w *Worker) Stop() {
	go func() {
		w.done <- struct{}{}
		w.done = nil
	}()
}

func (w Worker) Work() (interface{}, error) {
	if a, err := w.arguments(); err.e == nil {
		return w.work(a)
	} else {
		return nil, err.e
	}
}

func (w Worker) arguments() ([]interface{}, workerError) {
	var vs []interface{}

	for _, c := range w.in {
		var v = <-c

		if err, ok := v.(workerError); ok {
			return nil, err
		}

		vs = append(vs, v)
	}

	return vs, workerError{}
}

func (w Worker) result(v interface{}) {
	for _, c := range w.out {
		c <- v
	}
}

type workerError struct {
	e error
}
