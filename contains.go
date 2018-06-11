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
