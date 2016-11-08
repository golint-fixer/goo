//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro primitive.go -o ../../primitive_bool.go -d {"bool":true,"equatable":true,"name":"Bool","rec":"b","base":"bool","zero":"false"}
//go:generate goo macro primitive.go -o ../../primitive_byte.go -d {"comparable":true,"integer":true,"name":"Byte","number":true,"rec":"b","base":"byte","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_complex64.go -d {"name":"Complex64","number":true,"rec":"c","base":"complex64","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_complex128.go -d {"name":"Complex128","number":true,"rec":"c","base":"complex128","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_float32.go -d {"comparable":true,"name":"Float32","number":true,"rec":"f","base":"float32","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_float64.go -d {"comparable":true,"name":"Float64","number":true,"rec":"f","base":"float64","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_int.go -d {"comparable":true,"integer":true,"name":"Int","number":true,"rec":"i","base":"int","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_int8.go -d {"comparable":true,"integer":true,"name":"Int8","number":true,"rec":"i","base":"int8","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_int16.go -d {"comparable":true,"integer":true,"name":"Int16","number":true,"rec":"i","base":"int16","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_int32.go -d {"comparable":true,"integer":true,"name":"Int32","number":true,"rec":"i","base":"int32","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_int64.go -d {"comparable":true,"integer":true,"name":"Int64","number":true,"rec":"i","base":"int64","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_rune.go -d {"comparable":true,"integer":true,"name":"Rune","number":true,"rec":"r","base":"rune","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_string.go -d {"comparable":true,"name":"String","rec":"s","base":"string","zero":"\"\""}
//go:generate goo macro primitive.go -o ../../primitive_uint.go -d {"comparable":true,"integer":true,"name":"Uint","number":true,"rec":"u","base":"uint","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_uintptr.go -d {"comparable":true,"integer":true,"name":"Uintptr","number":true,"rec":"u","base":"uintptr","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_uint8.go -d {"comparable":true,"integer":true,"name":"Uint8","number":true,"rec":"u","base":"uint8","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_uint16.go -d {"comparable":true,"integer":true,"name":"Uint16","number":true,"rec":"u","base":"uint16","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_uint32.go -d {"comparable":true,"integer":true,"name":"Uint32","number":true,"rec":"u","base":"uint32","zero":"0"}
//go:generate goo macro primitive.go -o ../../primitive_uint64.go -d {"comparable":true,"integer":true,"name":"Uint64","number":true,"rec":"u","base":"uint64","zero":"0"}

package goo

/// {{if false}}
type (
	__FIELD_base__          int
	__FIELD_nameOMIT_bool__ bool
	__FIELD_nameOMIT_int__  int
	Bool                    bool
	Comparable              interface{}
	Complex                 interface{}
	Complex64               complex64
	Complex128              complex128
	Equatable               interface{}
	Integer                 interface{}
	Number                  interface{}
	String                  string
) /// {{end}}

/// {{if false}}
var __FIELD_zero__ = 0 /// {{end}}

// __FIELD_name__Zero is the __FIELD_name__ zero value.
var __FIELD_name__Zero = __FIELD_name__(__FIELD_zero__)

/// {{if .integer}}
var _ Integer = __FIELD_name__Zero    /// {{else if .number}}
var _ Number = __FIELD_name__Zero     /// {{else if .comparable}}
var _ Comparable = __FIELD_name__Zero /// {{else if .equatable}}
var _ Equatable = __FIELD_name__Zero  /// {{end}}

// {{.name}} is a {{.base}}.
type __FIELD_name__ __FIELD_base__

/// {{if .number}}
// Add implements Number.
func (__FIELD_rec__ __FIELD_name__) Add(n Number) Number {
	return __FIELD_rec__ + n.(__FIELD_name__)
} /// {{end}}

/// {{if .bool}}
// And returns the conjunction of __FIELD_rec__ and other.
func (__FIELD_rec__ Bool) And__X_OMIT_bool__(other Bool) Bool {
	return __FIELD_rec__ && other
} /// {{end}}

/// {{if .integer}}
// And implements Integer.
func (__FIELD_rec__ __FIELD_nameOMIT_int__) And__X_OMIT_int__(other Integer) Integer {
	return __FIELD_rec__ & other.(__FIELD_nameOMIT_int__)
}

// AndNot implements Integer.
func (__FIELD_rec__ __FIELD_name__) AndNot(other Integer) Integer {
	return __FIELD_rec__ &^ other.(__FIELD_name__)
} /// {{end}}

/// {{if eq .base "string"}}
// Concat returns the concatenation of {{.rec}} and other.
func (__FIELD_rec__ __FIELD_name__) Concat(other __FIELD_name__) __FIELD_name__ {
	return __FIELD_rec__ + other
} /// {{end}}

