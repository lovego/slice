package slice

import "reflect"

// Substract substracts right from left.
func SubstractGeneric(left, right interface{}) interface{} {
	if left == nil {
		return nil
	}
	return SubstractValue(reflect.ValueOf(left), reflect.ValueOf(right)).Interface()
}

// SubstractValue substracts right from left.
func SubstractValue(left, right reflect.Value) reflect.Value {
	var result = reflect.Zero(left.Type())
	length := left.Len()
	if length == 0 {
		return result
	}
	for i := 0; i < length; i++ {
		v := left.Index(i)
		if !right.IsValid() || !ContainsValue(right, v.Interface()) {
			result = reflect.Append(result, v)
		}
	}
	return result
}

// SubstractInterface substracts right from left.
func SubstractInterface(left, right []interface{}) []interface{} {
	var result []interface{}
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsInterface(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractString substracts right from left.
func SubstractString(left, right []string) []string {
	var result []string
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsString(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractInt substracts right from left.
func SubstractInt(left, right []int) []int {
	var result []int
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsInt(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractInt8 substracts right from left.
func SubstractInt8(left, right []int8) []int8 {
	var result []int8
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsInt8(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractInt16 substracts right from left.
func SubstractInt16(left, right []int16) []int16 {
	var result []int16
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsInt16(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractInt32 substracts right from left.
func SubstractInt32(left, right []int32) []int32 {
	var result []int32
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsInt32(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractInt64 substracts right from left.
func SubstractInt64(left, right []int64) []int64 {
	var result []int64
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsInt64(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractUint substracts right from left.
func SubstractUint(left, right []uint) []uint {
	var result []uint
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsUint(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractUint8 substracts right from left.
func SubstractUint8(left, right []uint8) []uint8 {
	var result []uint8
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsUint8(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractUint16 substracts right from left.
func SubstractUint16(left, right []uint16) []uint16 {
	var result []uint16
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsUint16(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractUint32 substracts right from left.
func SubstractUint32(left, right []uint32) []uint32 {
	var result []uint32
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsUint32(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// SubstractUint64 substracts right from left.
func SubstractUint64(left, right []uint64) []uint64 {
	var result []uint64
	if len(left) == 0 {
		return result
	}
	for _, v := range left {
		if !ContainsUint64(right, v) {
			result = append(result, v)
		}
	}
	return result
}
