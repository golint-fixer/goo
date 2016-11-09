//go:generate go get github.com/willfaught/goo/cmd/goo
//go:generate goo macro maps.go ../../map_bool_int.go -d {"Key":"Bool","KeyType":"bool","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_bool_string.go -d {"Key":"Bool","KeyType":"bool","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_bool_struct.go -d {"Key":"Bool","KeyType":"bool","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_byte_int.go -d {"Key":"Byte","KeyType":"byte","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_byte_string.go -d {"Key":"Byte","KeyType":"byte","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_byte_struct.go -d {"Key":"Byte","KeyType":"byte","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_complex128_int.go -d {"Key":"Complex128","KeyType":"complex128","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_complex128_string.go -d {"Key":"Complex128","KeyType":"complex128","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_complex128_struct.go -d {"Key":"Complex128","KeyType":"complex128","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_complex64_int.go -d {"Key":"Complex64","KeyType":"complex64","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_complex64_string.go -d {"Key":"Complex64","KeyType":"complex64","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_complex64_struct.go -d {"Key":"Complex64","KeyType":"complex64","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_float32_int.go -d {"Key":"Float32","KeyType":"float32","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_float32_string.go -d {"Key":"Float32","KeyType":"float32","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_float32_struct.go -d {"Key":"Float32","KeyType":"float32","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_float64_int.go -d {"Key":"Float64","KeyType":"float64","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_float64_string.go -d {"Key":"Float64","KeyType":"float64","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_float64_struct.go -d {"Key":"Float64","KeyType":"float64","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_int16_int.go -d {"Key":"Int16","KeyType":"int16","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_int16_string.go -d {"Key":"Int16","KeyType":"int16","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_int16_struct.go -d {"Key":"Int16","KeyType":"int16","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_int32_int.go -d {"Key":"Int32","KeyType":"int32","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_int32_string.go -d {"Key":"Int32","KeyType":"int32","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_int32_struct.go -d {"Key":"Int32","KeyType":"int32","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_int64_int.go -d {"Key":"Int64","KeyType":"int64","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_int64_string.go -d {"Key":"Int64","KeyType":"int64","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_int64_struct.go -d {"Key":"Int64","KeyType":"int64","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_int8_int.go -d {"Key":"Int8","KeyType":"int8","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_int8_string.go -d {"Key":"Int8","KeyType":"int8","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_int8_struct.go -d {"Key":"Int8","KeyType":"int8","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_int_bool.go -d {"Key":"Int","KeyType":"int","Value":"Bool","ValueType":"bool"}
//go:generate goo macro maps.go ../../map_int_byte.go -d {"Key":"Int","KeyType":"int","Value":"Byte","ValueType":"byte"}
//go:generate goo macro maps.go ../../map_int_complex128.go -d {"Key":"Int","KeyType":"int","Value":"Complex128","ValueType":"complex128"}
//go:generate goo macro maps.go ../../map_int_complex64.go -d {"Key":"Int","KeyType":"int","Value":"Complex64","ValueType":"complex64"}
//go:generate goo macro maps.go ../../map_int_float32.go -d {"Key":"Int","KeyType":"int","Value":"Float32","ValueType":"float32"}
//go:generate goo macro maps.go ../../map_int_float64.go -d {"Key":"Int","KeyType":"int","Value":"Float64","ValueType":"float64"}
//go:generate goo macro maps.go ../../map_int_int.go -d {"Key":"Int","KeyType":"int","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_int_int16.go -d {"Key":"Int","KeyType":"int","Value":"Int16","ValueType":"int16"}
//go:generate goo macro maps.go ../../map_int_int32.go -d {"Key":"Int","KeyType":"int","Value":"Int32","ValueType":"int32"}
//go:generate goo macro maps.go ../../map_int_int64.go -d {"Key":"Int","KeyType":"int","Value":"Int64","ValueType":"int64"}
//go:generate goo macro maps.go ../../map_int_int8.go -d {"Key":"Int","KeyType":"int","Value":"Int8","ValueType":"int8"}
//go:generate goo macro maps.go ../../map_int_interface.go -d {"Key":"Int","KeyType":"int","Value":"Interface","ValueType":"interface{}"}
//go:generate goo macro maps.go ../../map_int_rune.go -d {"Key":"Int","KeyType":"int","Value":"Rune","ValueType":"rune"}
//go:generate goo macro maps.go ../../map_int_string.go -d {"Key":"Int","KeyType":"int","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_int_struct.go -d {"Key":"Int","KeyType":"int","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_int_uint.go -d {"Key":"Int","KeyType":"int","Value":"Uint","ValueType":"uint"}
//go:generate goo macro maps.go ../../map_int_uint16.go -d {"Key":"Int","KeyType":"int","Value":"Uint16","ValueType":"uint16"}
//go:generate goo macro maps.go ../../map_int_uint32.go -d {"Key":"Int","KeyType":"int","Value":"Uint32","ValueType":"uint32"}
//go:generate goo macro maps.go ../../map_int_uint64.go -d {"Key":"Int","KeyType":"int","Value":"Uint64","ValueType":"uint64"}
//go:generate goo macro maps.go ../../map_int_uint8.go -d {"Key":"Int","KeyType":"int","Value":"Uint8","ValueType":"uint8"}
//go:generate goo macro maps.go ../../map_int_uintptr.go -d {"Key":"Int","KeyType":"int","Value":"Uintptr","ValueType":"uintptr"}
//go:generate goo macro maps.go ../../map_interface_bool.go -d {"Key":"Interface","KeyType":"interface{}","Value":"Bool","ValueType":"bool"}
//go:generate goo macro maps.go ../../map_interface_int.go -d {"Key":"Interface","KeyType":"interface{}","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_interface_interface.go -d {"Key":"Interface","KeyType":"interface{}","Value":"Interface","ValueType":"interface{}"}
//go:generate goo macro maps.go ../../map_interface_string.go -d {"Key":"Interface","KeyType":"interface{}","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_interface_struct.go -d {"Key":"Interface","KeyType":"interface{}","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_rune_int.go -d {"Key":"Rune","KeyType":"rune","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_rune_string.go -d {"Key":"Rune","KeyType":"rune","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_rune_struct.go -d {"Key":"Rune","KeyType":"rune","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_string_bool.go -d {"Key":"String","KeyType":"string","Value":"Bool","ValueType":"bool"}
//go:generate goo macro maps.go ../../map_string_byte.go -d {"Key":"String","KeyType":"string","Value":"Byte","ValueType":"byte"}
//go:generate goo macro maps.go ../../map_string_complex128.go -d {"Key":"String","KeyType":"string","Value":"Complex128","ValueType":"complex128"}
//go:generate goo macro maps.go ../../map_string_complex64.go -d {"Key":"String","KeyType":"string","Value":"Complex64","ValueType":"complex64"}
//go:generate goo macro maps.go ../../map_string_float32.go -d {"Key":"String","KeyType":"string","Value":"Float32","ValueType":"float32"}
//go:generate goo macro maps.go ../../map_string_float64.go -d {"Key":"String","KeyType":"string","Value":"Float64","ValueType":"float64"}
//go:generate goo macro maps.go ../../map_string_int.go -d {"Key":"String","KeyType":"string","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_string_int16.go -d {"Key":"String","KeyType":"string","Value":"Int16","ValueType":"int16"}
//go:generate goo macro maps.go ../../map_string_int32.go -d {"Key":"String","KeyType":"string","Value":"Int32","ValueType":"int32"}
//go:generate goo macro maps.go ../../map_string_int64.go -d {"Key":"String","KeyType":"string","Value":"Int64","ValueType":"int64"}
//go:generate goo macro maps.go ../../map_string_int8.go -d {"Key":"String","KeyType":"string","Value":"Int8","ValueType":"int8"}
//go:generate goo macro maps.go ../../map_string_interface.go -d {"Key":"String","KeyType":"string","Value":"Interface","ValueType":"interface{}"}
//go:generate goo macro maps.go ../../map_string_rune.go -d {"Key":"String","KeyType":"string","Value":"Rune","ValueType":"rune"}
//go:generate goo macro maps.go ../../map_string_string.go -d {"Key":"String","KeyType":"string","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_string_struct.go -d {"Key":"String","KeyType":"string","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_string_uint.go -d {"Key":"String","KeyType":"string","Value":"Uint","ValueType":"uint"}
//go:generate goo macro maps.go ../../map_string_uint16.go -d {"Key":"String","KeyType":"string","Value":"Uint16","ValueType":"uint16"}
//go:generate goo macro maps.go ../../map_string_uint32.go -d {"Key":"String","KeyType":"string","Value":"Uint32","ValueType":"uint32"}
//go:generate goo macro maps.go ../../map_string_uint64.go -d {"Key":"String","KeyType":"string","Value":"Uint64","ValueType":"uint64"}
//go:generate goo macro maps.go ../../map_string_uint8.go -d {"Key":"String","KeyType":"string","Value":"Uint8","ValueType":"uint8"}
//go:generate goo macro maps.go ../../map_string_uintptr.go -d {"Key":"String","KeyType":"string","Value":"Uintptr","ValueType":"uintptr"}
//go:generate goo macro maps.go ../../map_uint16_int.go -d {"Key":"Uint16","KeyType":"uint16","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_uint16_string.go -d {"Key":"Uint16","KeyType":"uint16","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_uint16_struct.go -d {"Key":"Uint16","KeyType":"uint16","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_uint32_int.go -d {"Key":"Uint32","KeyType":"uint32","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_uint32_string.go -d {"Key":"Uint32","KeyType":"uint32","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_uint32_struct.go -d {"Key":"Uint32","KeyType":"uint32","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_uint64_int.go -d {"Key":"Uint64","KeyType":"uint64","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_uint64_string.go -d {"Key":"Uint64","KeyType":"uint64","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_uint64_struct.go -d {"Key":"Uint64","KeyType":"uint64","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_uint8_int.go -d {"Key":"Uint8","KeyType":"uint8","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_uint8_string.go -d {"Key":"Uint8","KeyType":"uint8","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_uint8_struct.go -d {"Key":"Uint8","KeyType":"uint8","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_uint_int.go -d {"Key":"Uint","KeyType":"uint","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_uint_string.go -d {"Key":"Uint","KeyType":"uint","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_uint_struct.go -d {"Key":"Uint","KeyType":"uint","Value":"Struct","ValueType":"struct{}"}
//go:generate goo macro maps.go ../../map_uintptr_int.go -d {"Key":"Uintptr","KeyType":"uintptr","Value":"Int","ValueType":"int"}
//go:generate goo macro maps.go ../../map_uintptr_string.go -d {"Key":"Uintptr","KeyType":"uintptr","Value":"String","ValueType":"string"}
//go:generate goo macro maps.go ../../map_uintptr_struct.go -d {"Key":"Uintptr","KeyType":"uintptr","Value":"Struct","ValueType":"struct{}"}

