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

func FullContainsInt64(slice, target []int64) bool {
	if len(slice) == 0 && len(target) == 0 {
		return true
	}

	left, right := UniqueInt64(slice), UniqueInt64(target)
	if len(left) != len(right) {
		return false
	}

	for _, id := range left {
		if !ContainsInt64(right, id) {
			return false
		}
	}
	return true
}

func FullContainsString(slice, target []string) bool {
	if len(slice) == 0 && len(target) == 0 {
		return true
	}

	left, right := UniqueString(slice), UniqueString(target)
	if len(left) != len(right) {
		return false
	}

	for _, id := range left {
		if !ContainsString(right, id) {
			return false
		}
	}
	return true
}
