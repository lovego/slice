package slice

func AppendIfNoString(slice []string, target string) []string {
	if IndexString(slice, target) < 0 {
		slice = append(slice, target)
	}
	return slice
}

func AppendIfNoInt(slice []int, target int) []int {
	if IndexInt(slice, target) < 0 {
		slice = append(slice, target)
	}
	return slice
}

func AppendIfNoInt64(slice []int64, target int64) []int64 {
	if IndexInt64(slice, target) < 0 {
		slice = append(slice, target)
	}
	return slice
}
