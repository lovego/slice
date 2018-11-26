package slice

// string数组取并集
func UnionString(left, right []string) []string {
	if len(left) == 0 {
		return right
	}
	if len(right) == 0 {
		return left
	}
	m := make(map[string]bool)
	result := []string{}
	for _, v := range left {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	for _, v := range right {
		if !m[v] {
			m[v] = true
			result = append(result, v)
		}
	}
	return result
}