package goo

/// {{if false}}
type __FIELD_KeyType__ interface{}
type __FIELD_ValueType__ interface{}
type Equatable interface{}
type Map interface{} /// {{end}}

var _ Map = Map__FIELD_Key____FIELD_Value__(nil)

// Map__FIELD_Key____FIELD_Value__ is a map from __FIELD_KeyType__ to __FIELD_ValueType__.
type Map__FIELD_Key____FIELD_Value__ map[__FIELD_KeyType__]__FIELD_ValueType__

// Delete implements Map.
func (m Map__FIELD_Key____FIELD_Value__) Delete(k interface{}) {
	delete(m, k.(__FIELD_KeyType__))
}

// Equals implements Map.
func (m Map__FIELD_Key____FIELD_Value__) Equals(other Equatable) bool {
	var n = other.(Map__FIELD_Key____FIELD_Value__)

	if len(n) != len(m) {
		return false
	}

	for k, v := range m {
		if nv, ok := n[k]; !ok {
			return false
		} else if nv != v {
			return false
		}
	}

	return true
}

// Get implements Map.
func (m Map__FIELD_Key____FIELD_Value__) Get(k interface{}) interface{} {
	return m[k.(__FIELD_KeyType__)]
}

// GetCheck implements Map.
func (m Map__FIELD_Key____FIELD_Value__) GetCheck(k interface{}) (interface{}, bool) {
	var v, ok = m[k.(__FIELD_KeyType__)]

	return v, ok
}

// KeyValues implements Map.
func (m Map__FIELD_Key____FIELD_Value__) KeyValues() [][2]interface{} {
	var kvs [][2]interface{}

	for k, v := range m {
		kvs = append(kvs, [2]interface{}{k, v})
	}

	return kvs
}

// Keys implements Map.
func (m Map__FIELD_Key____FIELD_Value__) Keys() []interface{} {
	var ks []interface{}

	for k := range m {
		ks = append(ks, k)
	}

	return ks
}

// Len implements Map.
func (m Map__FIELD_Key____FIELD_Value__) Len() int {
	return len(m)
}

// Make implements Map.
func (m Map__FIELD_Key____FIELD_Value__) Make(c int) Map {
	return make(Map__FIELD_Key____FIELD_Value__, c)
}

// NotEquals implements Map.
func (m Map__FIELD_Key____FIELD_Value__) NotEquals(other Equatable) bool {
	return !m.Equals(other)
}

// Set implements Map.
func (m Map__FIELD_Key____FIELD_Value__) Set(k, v interface{}) {
	m[k.(__FIELD_KeyType__)] = v.(__FIELD_ValueType__)
}
