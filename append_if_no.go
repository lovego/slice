package slice

import "reflect"

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

func AppendIfNoInt64(slice []int64, targets ...int64) []int64 {
	for _, target := range targets {
		if IndexInt64(slice, target) < 0 {
			slice = append(slice, target)
		}
	}
	return slice
}

func AppendIfNo(slice interface{}, targets ...interface{}) interface{} {
	sliceV := reflect.ValueOf(slice)
	for i := range targets {
		targetV := reflect.ValueOf(targets[i])
		if IndexValue(sliceV, targetV) < 0 {
			sliceV = reflect.Append(sliceV, targetV)
		}
	}
	return sliceV.Interface()
}
