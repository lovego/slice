package slice

import "reflect"

func Contains(slice interface{}, target interface{}) bool {
	return Index(slice, target) >= 0
}

func ContainsValue(slice reflect.Value, target interface{}) bool {
	return IndexValue(slice, target) >= 0
}

func ContainsInterface(slice []interface{}, target interface{}) bool {
	return IndexInterface(slice, target) >= 0
}

func ContainsString(slice []string, target string) bool {
	return IndexString(slice, target) >= 0
}

func ContainsInt(slice []int, target int) bool {
	return IndexInt(slice, target) >= 0
}

func ContainsInt8(slice []int8, target int8) bool {
	return IndexInt8(slice, target) >= 0
}

func ContainsInt16(slice []int16, target int16) bool {
	return IndexInt16(slice, target) >= 0
}

func ContainsInt32(slice []int32, target int32) bool {
	return IndexInt32(slice, target) >= 0
}

func ContainsInt64(slice []int64, target int64) bool {
	return IndexInt64(slice, target) >= 0
}

func ContainsUint(slice []uint, target uint) bool {
	return IndexUint(slice, target) >= 0
}

func ContainsUint8(slice []uint8, target uint8) bool {
	return IndexUint8(slice, target) >= 0
}

func ContainsUint16(slice []uint16, target uint16) bool {
	return IndexUint16(slice, target) >= 0
}

func ContainsUint32(slice []uint32, target uint32) bool {
	return IndexUint32(slice, target) >= 0
}

func ContainsUint64(slice []uint64, target uint64) bool {
	return IndexUint64(slice, target) >= 0
}
