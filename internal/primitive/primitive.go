//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo -in primitive.go -out ../../primitive_bool.go -json {"equatable":true,"name":"Bool","rec":"b","prim":"bool","zero":"false"}
//go:generate goo -in primitive.go -out ../../primitive_byte.go -json {"comparable":true,"integer":true,"name":"Byte","number":true,"rec":"b","prim":"byte","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_complex64.go -json {"name":"Complex64","number":true,"rec":"c","prim":"complex64","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_complex128.go -json {"name":"Complex128","number":true,"rec":"c","prim":"complex128","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_float32.go -json {"comparable":true,"name":"Float32","number":true,"rec":"f","prim":"float32","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_float64.go -json {"comparable":true,"name":"Float64","number":true,"rec":"f","prim":"float64","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int.go -json {"comparable":true,"integer":true,"name":"Int","number":true,"rec":"i","prim":"int","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int8.go -json {"comparable":true,"integer":true,"name":"Int8","number":true,"rec":"i","prim":"int8","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int16.go -json {"comparable":true,"integer":true,"name":"Int16","number":true,"rec":"i","prim":"int16","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int32.go -json {"comparable":true,"integer":true,"name":"Int32","number":true,"rec":"i","prim":"int32","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int64.go -json {"comparable":true,"integer":true,"name":"Int64","number":true,"rec":"i","prim":"int64","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_rune.go -json {"comparable":true,"integer":true,"name":"Rune","number":true,"rec":"r","prim":"rune","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_string.go -json {"comparable":true,"name":"String","rec":"s","prim":"string","zero":"\"\""}
//go:generate goo -in primitive.go -out ../../primitive_uint.go -json {"comparable":true,"integer":true,"name":"Uint","number":true,"rec":"u","prim":"uint","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uintptr.go -json {"comparable":true,"integer":true,"name":"Uintptr","number":true,"rec":"u","prim":"uintptr","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint8.go -json {"comparable":true,"integer":true,"name":"Uint8","number":true,"rec":"u","prim":"uint8","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint16.go -json {"comparable":true,"integer":true,"name":"Uint16","number":true,"rec":"u","prim":"uint16","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint32.go -json {"comparable":true,"integer":true,"name":"Uint32","number":true,"rec":"u","prim":"uint32","zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint64.go -json {"comparable":true,"integer":true,"name":"Uint64","number":true,"rec":"u","prim":"uint64","zero":"0"}

package goo

/// {{if false}}
var __zero__ = 0 /// {{end}}

/// {{if .integer}}
var _ Integer = __name__(__zero__)

/// {{else if .number}}
var _ Number = __name__(__zero__)

/// {{else if .comparable}}
var _ Comparable = __name__(__zero__)

/// {{else if .equatable}}
var _ Equatable = __name__(__zero__) /// {{end}}

// {{.name}}Zero is the {{.name}} zero value.
var __name__Zero = __name__(__zero__)

/// {{if false}}
type (
	__name____goo_comment_bool__ bool
	__name____goo_comment_int__  int
) /// {{end}}

// {{.name}} is a {{.prim}}.
type __name__ __prim__

/// {{if false}}
type (
	__prim__   int
	Bool       bool
	Comparable interface{}
	Equatable  interface{}
	Integer    interface{}
	Number     interface{}
	String     string
) /// {{end}}

/// {{if .number}}
// Add implements Number.
func (__rec__ __name__) Add(n Number) Number {
	return __rec__ + n.(__name__)
} /// {{end}}

/// {{if eq .prim "bool"}}
// And returns the conjunction of __rec__ and other.
func (__rec__ __name____goo_comment_bool__) And__Goo_Comment_Bool__(other __name____goo_comment_bool__) __name____goo_comment_bool__ {
	return __rec__ && other
} /// {{end}}

/// {{if .integer}}
// And returns the bitwise conjunction of __rec__ and other.
func (__rec__ __name____goo_comment_int__) And__Goo_Comment_Int__(other Integer) Integer {
	return __rec__ & other.(__name____goo_comment_int__)
} /// {{end}}

/// {{if eq .prim "string"}}
// Concat returns the concatenation of {{.rec}} and other.
func (__rec__ __name__) Concat(other __name__) __name__ {
	return __rec__ + other
} /// {{end}}

/// {{if .number}}
// Divide implements Number.
func (__rec__ __name__) Divide(n Number) Number {
	return __rec__ / n.(__name__)
} /// {{end}}

/// {{if or .comparable .equatable .number}}
// Equals implements Equatable.
func (__rec__ __name__) Equals(e Equatable) bool {
	return __rec__ == e.(__name__)
} /// {{end}}

/// {{if .comparable}}
// Greater implements Comparable.
func (__rec__ __name__) Greater(c Comparable) bool {
	return __rec__ > c.(__name__)
}

// GreaterEqual implements Comparable.
func (__rec__ __name__) GreaterEqual(c Comparable) bool {
	return __rec__ >= c.(__name__)
} /// {{end}}

/// {{if eq .prim "complex64" "complex128"}}
// Imag returns the imaginary part.
func (__rec__ __name__) Imag() /** {{if eq .prim "complex64"}} float32 {{else}} float64 {{end}} **/ {
	/// return imag(__rec__)
} /// {{end}}

/// {{if .comparable}}
// Less implements Comparable.
func (__rec__ __name__) Less(c Comparable) bool {
	return __rec__ < c.(__name__)
}

// LessEqual implements Comparable.
func (__rec__ __name__) LessEqual(c Comparable) bool {
	return __rec__ <= c.(__name__)
} /// {{end}}

/// {{if .integer}}
// Modulo implements Number.
func (__rec__ __name__) Modulo(other Integer) Integer {
	return __rec__ % other.(__name__)
} /// {{end}}

/// {{if .number}}
// Multiply implements Number.
func (__rec__ __name__) Multiply(n Number) Number {
	return __rec__ * n.(__name__)
}

// Negate implements Number.
func (__rec__ __name__) Negate() Number {
	return -__rec__
} /// {{end}}

/// {{if eq .prim "bool"}}
// Not returns the negation of b.
func (b Bool) Not() Bool {
	return !b
} /// {{end}}

/// {{if or .comparable .equatable .number}}
// NotEquals implements Equatable.
func (__rec__ __name__) NotEquals(e Equatable) bool {
	return __rec__ != e.(__name__)
} /// {{end}}

/// {{if eq .prim "bool"}}
// Or returns the disjunction of b and other.
func (b Bool) Or(other Bool) Bool {
	return b || other
} /// {{end}}

/// {{if eq .prim "complex64" "complex128"}}
// Real returns the real part.
func (__rec__ __name__) Real() /** {{if eq .prim "complex64"}} float32 {{else}} float64 {{end}} **/ {
	/// return real(__rec__)
} /// {{end}}

/// {{if .number}}
// Subtract implements Number.
func (__rec__ __name__) Subtract(n Number) Number {
	return __rec__ - n.(__name__)
} /// {{end}}
