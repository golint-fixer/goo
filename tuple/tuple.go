package tuples

import "github.com/willfaught/goo"

var (
	_ lang.Equatable = Tuple2{}
	_ lang.Equatable = Tuple3{}
	_ lang.Equatable = Tuple4{}
	_ lang.Equatable = Tuple5{}
	_ lang.Equatable = Tuple6{}
	_ lang.Equatable = Tuple7{}
)

var (
	_ lang.EquatableDeep = Tuple2{}
	_ lang.EquatableDeep = Tuple3{}
	_ lang.EquatableDeep = Tuple4{}
	_ lang.EquatableDeep = Tuple5{}
	_ lang.EquatableDeep = Tuple6{}
	_ lang.EquatableDeep = Tuple7{}
)

type Tuple2 [2]interface{}

func (t Tuple2) Equals(v interface{}) bool {
	return t == v.(Tuple2)
}

func (t Tuple2) EqualsDeep(v interface{}) bool {
	var u = v.(Tuple2)

	for i := range t {
		if !lang.Equal(t[i], u[i]) {
			return false
		}
	}

	return true
}

func (t Tuple2) First() interface{} {
	return t[0]
}

func (t Tuple2) Second() interface{} {
	return t[1]
}

type Tuple3 [3]interface{}

func (t Tuple3) Equals(v interface{}) bool {
	return t == v.(Tuple3)
}

func (t Tuple3) EqualsDeep(v interface{}) bool {
	var u = v.(Tuple3)

	for i := range t {
		if !lang.Equal(t[i], u[i]) {
			return false
		}
	}

	return true
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

func (t Tuple4) Equals(v interface{}) bool {
	return t == v.(Tuple4)
}

func (t Tuple4) EqualsDeep(v interface{}) bool {
	var u = v.(Tuple4)

	for i := range t {
		if !lang.Equal(t[i], u[i]) {
			return false
		}
	}

	return true
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

func (t Tuple5) Equals(v interface{}) bool {
	return t == v.(Tuple5)
}

func (t Tuple5) EqualsDeep(v interface{}) bool {
	var u = v.(Tuple5)

	for i := range t {
		if !lang.Equal(t[i], u[i]) {
			return false
		}
	}

	return true
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

func (t Tuple6) Equals(v interface{}) bool {
	return t == v.(Tuple6)
}

func (t Tuple6) EqualsDeep(v interface{}) bool {
	var u = v.(Tuple6)

	for i := range t {
		if !lang.Equal(t[i], u[i]) {
			return false
		}
	}

	return true
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

func (t Tuple7) Equals(v interface{}) bool {
	return t == v.(Tuple7)
}

func (t Tuple7) EqualsDeep(v interface{}) bool {
	var u = v.(Tuple7)

	for i := range t {
		if !lang.Equal(t[i], u[i]) {
			return false
		}
	}

	return true
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
