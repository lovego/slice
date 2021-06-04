package slice

import "reflect"

func ContainsAny(slice interface{}, targets ...interface{}) bool {
	if slice == nil {
		return false
	}
	sliceValue := reflect.ValueOf(slice)
	if sliceValue.Len() == 0 {
		return false
	}
	for _, target := range targets {
		if index(sliceValue, target) >= 0 {
			return true
		}
	}
	return false
}

func ContainsAnyValue(slice reflect.Value, targets ...reflect.Value) bool {
	if slice.Len() == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsValue(slice, target) {
			return true
		}
	}
	return false
}

func ContainsAnyInterface(slice []interface{}, targets ...interface{}) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsInterface(slice, target) {
			return true
		}
	}
	return false
}

func ContainsAnyString(slice []string, targets ...string) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsString(slice, target) {
			return true
		}
	}
	return false
}

func ContainsAnyInt(slice []int, targets ...int) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsInt(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyInt8(slice []int8, targets ...int8) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsInt8(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyInt16(slice []int16, targets ...int16) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsInt16(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyInt32(slice []int32, targets ...int32) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsInt32(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyInt64(slice []int64, targets ...int64) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsInt64(slice, target) {
			return true
		}
	}
	return false
}

func ContainsAnyUint(slice []uint, targets ...uint) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsUint(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyUint8(slice []uint8, targets ...uint8) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsUint8(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyUint16(slice []uint16, targets ...uint16) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsUint16(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyUint32(slice []uint32, targets ...uint32) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsUint32(slice, target) {
			return true
		}
	}
	return false
}
func ContainsAnyUint64(slice []uint64, targets ...uint64) bool {
	if len(slice) == 0 {
		return false
	}
	for _, target := range targets {
		if ContainsUint64(slice, target) {
			return true
		}
	}
	return false
}
