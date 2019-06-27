package slice

import "reflect"

// Unique will change original slice
func Unique(s interface{}, primaryKey func(i int) interface{}) {
	val := reflect.ValueOf(s)
	if val.Kind() != reflect.Ptr {
		panic(`Unique need pointer`)
	}
	val = val.Elem()
	if val.Kind() != reflect.Slice {
		panic(`Unique need pointer to Slice`)
	}
	seen := make(map[interface{}]struct{}, val.Len())
	j := 0
	for i := 0; i < val.Len(); i++ {
		key := primaryKey(i)
		if _, ok := seen[key]; ok {
			continue
		}
		seen[key] = struct{}{}
		val.Index(j).Set(val.Index(i))
		j++
	}
	val.SetLen(j)
}

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