/// {{if .number}}
// Divide implements Number.
func (__FIELD_rec__ __FIELD_name__) Divide(other Number) Number {
	return __FIELD_rec__ / other.(__FIELD_name__)
} /// {{end}}

/// {{if or .comparable .equatable .number}}
// Equals implements Equatable.
func (__FIELD_rec__ __FIELD_name__) Equals(other Equatable) bool {
	return __FIELD_rec__ == other.(__FIELD_name__)
} /// {{end}}

/// {{if .comparable}}
// Greater implements Comparable.
func (__FIELD_rec__ __FIELD_name__) Greater(other Comparable) bool {
	return __FIELD_rec__ > other.(__FIELD_name__)
}

// GreaterEqual implements Comparable.
func (__FIELD_rec__ __FIELD_name__) GreaterEqual(other Comparable) bool {
	return __FIELD_rec__ >= other.(__FIELD_name__)
} /// {{end}}

/// {{if eq .base "complex64"}}
// Imag implements Complex.
func (__FIELD_rec__ Complex64) Imag() float32 {
	return imag(__FIELD_rec__)
} /// {{end}}

/// {{if eq .base "complex128"}}
// Imag implements Complex.
func (__FIELD_rec__ Complex128) Imag() float64 {
	return imag(__FIELD_rec__)
} /// {{end}}

/// {{if .integer}}
// Left implements Integer.
func (__FIELD_rec__ __FIELD_name__) Left(n uint) Integer {
	return __FIELD_rec__ << n
} /// {{end}}

/// {{if .comparable}}
// Less implements Comparable.
func (__FIELD_rec__ __FIELD_name__) Less(other Comparable) bool {
	return __FIELD_rec__ < other.(__FIELD_name__)
}

// LessEqual implements Comparable.
func (__FIELD_rec__ __FIELD_name__) LessEqual(other Comparable) bool {
	return __FIELD_rec__ <= other.(__FIELD_name__)
} /// {{end}}

/// {{if .integer}}
// Modulo implements Integer.
func (__FIELD_rec__ __FIELD_name__) Modulo(other Integer) Integer {
	return __FIELD_rec__ % other.(__FIELD_name__)
} /// {{end}}

/// {{if .number}}
// Multiply implements Number.
func (__FIELD_rec__ __FIELD_name__) Multiply(other Number) Number {
	return __FIELD_rec__ * other.(__FIELD_name__)
}

// Negate implements Number.
func (__FIELD_rec__ __FIELD_name__) Negate() Number {
	return -__FIELD_rec__
} /// {{end}}

/// {{if .bool}}
// Not returns the negation of __FIELD_rec__.
func (__FIELD_rec__ Bool) Not__X_OMIT_bool__() Bool {
	return !__FIELD_rec__
} /// {{end}}

/// {{if .integer}}
// Not implements Integer.
func (__FIELD_rec__ __FIELD_nameOMIT_int__) Not__X_OMIT_int__() Integer {
	return ^__FIELD_rec__
} /// {{end}}

/// {{if or .comparable .equatable .number}}
// NotEquals implements Equatable.
func (__FIELD_rec__ __FIELD_name__) NotEquals(other Equatable) bool {
	return __FIELD_rec__ != other.(__FIELD_name__)
} /// {{end}}

/// {{if .bool}}
// Or returns the disjunction of b and other.
func (__FIELD_rec__ Bool) Or__X_OMIT_bool__(other Bool) Bool {
	return __FIELD_rec__ || other
} /// {{end}}

/// {{if .integer}}
// Or implements Integer.
func (__FIELD_rec__ __FIELD_nameOMIT_int__) Or__X_OMIT_int__(other Integer) Integer {
	return __FIELD_rec__ | other.(__FIELD_nameOMIT_int__)
} /// {{end}}

/// {{if eq .base "complex64"}}
// Real implements Complex.
func (__FIELD_rec__ Complex64) Real__X_OMIT_complex64__() float32 {
	return real(__FIELD_rec__)
} /// {{end}}

/// {{if eq .base "complex128"}}
// Real implements Complex.
func (__FIELD_rec__ Complex128) Real_OMIT_complex128() float64 {
	return real(__FIELD_rec__)
} /// {{end}}

/// {{if .integer}}
// Right implements Integer.
func (__FIELD_rec__ __FIELD_name__) Right(n uint) Integer {
	return __FIELD_rec__ >> n
} /// {{end}}

/// {{if .number}}
// Subtract implements Number.
func (__FIELD_rec__ __FIELD_name__) Subtract(other Number) Number {
	return __FIELD_rec__ - other.(__FIELD_name__)
} /// {{end}}

/// {{if .integer}}
// Xor implements Integer.
func (__FIELD_rec__ __FIELD_name__) Xor(other Integer) Integer {
	return __FIELD_rec__ ^ other.(__FIELD_name__)
} /// {{end}}
