package slice

import "reflect"

// InsertInt64 panics if i > len(s)
func InsertInt64(s []int64, i int, v int64) []int64 {
	s = append(s, v)
	if i != len(s)-1 {
		copy(s[i+1:], s[i:len(s)-1])
		s[i] = v
	}
	return s
}

// InsertString panics if i > len(s)
func InsertString(s []string, i int, v string) []string {
	s = append(s, v)
	if i != len(s)-1 {
		copy(s[i+1:], s[i:len(s)-1])
		s[i] = v
	}
	return s
}

// InsertValue panics if i > len(s)
func InsertValue(s reflect.Value, i int, v reflect.Value) reflect.Value {
	s = reflect.Append(s, v)
	if i != s.Len()-1 {
		reflect.Copy(s.Slice(i+1, s.Len()), s.Slice(i, s.Len()-1))
		s.Index(i).Set(v)
	}
	return s
}
