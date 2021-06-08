package slice

import "reflect"

func SplitGeneric(slice, sep interface{}) interface{} {
	return SplitValue(reflect.ValueOf(slice), sep).Interface()
}

func SplitValue(slice reflect.Value, sep interface{}) reflect.Value {
	result := reflect.Zero(reflect.SliceOf(slice.Type()))
	length := slice.Len()
	start := 0
	for i := 0; i < length; i++ {
		if slice.Index(i).Interface() == sep {
			result = reflect.Append(result, slice.Slice(start, i))
			start = i + 1
		}
	}

	if length > 0 {
		result = reflect.Append(result, slice.Slice(start, length))
	}
	return result
}

func SplitInterface(slice []interface{}, sep interface{}) (result [][]interface{}) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}

func SplitString(slice []string, sep string) (result [][]string) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitInt(slice []int, sep int) (result [][]int) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitInt8(slice []int8, sep int8) (result [][]int8) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitInt16(slice []int16, sep int16) (result [][]int16) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitInt32(slice []int32, sep int32) (result [][]int32) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitInt64(slice []int64, sep int64) (result [][]int64) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}

func SplitUint(slice []uint, sep uint) (result [][]uint) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitUint8(slice []uint8, sep uint8) (result [][]uint8) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitUint16(slice []uint16, sep uint16) (result [][]uint16) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitUint32(slice []uint32, sep uint32) (result [][]uint32) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
func SplitUint64(slice []uint64, sep uint64) (result [][]uint64) {
	start := 0
	for i, v := range slice {
		if v == sep {
			result = append(result, slice[start:i])
			start = i + 1
		}
	}

	if len(slice) > 0 {
		result = append(result, slice[start:len(slice)])
	}
	return
}
