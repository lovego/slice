package slice

import "reflect"

// AppendIfNo append value in targets to slice if slice has no this value.
func AppendIfNo(slice interface{}, targets ...interface{}) interface{} {
	sliceValue := reflect.ValueOf(slice)
	for i := range targets {
		target := reflect.ValueOf(targets[i])
		if IndexValue(sliceValue, target) < 0 {
			sliceValue = reflect.Append(sliceValue, target)
		}
	}
	return sliceValue.Interface()
}

func AppendIfNoValue(slice reflect.Value, targets ...reflect.Value) reflect.Value {
	for _, target := range targets {
		if IndexValue(slice, target) < 0 {
			slice = reflect.Append(slice, target)
		}
	}
	return slice
}

func AppendIfNoInterface(slice []interface{}, targets ...interface{}) []interface{} {
	for _, target := range targets {
		if IndexInterface(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoString(slice []string, targets ...string) []string {
	for _, target := range targets {
		if IndexString(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoInt(slice []int, targets ...int) []int {
	for _, target := range targets {
		if IndexInt(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoInt8(slice []int8, targets ...int8) []int8 {
	for _, target := range targets {
		if IndexInt8(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoInt16(slice []int16, targets ...int16) []int16 {
	for _, target := range targets {
		if IndexInt16(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoInt32(slice []int32, targets ...int32) []int32 {
	for _, target := range targets {
		if IndexInt32(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoInt64(slice []int64, targets ...int64) []int64 {
	for _, target := range targets {
		if IndexInt64(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoUint(slice []uint, targets ...uint) []uint {
	for _, target := range targets {
		if IndexUint(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoUint8(slice []uint8, targets ...uint8) []uint8 {
	for _, target := range targets {
		if IndexUint8(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoUint16(slice []uint16, targets ...uint16) []uint16 {
	for _, target := range targets {
		if IndexUint16(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoUint32(slice []uint32, targets ...uint32) []uint32 {
	for _, target := range targets {
		if IndexUint32(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNoUint64(slice []uint64, targets ...uint64) []uint64 {
	for _, target := range targets {
		if IndexUint64(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}
