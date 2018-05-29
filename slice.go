package slice

func ContainsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func ContainsInt(slice []int, target int) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func ContainsInt64(slice []int64, target int64) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

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

// int64数组交集
func IntersectInt64(left, right []int64) []int64 {
	m := make(map[int64]bool)
	for _, elem := range left {
		m[elem] = true
	}
	var result []int64
	for _, elem := range right {
		if m[elem] {
			result = append(result, elem)
		}
	}
	return result
}
