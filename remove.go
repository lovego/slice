package slice

func Remove(slice, toRemove []interface{}) (result []interface{}) {
	if len(slice) == 0 || len(toRemove) == 0 {
		return slice
	}
	m := make(map[interface{}]bool)
	for _, v := range toRemove {
		m[v] = true
	}
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}

func RemoveStrings(slice, toRemove []string) (result []string) {
	if len(slice) == 0 || len(toRemove) == 0 {
		return slice
	}
	m := make(map[string]bool)
	for _, v := range toRemove {
		m[v] = true
	}
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}

func RemoveInts(slice, toRemove []int) (result []int) {
	if len(slice) == 0 || len(toRemove) == 0 {
		return slice
	}
	m := make(map[int]bool)
	for _, v := range toRemove {
		m[v] = true
	}
	for _, v := range slice {
		if !m[v] {
			result = append(result, v)
		}
	}
	return
}

// func doRemove(slice, toRemove interface{}) (result interface{}) {
// 	fmt.Println(reflect.TypeOf(slice), reflect.TypeOf(toRemove))
// 	doSlice, ok := slice.([]string)
// 	if !ok {
// 		fmt.Println("...")
// 		return
// 	}
// 	doToRemove, ok := toRemove.([]interface{})
// 	if !ok {
// 		fmt.Println("...-===")
// 		return
// 	}
// 	if len(doSlice) == 0 || len(doToRemove) == 0 {
// 		return doSlice
// 	}
// 	m := make(map[interface{}]bool)
// 	for _, v := range doToRemove {
// 		m[v] = true
// 	}
// 	var resultSlice []interface{}
// 	for _, v := range doSlice {
// 		if !m[v] {
// 			resultSlice = append(resultSlice, v)
// 		}
// 	}
// 	result = resultSlice
// 	return
// }
