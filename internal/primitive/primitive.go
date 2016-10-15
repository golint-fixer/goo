//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro -i primitive.go -o ../../primitive_bool.go -j {"equatable":true,"name":"Bool","rec":"b","base":"bool","zero":"false"}
//go:generate goo macro -i primitive.go -o ../../primitive_byte.go -j {"comparable":true,"integer":true,"name":"Byte","number":true,"rec":"b","base":"byte","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_complex64.go -j {"name":"Complex64","number":true,"rec":"c","base":"complex64","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_complex128.go -j {"name":"Complex128","number":true,"rec":"c","base":"complex128","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_float32.go -j {"comparable":true,"name":"Float32","number":true,"rec":"f","base":"float32","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_float64.go -j {"comparable":true,"name":"Float64","number":true,"rec":"f","base":"float64","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_int.go -j {"comparable":true,"integer":true,"name":"Int","number":true,"rec":"i","base":"int","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_int8.go -j {"comparable":true,"integer":true,"name":"Int8","number":true,"rec":"i","base":"int8","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_int16.go -j {"comparable":true,"integer":true,"name":"Int16","number":true,"rec":"i","base":"int16","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_int32.go -j {"comparable":true,"integer":true,"name":"Int32","number":true,"rec":"i","base":"int32","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_int64.go -j {"comparable":true,"integer":true,"name":"Int64","number":true,"rec":"i","base":"int64","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_rune.go -j {"comparable":true,"integer":true,"name":"Rune","number":true,"rec":"r","base":"rune","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_string.go -j {"comparable":true,"name":"String","rec":"s","base":"string","zero":"\"\""}
//go:generate goo macro -i primitive.go -o ../../primitive_uint.go -j {"comparable":true,"integer":true,"name":"Uint","number":true,"rec":"u","base":"uint","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_uintptr.go -j {"comparable":true,"integer":true,"name":"Uintptr","number":true,"rec":"u","base":"uintptr","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_uint8.go -j {"comparable":true,"integer":true,"name":"Uint8","number":true,"rec":"u","base":"uint8","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_uint16.go -j {"comparable":true,"integer":true,"name":"Uint16","number":true,"rec":"u","base":"uint16","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_uint32.go -j {"comparable":true,"integer":true,"name":"Uint32","number":true,"rec":"u","base":"uint32","zero":"0"}
//go:generate goo macro -i primitive.go -o ../../primitive_uint64.go -j {"comparable":true,"integer":true,"name":"Uint64","number":true,"rec":"u","base":"uint64","zero":"0"}

package goo

/// {{if false}}
var __FIELD_zero__ = 0 /// {{end}}

/// {{if .integer}}
var _ Integer = __FIELD_name__(__FIELD_zero__)

/// {{else if .number}}
var _ Number = __FIELD_name__(__FIELD_zero__)

/// {{else if .comparable}}
var _ Comparable = __FIELD_name__(__FIELD_zero__)

/// {{else if .equatable}}
var _ Equatable = __FIELD_name__(__FIELD_zero__) /// {{end}}

// {{.name}}Zero is the {{.name}} zero value.
var __FIELD_name__Zero = __FIELD_name__(__FIELD_zero__)

/// {{if false}}
type (
	__X_IFE_eq_FIELD_base_STRING_complex64_THEN_float32_ELSE_float64_ENDIF__ interface{}

	__FIELD_nameOMIT_bool__ bool
	__FIELD_nameOMIT_imag__ complex64
	__FIELD_nameOMIT_int__  int
) /// {{end}}

// {{.name}} is a {{.base}}.
type __FIELD_name__ __FIELD_base__

/// {{if false}}
type (
	__FIELD_base__ int
	Bool           bool
	Comparable     interface{}
	Equatable      interface{}
	Integer        interface{}
	Number         interface{}
	String         string
) /// {{end}}

/// {{if .number}}
// Add implements Number.
func (__FIELD_rec__ __FIELD_name__) Add(n Number) Number {
	return __FIELD_rec__ + n.(__FIELD_name__)
} /// {{end}}

/// {{if eq .base "bool"}}
// And returns the conjunction of __FIELD_rec__ and other.
func (__FIELD_rec__ __FIELD_nameOMIT_bool__) And__X_OMIT_bool__(other __FIELD_nameOMIT_bool__) __FIELD_nameOMIT_bool__ {
	return __FIELD_rec__ && other
} /// {{end}}

