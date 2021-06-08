package slice

import (
	"reflect"
)

// UnionGeneric returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionGeneric(left, right interface{}) interface{} {
	if left == nil && right == nil {
		return nil
	}
	return UnionValue(reflect.ValueOf(left), reflect.ValueOf(right)).Interface()
}

// UnionValue returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionValue(left, right reflect.Value) reflect.Value {
	var result reflect.Value
	if left.IsValid() {
		length := left.Len()
		result = reflect.MakeSlice(left.Type(), length, length)
		reflect.Copy(result, left)
	} else if right.IsValid() {
		result = reflect.Zero(right.Type())
	} else {
		return result
	}

	if right.IsValid() {
		length := right.Len()
		for i := 0; i < length; i++ {
			v := right.Index(i)
			if !ContainsValue(result, v.Interface()) {
				result = reflect.Append(result, v)
			}
		}
	}
	return result
}

// UnionInterface returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionInterface(left, right []interface{}) []interface{} {
	result := make([]interface{}, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsInterface(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionString returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionString(left, right []string) []string {
	result := make([]string, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsString(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionInt returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionInt(left, right []int) []int {
	result := make([]int, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsInt(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionInt8 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionInt8(left, right []int8) []int8 {
	result := make([]int8, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsInt8(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionInt16 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionInt16(left, right []int16) []int16 {
	result := make([]int16, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsInt16(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionInt32 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionInt32(left, right []int32) []int32 {
	result := make([]int32, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsInt32(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionInt64 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionInt64(left, right []int64) []int64 {
	result := make([]int64, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsInt64(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionUint returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionUint(left, right []uint) []uint {
	result := make([]uint, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsUint(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionUint8 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionUint8(left, right []uint8) []uint8 {
	result := make([]uint8, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsUint8(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionUint16 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionUint16(left, right []uint16) []uint16 {
	result := make([]uint16, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsUint16(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionUint32 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionUint32(left, right []uint32) []uint32 {
	result := make([]uint32, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsUint32(result, v) {
			result = append(result, v)
		}
	}
	return result
}

// UnionUint64 returns union set of left and right, with right follows left.
// The duplicate members in left are kept.
func UnionUint64(left, right []uint64) []uint64 {
	result := make([]uint64, len(left))
	copy(result, left)
	for _, v := range right {
		if !ContainsUint64(result, v) {
			result = append(result, v)
		}
	}
	return result
}
