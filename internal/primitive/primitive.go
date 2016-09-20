//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo -in primitive.go -out ../../primitive_bool.go -json {"Equatable":true,"Name":"Bool","Rec":"b","Type":"bool","Zero":"false"}
//go:generate goo -in primitive.go -out ../../primitive_byte.go -json {"Comparable":true,"Int":true,"Name":"Byte","Number":true,"Rec":"b","Type":"byte","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_complex64.go -json {"Name":"Complex64","Number":true,"Rec":"c","Type":"complex64","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_complex128.go -json {"Name":"Complex128","Number":true,"Rec":"c","Type":"complex128","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_float32.go -json {"Comparable":true,"Name":"Float32","Number":true,"Rec":"f","Type":"float32","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_float64.go -json {"Comparable":true,"Name":"Float64","Number":true,"Rec":"f","Type":"float64","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int.go -json {"Comparable":true,"Int":true,"Name":"Int","Number":true,"Rec":"i","Type":"int","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int8.go -json {"Comparable":true,"Int":true,"Name":"Int8","Number":true,"Rec":"i","Type":"int8","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int16.go -json {"Comparable":true,"Int":true,"Name":"Int16","Number":true,"Rec":"i","Type":"int16","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int32.go -json {"Comparable":true,"Int":true,"Name":"Int32","Number":true,"Rec":"i","Type":"int32","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_int64.go -json {"Comparable":true,"Int":true,"Name":"Int64","Number":true,"Rec":"i","Type":"int64","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_rune.go -json {"Comparable":true,"Int":true,"Name":"Rune","Number":true,"Rec":"r","Type":"rune","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_string.go -json {"Comparable":true,"Name":"String","Rec":"s","Type":"string","Zero":"\"\""}
//go:generate goo -in primitive.go -out ../../primitive_uint.go -json {"Comparable":true,"Int":true,"Name":"Uint","Number":true,"Rec":"u","Type":"uint","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uintptr.go -json {"Comparable":true,"Int":true,"Name":"Uintptr","Number":true,"Rec":"u","Type":"uintptr","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint8.go -json {"Comparable":true,"Int":true,"Name":"Uint8","Number":true,"Rec":"u","Type":"uint8","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint16.go -json {"Comparable":true,"Int":true,"Name":"Uint16","Number":true,"Rec":"u","Type":"uint16","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint32.go -json {"Comparable":true,"Int":true,"Name":"Uint32","Number":true,"Rec":"u","Type":"uint32","Zero":"0"}
//go:generate goo -in primitive.go -out ../../primitive_uint64.go -json {"Comparable":true,"Int":true,"Name":"Uint64","Number":true,"Rec":"u","Type":"uint64","Zero":"0"}

package goo

/// {{if false}}
var __Zero__ = 0 /// {{end}}

/// {{if .Int}}
var _ Integer = __Name__(__Zero__)

/// {{else if .Number}}
var _ Number = __Name__(__Zero__)

/// {{else if .Comparable}}
var _ Comparable = __Name__(__Zero__)

/// {{else if .Equatable}}
var _ Equatable = __Name__(__Zero__) /// {{end}}

// {{.Name}}Zero is the {{.Name}} zero value.
var __Name__Zero = __Name__(__Zero__)

/// {{if false}}
type (
	__Name____Goo_Comment_Bool__ bool
	__Name____Goo_Comment_Int__  int
) /// {{end}}

// {{.Name}} is a {{.Type}}.
type __Name__ __Type__

/// {{if false}}
type (
	__Type__   int
	Bool       bool
	Comparable interface{}
	Equatable  interface{}
	Integer    interface{}
	Number     interface{}
	String     string
) /// {{end}}

/// {{if .Number}}
// Add implements Number.
func (__Rec__ __Name__) Add(n Number) Number {
	return __Rec__ + n.(__Name__)
} /// {{end}}

/// {{if eq .Type "bool"}}
// And returns the conjunction of __Rec__ and other.
func (__Rec__ __Name____Goo_Comment_Bool__) And__Goo_Comment_Bool__(other __Name____Goo_Comment_Bool__) __Name____Goo_Comment_Bool__ {
	return __Rec__ && other
} /// {{end}}

/// {{if .Int}}
// And returns the bitwise conjunction of __Rec__ and other.
func (__Rec__ __Name____Goo_Comment_Int__) And__Goo_Comment_Int__(other Integer) Integer {
	return __Rec__ & other.(__Name____Goo_Comment_Int__)
} /// {{end}}

/// {{if eq .Type "string"}}
// Concat returns the concatenation of {{.Rec}} and other.
func (__Rec__ __Name__) Concat(other __Name__) __Name__ {
	return __Rec__ + other
} /// {{end}}

/// {{if .Number}}
// Divide implements Number.
func (__Rec__ __Name__) Divide(n Number) Number {
	return __Rec__ / n.(__Name__)
} /// {{end}}

/// {{if or .Comparable .Equatable .Number}}
// Equals implements Equatable.
func (__Rec__ __Name__) Equals(e Equatable) bool {
	return __Rec__ == e.(__Name__)
} /// {{end}}

/// {{if .Comparable}}
// Greater implements Comparable.
func (__Rec__ __Name__) Greater(c Comparable) bool {
	return __Rec__ > c.(__Name__)
}

// GreaterEqual implements Comparable.
func (__Rec__ __Name__) GreaterEqual(c Comparable) bool {
	return __Rec__ >= c.(__Name__)
} /// {{end}}

/// {{if eq .Type "complex64" "complex128"}}
// Imag returns the imaginary part.
func (__Rec__ __Name__) Imag() /** {{if eq .Type "complex64"}} float32 {{else}} float64 {{end}} **/ {
	/// return imag(__Rec__)
} /// {{end}}

/// {{if .Comparable}}
// Less implements Comparable.
func (__Rec__ __Name__) Less(c Comparable) bool {
	return __Rec__ < c.(__Name__)
}

// LessEqual implements Comparable.
func (__Rec__ __Name__) LessEqual(c Comparable) bool {
	return __Rec__ <= c.(__Name__)
} /// {{end}}

/// {{if .Int}}
// Modulo implements Number.
func (__Rec__ __Name__) Modulo(other Integer) Integer {
	return __Rec__ % other.(__Name__)
} /// {{end}}

/// {{if .Number}}
// Multiply implements Number.
func (__Rec__ __Name__) Multiply(n Number) Number {
	return __Rec__ * n.(__Name__)
}

// Negate implements Number.
func (__Rec__ __Name__) Negate() Number {
	return -__Rec__
} /// {{end}}

/// {{if eq .Type "bool"}}
// Not returns the negation of b.
func (b Bool) Not() Bool {
	return !b
} /// {{end}}

/// {{if or .Comparable .Equatable .Number}}
// NotEquals implements Equatable.
func (__Rec__ __Name__) NotEquals(e Equatable) bool {
	return __Rec__ != e.(__Name__)
} /// {{end}}

/// {{if eq .Type "bool"}}
// Or returns the disjunction of b and other.
func (b Bool) Or(other Bool) Bool {
	return b || other
} /// {{end}}

/// {{if eq .Type "complex64" "complex128"}}
// Real returns the real part.
func (__Rec__ __Name__) Real() /** {{if eq .Type "complex64"}} float32 {{else}} float64 {{end}} **/ {
	/// return real(__Rec__)
} /// {{end}}

/// {{if .Number}}
// Subtract implements Number.
func (__Rec__ __Name__) Subtract(n Number) Number {
	return __Rec__ - n.(__Name__)
} /// {{end}}
