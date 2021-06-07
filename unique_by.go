package slice

import "reflect"

// UniqueBy will change the original slice, Use copy to keep the original slice.
// Param slice should be a pointer to a slice.
// If param keepLast is true, the last one of the same elements is kept,  otherwise the first one is kept.
func UniqueBy(slice interface{}, keyFunc func(i int) interface{}, keepLast bool) {
	UniqueByValue(reflect.ValueOf(slice), keyFunc, keepLast)
}

// UniqueByValue will change the original slice, Use copy to keep the original slice.
// Param slice should be a pointer to a slice.
// If param keepLast is true, the last one of the same elements is kept,  otherwise the first one is kept.
func UniqueByValue(slice reflect.Value, keyFunc func(i int) interface{}, keepLast bool) {
	slice = slice.Elem()
	length := slice.Len()
	if keepLast {
		j := length
		for i := length - 1; i >= 0; i-- {
			key := keyFunc(i)
			if !containsInRange(j, length, keyFunc, key) {
				j--
				slice.Index(j).Set(slice.Index(i))
			}
		}
		slice.Set(slice.Slice(j, length))
	} else {
		j := 0
		for i := 0; i < length; i++ {
			key := keyFunc(i)
			if !containsInRange(0, j, keyFunc, key) {
				slice.Index(j).Set(slice.Index(i))
				j++
			}
		}
		slice.Set(slice.Slice(0, j))
	}
}

func containsInRange(start, end int, keyFunc func(i int) interface{}, target interface{}) bool {
	for i := start; i < end; i++ {
		if keyFunc(i) == target {
			return true
		}
	}
	return false
}
