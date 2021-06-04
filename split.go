package slice

func SplitString(s []string, sep string) (result [][]string) {
	start := 0
	for i, v := range s {
		if v == sep {
			result = append(result, s[start:i])
			start = i + 1
		}
	}

	if len(s) > 0 {
		result = append(result, s[start:len(s)])
	}
	return
}
