package slice

import "reflect"

func LastIndex(slice interface{}, target interface{}) int {
	if slice == nil {
		return -1
	}
	return LastIndexValue(reflect.ValueOf(slice), target)
}

func LastIndexValue(slice reflect.Value, target interface{}) int {
	for i := slice.Len() - 1; i >= 0; i-- {
		if slice.Index(i).Interface() == target {
			return i
		}
	}
	return -1
}

func LastIndexInterface(slice []interface{}, target interface{}) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

func LastIndexString(slice []string, target string) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

func LastIndexInt(slice []int, target int) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexInt8(slice []int8, target int8) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexInt16(slice []int16, target int16) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexInt32(slice []int32, target int32) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexInt64(slice []int64, target int64) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}

func LastIndexUint(slice []uint, target uint) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexUint8(slice []uint8, target uint8) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexUint16(slice []uint16, target uint16) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexUint32(slice []uint32, target uint32) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
func LastIndexUint64(slice []uint64, target uint64) int {
	for i := len(slice) - 1; i >= 0; i-- {
		if slice[i] == target {
			return i
		}
	}
	return -1
}
