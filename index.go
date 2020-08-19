package slice

import "reflect"

func IndexString(slice []string, target string) int {
	for i := 0; i < len(slice); i++ {
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

func IndexInt(slice []int, target int) int {
	for i := 0; i < len(slice); i++ {
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

func IndexInt64(slice []int64, target int64) int {
	for i := 0; i < len(slice); i++ {
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

func Index(slice interface{}, target interface{}) int {
	if slice == nil {
		return -1
	}
	v := reflect.ValueOf(slice)
	for i := 0; i < v.Len(); i++ {
		if v.Index(i).Interface() == target {
			return i
		}
	}
	return -1
}

func LastIndex(slice interface{}, target interface{}) int {
	if slice == nil {
		return -1
	}
	v := reflect.ValueOf(slice)
	for i := v.Len() - 1; i >= 0; i-- {
		if v.Index(i).Interface() == target {
			return i
		}
	}
	return -1
}

func IndexValue(slice, target reflect.Value) int {
	if slice.Len() == 0 {
		return -1
	}
	t := target.Interface()
	for i := 0; i < slice.Len(); i++ {
		if slice.Index(i).Interface() == t {
			return i
		}
	}
	return -1
}

func LastIndexValue(slice, target reflect.Value) int {
	if slice.Len() == 0 {
		return -1
	}
	t := target.Interface()
	for i := slice.Len() - 1; i >= 0; i-- {
		if slice.Index(i).Interface() == t {
			return i
		}
	}
	return -1
}
