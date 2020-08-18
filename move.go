package slice

import "reflect"

// MoveBackward move elements starting at i backward n steps. It panics if i > len(s)
func MoveBackward(s interface{}, i, n int) interface{} {
	v := reflect.ValueOf(s)
	l := v.Len() + n
	c := v.Cap()
	if l <= c {
		v = v.Slice(0, l)
	} else {
		v = v.Slice(0, c)
		v = reflect.AppendSlice(v, reflect.MakeSlice(v.Type(), l-c, l-c))
	}

	if i != l-n { // let it panic if i > l - n
		reflect.Copy(v.Slice(i+n, l), v.Slice(i, l-n))
	}
	return v.Interface()
}
