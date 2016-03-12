package lang

type Iterator interface {
	More() bool

	Next() interface{}
}
