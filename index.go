package slice

import "reflect"

func Index(length int, fun func(int) bool) int {
	for i := 0; i < length; i++ {
		if fun(i) {
			return i
		}
	}
	return -1
}

func IndexGeneric(slice interface{}, target interface{}) int {
	if slice == nil {
		return -1
	}
	return IndexValue(reflect.ValueOf(slice), target)
}

func IndexValue(slice reflect.Value, target interface{}) int {
	length := slice.Len()
	for i := 0; i < length; i++ {
		if slice.Index(i).Interface() == target {
			return i
		}
	}
	return -1
}

func IndexInterface(slice []interface{}, target interface{}) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

func IndexString(slice []string, target string) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

func IndexInt(slice []int, target int) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexInt8(slice []int8, target int8) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexInt16(slice []int16, target int16) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexInt32(slice []int32, target int32) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexInt64(slice []int64, target int64) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

func IndexUint(slice []uint, target uint) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexUint8(slice []uint8, target uint8) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexUint16(slice []uint16, target uint16) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexUint32(slice []uint32, target uint32) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func IndexUint64(slice []uint64, target uint64) int {
	for i := 0; i < len(slice); i++ {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
