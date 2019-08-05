package slice

func DifferenceInt64(left, right []int64) []int64 {
	if len(left) == 0 || len(right) == 0 {
		return []int64{}
	}
	var result []int64
	for _, id := range left {
		if ContainsInt64(right, id) {
			continue
		}
		result = append(result, id)
	}
	return result
}

func DifferenceString(left, right []string) []string {
	if len(left) == 0 || len(right) == 0 {
		return []string{}
	}
	var result []string
	for _, id := range left {
		if ContainsString(right, id) {
			continue
		}
		result = append(result, id)
	}
	return result
}
