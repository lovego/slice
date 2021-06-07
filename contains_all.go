package slice

import "reflect"

func ContainsAll(slice interface{}, targets ...interface{}) bool {
	if slice == nil {
		return len(targets) == 0
	}
	return ContainsAllValue(reflect.ValueOf(slice), targets...)
}

func ContainsAllValue(slice reflect.Value, targets ...interface{}) bool {
	for _, target := range targets {
		if !ContainsValue(slice, target) {
			return false
		}
	}
	return true
}

func ContainsAllInterface(slice []interface{}, targets ...interface{}) bool {
	for _, target := range targets {
		if !ContainsInterface(slice, target) {
			return false
		}
	}
	return true
}

func ContainsAllString(slice []string, targets ...string) bool {
	for _, target := range targets {
		if !ContainsString(slice, target) {
			return false
		}
	}
	return true
}

func ContainsAllInt(slice []int, targets ...int) bool {
	for _, target := range targets {
		if !ContainsInt(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllInt8(slice []int8, targets ...int8) bool {
	for _, target := range targets {
		if !ContainsInt8(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllInt16(slice []int16, targets ...int16) bool {
	for _, target := range targets {
		if !ContainsInt16(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllInt32(slice []int32, targets ...int32) bool {
	for _, target := range targets {
		if !ContainsInt32(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllInt64(slice []int64, targets ...int64) bool {
	for _, target := range targets {
		if !ContainsInt64(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllUint(slice []uint, targets ...uint) bool {
	for _, target := range targets {
		if !ContainsUint(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllUint8(slice []uint8, targets ...uint8) bool {
	for _, target := range targets {
		if !ContainsUint8(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllUint16(slice []uint16, targets ...uint16) bool {
	for _, target := range targets {
		if !ContainsUint16(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllUint32(slice []uint32, targets ...uint32) bool {
	for _, target := range targets {
		if !ContainsUint32(slice, target) {
			return false
		}
	}
	return true
}
func ContainsAllUint64(slice []uint64, targets ...uint64) bool {
	for _, target := range targets {
		if !ContainsUint64(slice, target) {
			return false
		}
	}
	return true
}
