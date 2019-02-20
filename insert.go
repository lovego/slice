package slice

import "reflect"

func InsertInt64(s []int64, i int, v int64) []int64 {
	s = append(s, v)
	if i != len(s) { // let it panic if i > len(s)
		copy(s[i+1:], s[i:len(s)-1])
		s[i] = v
	}
	return s
}

func InsertString(s []string, i int, v string) []string {
	s = append(s, v)
	if i != len(s) { // let it panic if i > len(s)
		copy(s[i+1:], s[i:len(s)-1])
		s[i] = v
	}
	return s
}

func InsertValue(s reflect.Value, i int, v reflect.Value) reflect.Value {
	s = reflect.Append(s, v)
	if i != s.Len() { // let it panic if i > len(s)
		reflect.Copy(s.Slice(i+1, s.Len()), s.Slice(i, s.Len()-1))
		s.Index(i).Set(v)
	}
	return s
}
