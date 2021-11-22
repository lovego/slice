package slice

import "reflect"

// IntersectGeneric returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectGeneric(left, right interface{}) interface{} {
	if left == nil || right == nil {
		return nil
	}
	return IntersectValue(reflect.ValueOf(left), reflect.ValueOf(right)).Interface()
}

// IntersectValue returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectValue(left, right reflect.Value) reflect.Value {
	var result = reflect.Zero(left.Type())
	length := left.Len()
	if length == 0 || !right.IsValid() || right.Len() == 0 {
		return result
	}
	for i := 0; i < length; i++ {
		v := left.Index(i)
		if ContainsValue(right, v.Interface()) {
			result = reflect.Append(result, v)
		}
	}
	return result
}

// IntersectInterface returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectInterface(left, right []interface{}) []interface{} {
	var result []interface{}
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsInterface(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectString returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectString(left, right []string) []string {
	var result []string
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsString(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectBool returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectBool(left, right []bool) []bool {
	var result []bool
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsBool(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectInt returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectInt(left, right []int) []int {
	var result []int
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsInt(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectInt8 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectInt8(left, right []int8) []int8 {
	var result []int8
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsInt8(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectInt16 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectInt16(left, right []int16) []int16 {
	var result []int16
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsInt16(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectInt32 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectInt32(left, right []int32) []int32 {
	var result []int32
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsInt32(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectInt64 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectInt64(left, right []int64) []int64 {
	var result []int64
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsInt64(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectUint returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectUint(left, right []uint) []uint {
	var result []uint
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsUint(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectUint8 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectUint8(left, right []uint8) []uint8 {
	var result []uint8
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsUint8(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectUint16 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectUint16(left, right []uint16) []uint16 {
	var result []uint16
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsUint16(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectUint32 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectUint32(left, right []uint32) []uint32 {
	var result []uint32
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsUint32(right, v) {
			result = append(result, v)
		}
	}
	return result
}

// IntersectUint64 returns intersection of left and right, in the left order.
// The duplicate members in left are kept.
func IntersectUint64(left, right []uint64) []uint64 {
	var result []uint64
	if len(left) == 0 || len(right) == 0 {
		return result
	}
	for _, v := range left {
		if ContainsUint64(right, v) {
			result = append(result, v)
		}
	}
	return result
}
