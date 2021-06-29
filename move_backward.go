package slice

import "reflect"

// MoveBackward move elements starting at i backward n steps. It panics if i > len(s)
func MoveBackward(slicePointer interface{}, i, n int) {
	slice := reflect.ValueOf(slicePointer).Elem()
	l := slice.Len() + n
	c := slice.Cap()
	if l <= c {
		slice.Set(slice.Slice(0, l))
	} else {
		slice.Set(slice.Slice(0, c))
		slice.Set(reflect.AppendSlice(slice, reflect.MakeSlice(slice.Type(), l-c, l-c)))
	}

	if i != l-n { // let it panic if i > l - n
		reflect.Copy(slice.Slice(i+n, l), slice.Slice(i, l-n))
	}
}
