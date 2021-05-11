package slice

func ContainsString(slice []string, target string) bool {
	for _, s := range slice {
		if s == target {
			return true
		}
	}
	return false
}

func ContainsAnyString(slice []string, targets []string) bool {
	targetMap := map[string]bool{}
	for _, target := range targets {
		targetMap[target] = true
	}
	for _, s := range slice {
		if targetMap[s] {
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

func ContainsInt8(slice []int8, target int8) bool {
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

func ContainsAllInt64(slice, target []int64) bool {
	for _, id := range target {
		if !ContainsInt64(slice, id) {
			return false
		}
	}
	return true
}

func ContainsAllString(slice, target []string) bool {
	for _, id := range target {
		if !ContainsString(slice, id) {
			return false
		}
	}
	return true
}
