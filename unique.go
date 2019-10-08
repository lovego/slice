package slice

import "reflect"

// Unique will change the original slice! use copy to keep the original slice.
// if trailing is true, will use the last element.
func Unique(s interface{}, primaryKey func(i int) interface{}, trailing bool) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr {
		panic(`Unique need pointer`)
	}
	val = val.Elem()
	if val.Kind() != reflect.Slice {
		panic(`Unique need pointer to Slice`)
	}
	length := val.Len()
	if length == 0 {
		return
	}
	seen := make(map[interface{}]struct{}, length)
	j := 0
	for i := 0; i < length; i++ {
		cursor := i
		cursor2 := j
		if trailing {
			cursor = length - i - 1
			cursor2 = length - j - 1
		}

		key := primaryKey(cursor)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		val.Index(cursor2).Set(val.Index(cursor))
		j++
	}
	start := 0
	end := j
	if trailing {
		start = length - j
		end = length
	}
	val.Set(val.Slice(start, end))
}

// UniqueInt will change the original slice!
func UniqueInt(s []int) []int {
	seen := make(map[int]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

// UniqueInt64 will change the original slice!
func UniqueInt64(s []int64) []int64 {
	seen := make(map[int64]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}

// UniqueString will change the original slice!
func UniqueString(s []string) []string {
	seen := make(map[string]struct{}, len(s))
	j := 0
	for _, v := range s {
		if _, ok := seen[v]; ok {
			continue
		}
		seen[v] = struct{}{}
		s[j] = v
		j++
	}
	return s[:j]
}
