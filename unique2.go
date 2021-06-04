package slice

import "reflect"

// UniqueField returns unique values of a field from a struct slice
func UniqueField(slice interface{}, fields ...string) (result []interface{}) {
	if len(fields) == 0 {
		panic(`fields is required.`)
	}

	value := reflect.ValueOf(slice)
	switch value.Kind() {
	case reflect.Slice, reflect.Array:
	default:
		panic(`first param must be a slice or array.`)
	}

	l := value.Len()
	if l == 0 {
		return
	}

	seen := make(map[interface{}]struct{}, l)
	for i := 0; i < l; i++ {
		v := value.Index(i)
		for _, field := range fields {
			v = v.FieldByName(field)
		}
		ifc := v.Interface()
		if _, ok := seen[ifc]; !ok {
			result = append(result, ifc)
			seen[ifc] = struct{}{}
		}
	}
	return
}

// UniqueBy will change the original slice! use copy to keep the original slice.
// if trailing is true, will use the last element.
func UniqueBy(s interface{}, uniqueKey func(i int) interface{}, trailing bool) {
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

		key := uniqueKey(cursor)
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
