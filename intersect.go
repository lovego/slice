package slice

// int64数组交集
func IntersectInt64(left, right []int64) []int64 {
	if len(left) == 0 || len(right) == 0 {
		return []int64{}
	}
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

// int64数组是否有交集
func HasIntersectInt64(left, right []int64) bool {
	if len(left) == 0 || len(right) == 0 {
		return false
	}
	for _, elem := range left {
		if ContainsInt64(right, elem) {
			return true
		}
	}
	return false
}

// string数组是否有交集
func HasIntersectString(left, right []string) bool {
	if len(left) == 0 || len(right) == 0 {
		return false
	}
	for _, elem := range left {
		if ContainsString(right, elem) {
			return true
		}
	}
	return false
}
