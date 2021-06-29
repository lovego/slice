package slice

import "reflect"

// Input slice will be modified.
func UniqueGeneric(slice interface{}) interface{} {
	return UniqueValue(reflect.ValueOf(slice)).Interface()
}

// Input slice will be modified.
func UniqueValue(slice reflect.Value) reflect.Value {
	if slice.Kind() == reflect.Ptr {
		slice = slice.Elem()
	}
	j := 0
	length := slice.Len()
	for i := 0; i < length; i++ {
		v := slice.Index(i)
		if !ContainsValue(slice.Slice(0, j), v.Interface()) {
			slice.Index(j).Set(v)
			j++
		}
	}
	return slice.Slice(0, j)
}

// Input slice will be modified.
func UniqueInterface(slice []interface{}) []interface{} {
	j := 0
	for _, v := range slice {
		if !ContainsInterface(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueString(slice []string) []string {
	j := 0
	for _, v := range slice {
		if !ContainsString(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueInt(slice []int) []int {
	j := 0
	for _, v := range slice {
		if !ContainsInt(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueInt8(slice []int8) []int8 {
	j := 0
	for _, v := range slice {
		if !ContainsInt8(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueInt16(slice []int16) []int16 {
	j := 0
	for _, v := range slice {
		if !ContainsInt16(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueInt32(slice []int32) []int32 {
	j := 0
	for _, v := range slice {
		if !ContainsInt32(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueInt64(slice []int64) []int64 {
	j := 0
	for _, v := range slice {
		if !ContainsInt64(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueUint(slice []uint) []uint {
	j := 0
	for _, v := range slice {
		if !ContainsUint(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueUint8(slice []uint8) []uint8 {
	j := 0
	for _, v := range slice {
		if !ContainsUint8(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueUint16(slice []uint16) []uint16 {
	j := 0
	for _, v := range slice {
		if !ContainsUint16(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueUint32(slice []uint32) []uint32 {
	j := 0
	for _, v := range slice {
		if !ContainsUint32(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}

// Input slice will be modified.
func UniqueUint64(slice []uint64) []uint64 {
	j := 0
	for _, v := range slice {
		if !ContainsUint64(slice[0:j], v) {
			slice[j] = v
			j++
		}
	}
	return slice[:j]
}
