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
