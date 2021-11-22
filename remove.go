package slice

import "reflect"

func Remove(slice interface{}, fun func(int) bool) interface{} {
	sliceValue := reflect.ValueOf(slice)
	length := sliceValue.Len()
	j := 0
	for i := 0; i < length; i++ {
		if !fun(i) {
			if i != j {
				sliceValue.Index(j).Set(sliceValue.Index(i))
			}
			j++
		}
	}
	return sliceValue.Slice(0, j).Interface()
}

func RemoveGeneric(slice interface{}, targets ...interface{}) interface{} {
	if len(targets) == 0 {
		return slice
	}
	return RemoveValue(reflect.ValueOf(slice), targets...)
}

func RemoveValue(slice reflect.Value, targets ...interface{}) interface{} {
	targetsValue := reflect.ValueOf(targets)
	length := slice.Len()
	j := 0
	for i := 0; i < length; i++ {
		v := slice.Index(i)
		if !ContainsValue(targetsValue, v.Interface()) {
			if i != j {
				slice.Index(j).Set(v)
			}
			j++
		}
	}
	return slice.Slice(0, j).Interface()
}

func RemoveInterface(slice []interface{}, targets ...interface{}) []interface{} {
	j := 0
	for i, v := range slice {
		if !ContainsInterface(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}

func RemoveString(slice []string, targets ...string) []string {
	j := 0
	for i, v := range slice {
		if !ContainsString(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}

func RemoveBool(slice []bool, targets ...bool) []bool {
	j := 0
	for i, v := range slice {
		if !ContainsBool(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}

func RemoveInt(slice []int, targets ...int) []int {
	j := 0
	for i, v := range slice {
		if !ContainsInt(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveInt8(slice []int8, targets ...int8) []int8 {
	j := 0
	for i, v := range slice {
		if !ContainsInt8(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveInt16(slice []int16, targets ...int16) []int16 {
	j := 0
	for i, v := range slice {
		if !ContainsInt16(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveInt32(slice []int32, targets ...int32) []int32 {
	j := 0
	for i, v := range slice {
		if !ContainsInt32(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveInt64(slice []int64, targets ...int64) []int64 {
	j := 0
	for i, v := range slice {
		if !ContainsInt64(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveUint(slice []uint, targets ...uint) []uint {
	j := 0
	for i, v := range slice {
		if !ContainsUint(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveUint8(slice []uint8, targets ...uint8) []uint8 {
	j := 0
	for i, v := range slice {
		if !ContainsUint8(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveUint16(slice []uint16, targets ...uint16) []uint16 {
	j := 0
	for i, v := range slice {
		if !ContainsUint16(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveUint32(slice []uint32, targets ...uint32) []uint32 {
	j := 0
	for i, v := range slice {
		if !ContainsUint32(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
func RemoveUint64(slice []uint64, targets ...uint64) []uint64 {
	j := 0
	for i, v := range slice {
		if !ContainsUint64(targets, v) {
			if i != j {
				slice[j] = v
			}
			j++
		}
	}
	return slice[:j]
}
