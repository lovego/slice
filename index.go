package slice

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
