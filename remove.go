package slice

func Remove(slice, toRemove []interface{}) (result []interface{}) {
	if len(slice) == 0 || len(toRemove) == 0 {
		return slice
	}
	m := make(map[interface{}]bool)
	for _, v := range toRemove {
		m[v] = true
	}
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}

func RemoveStrings(slice, toRemove []string) (result []string) {
	if len(slice) == 0 || len(toRemove) == 0 {
		return slice
	}
	m := make(map[string]bool)
	for _, v := range toRemove {
		m[v] = true
	}
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}

func RemoveInts(slice, toRemove []int) (result []int) {
	if len(slice) == 0 || len(toRemove) == 0 {
		return slice
	}
	m := make(map[int]bool)
	for _, v := range toRemove {
		m[v] = true
	}
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}

func RemoveInt64s(slice, toRemove []int64) (result []int64) {
	if len(slice) == 0 || len(toRemove) == 0 {
		return slice
	}
	m := make(map[int64]bool)
	for _, v := range toRemove {
		m[v] = true
	}
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}
