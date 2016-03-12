package lang

type Iterable interface {
	More() bool

	Next() interface{}
}
