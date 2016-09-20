package goo

var (
	_ Equatable = Tuple2{}
	_ Equatable = Tuple3{}
	_ Equatable = Tuple4{}
	_ Equatable = Tuple5{}
	_ Equatable = Tuple6{}
	_ Equatable = Tuple7{}
)

type Tuple2 [2]interface{}

func (t Tuple2) Equals(e Equatable) bool {
	return t == e.(Tuple2)
}

func (t Tuple2) NotEquals(e Equatable) bool {
	return t != e.(Tuple2)
}

func (t Tuple2) First() interface{} {
	return t[0]
}

func (t Tuple2) Second() interface{} {
	return t[1]
}

type Tuple3 [3]interface{}

func (t Tuple3) Equals(e Equatable) bool {
	return t == e.(Tuple3)
}

func (t Tuple3) NotEquals(e Equatable) bool {
	return t != e.(Tuple3)
}

func (t Tuple3) First() interface{} {
	return t[0]
}

func (t Tuple3) Second() interface{} {
	return t[1]
}

func (t Tuple3) Third() interface{} {
	return t[2]
}

type Tuple4 [4]interface{}

func (t Tuple4) Equals(e Equatable) bool {
	return t == e.(Tuple4)
}

func (t Tuple4) NotEquals(e Equatable) bool {
	return t != e.(Tuple4)
}

func (t Tuple4) First() interface{} {
	return t[0]
}

func (t Tuple4) Second() interface{} {
	return t[1]
}

func (t Tuple4) Third() interface{} {
	return t[2]
}

func (t Tuple4) Fourth() interface{} {
	return t[3]
}

type Tuple5 [5]interface{}

func (t Tuple5) Equals(e Equatable) bool {
	return t == e.(Tuple5)
}

func (t Tuple5) NotEquals(e Equatable) bool {
	return t != e.(Tuple5)
}

func (t Tuple5) First() interface{} {
	return t[0]
}

func (t Tuple5) Second() interface{} {
	return t[1]
}

func (t Tuple5) Third() interface{} {
	return t[2]
}

func (t Tuple5) Fourth() interface{} {
	return t[3]
}

func (t Tuple5) Fifth() interface{} {
	return t[4]
}

type Tuple6 [6]interface{}

func (t Tuple6) Equals(e Equatable) bool {
	return t == e.(Tuple6)
}

func (t Tuple6) NotEquals(e Equatable) bool {
	return t == e.(Tuple6)
}

func (t Tuple6) First() interface{} {
	return t[0]
}

func (t Tuple6) Second() interface{} {
	return t[1]
}

func (t Tuple6) Third() interface{} {
	return t[2]
}

func (t Tuple6) Fourth() interface{} {
	return t[3]
}

func (t Tuple6) Fifth() interface{} {
	return t[4]
}

func (t Tuple6) Sixth() interface{} {
	return t[5]
}

type Tuple7 [7]interface{}

func (t Tuple7) Equals(e Equatable) bool {
	return t == e.(Tuple7)
}

func (t Tuple7) NotEquals(e Equatable) bool {
	return t != e.(Tuple7)
}

func (t Tuple7) First() interface{} {
	return t[0]
}

func (t Tuple7) Second() interface{} {
	return t[1]
}

func (t Tuple7) Third() interface{} {
	return t[2]
}

func (t Tuple7) Fourth() interface{} {
	return t[3]
}

func (t Tuple7) Fifth() interface{} {
	return t[4]
}

func (t Tuple7) Sixth() interface{} {
	return t[5]
}

func (t Tuple7) Seventh() interface{} {
	return t[6]
}