/// {{if .integer}}
// And returns the bitwise conjunction of __FIELD_rec__ and other.
func (__FIELD_rec__ __FIELD_nameOMIT_int__) And__X_OMIT_int__(other Integer) Integer {
	return __FIELD_rec__ & other.(__FIELD_nameOMIT_int__)
} /// {{end}}

/// {{if eq .base "string"}}
// Concat returns the concatenation of {{.rec}} and other.
func (__FIELD_rec__ __FIELD_name__) Concat(other __FIELD_name__) __FIELD_name__ {
	return __FIELD_rec__ + other
} /// {{end}}

/// {{if .number}}
// Divide implements Number.
func (__FIELD_rec__ __FIELD_name__) Divide(n Number) Number {
	return __FIELD_rec__ / n.(__FIELD_name__)
} /// {{end}}

/// {{if or .comparable .equatable .number}}
// Equals implements Equatable.
func (__FIELD_rec__ __FIELD_name__) Equals(e Equatable) bool {
	return __FIELD_rec__ == e.(__FIELD_name__)
} /// {{end}}

/// {{if .comparable}}
// Greater implements Comparable.
func (__FIELD_rec__ __FIELD_name__) Greater(c Comparable) bool {
	return __FIELD_rec__ > c.(__FIELD_name__)
}

// GreaterEqual implements Comparable.
func (__FIELD_rec__ __FIELD_name__) GreaterEqual(c Comparable) bool {
	return __FIELD_rec__ >= c.(__FIELD_name__)
} /// {{end}}

/// {{if eq .base "complex64" "complex128"}}
// Imag returns the imaginary part.
func (__FIELD_rec__ __FIELD_nameOMIT_imag__) Imag() __X_IFE_eq_FIELD_base_STRING_complex64_THEN_float32_ELSE_float64_ENDIF__ {
	return imag(__FIELD_rec__)
} /// {{end}}

/// {{if .comparable}}
// Less implements Comparable.
func (__FIELD_rec__ __FIELD_name__) Less(c Comparable) bool {
	return __FIELD_rec__ < c.(__FIELD_name__)
}

// LessEqual implements Comparable.
func (__FIELD_rec__ __FIELD_name__) LessEqual(c Comparable) bool {
	return __FIELD_rec__ <= c.(__FIELD_name__)
} /// {{end}}

/// {{if .integer}}
// Modulo implements Number.
func (__FIELD_rec__ __FIELD_name__) Modulo(other Integer) Integer {
	return __FIELD_rec__ % other.(__FIELD_name__)
} /// {{end}}

/// {{if .number}}
// Multiply implements Number.
func (__FIELD_rec__ __FIELD_name__) Multiply(n Number) Number {
	return __FIELD_rec__ * n.(__FIELD_name__)
}

// Negate implements Number.
func (__FIELD_rec__ __FIELD_name__) Negate() Number {
	return -__FIELD_rec__
} /// {{end}}

/// {{if eq .base "bool"}}
// Not returns the negation of __FIELD_rec__.
func (__FIELD_rec__ Bool) Not() Bool {
	return !__FIELD_rec__
} /// {{end}}

/// {{if or .comparable .equatable .number}}
// NotEquals implements Equatable.
func (__FIELD_rec__ __FIELD_name__) NotEquals(e Equatable) bool {
	return __FIELD_rec__ != e.(__FIELD_name__)
} /// {{end}}

/// {{if eq .base "bool"}}
// Or returns the disjunction of b and other.
func (b Bool) Or(other Bool) Bool {
	return b || other
} /// {{end}}

/// {{if eq .base "complex64" "complex128"}}
// Real returns the real part.
func (__FIELD_rec__ __FIELD_nameOMIT_imag__) Real() __X_IFE_eq_FIELD_base_STRING_complex64_THEN_float32_ELSE_float64_ENDIF__ {
	return real(__FIELD_rec__)
} /// {{end}}

/// {{if .number}}
// Subtract implements Number.
func (__FIELD_rec__ __FIELD_name__) Subtract(n Number) Number {
	return __FIELD_rec__ - n.(__FIELD_name__)
} /// {{end}}
