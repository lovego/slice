package slice

func DiffInt64(left, right []int64) []int64 {
	if len(left) == 0 {
		return []int64{}
	}
	var result []int64
	for _, id := range left {
		if !ContainsInt64(right, id) {
			result = append(result, id)
		}
	}
	return result
}

func DiffString(left, right []string) []string {
	if len(left) == 0 {
		return []string{}
	}
	var result []string
	for _, id := range left {
		if !ContainsString(right, id) {
			result = append(result, id)
		}
	}
	return result
}
