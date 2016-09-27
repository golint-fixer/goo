//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo -in tuple.go -out ../../tuple_2.go -json {"n":2,"f":["First","Second"]}
//go:generate goo -in tuple.go -out ../../tuple_3.go -json {"n":3,"f":["First","Second","Third"]}
//go:generate goo -in tuple.go -out ../../tuple_4.go -json {"n":4,"f":["First","Second","Third","Fourth"]}
//go:generate goo -in tuple.go -out ../../tuple_5.go -json {"n":5,"f":["First","Second","Third","Fourth","Fifth"]}
//go:generate goo -in tuple.go -out ../../tuple_6.go -json {"n":6,"f":["First","Second","Third","Fourth","Fifth","Sixth"]}
//go:generate goo -in tuple.go -out ../../tuple_7.go -json {"n":7,"f":["First","Second","Third","Fourth","Fifth","Sixth","Seventh"]}

package goo

/// {{$n := .n}}

/// {{if false}}
const __VAR_n__ = 0 /// {{end}}

/// {{if false}}
/// {{$i := 0}}
var __VAR_i__ int /// {{end}}

var _ Equatable = Tuple__VAR_n__{}

/// {{if false}}
type Equatable interface{} /// {{end}}

type Tuple__VAR_n__ [__VAR_n__]interface{}

func (t Tuple__VAR_n__) Equals(e Equatable) bool {
	return t == e.(Tuple__VAR_n__)
}

func (t Tuple__VAR_n__) NotEquals(e Equatable) bool {
	return t != e.(Tuple__VAR_n__)
}

/// {{range $i, $s := .f}}
func (t Tuple__VAR_n__) __VAR_s__() interface{} {
	return t[__VAR_i__]
}

/// {{end}}
