package slice

import "reflect"

func Unique(slicePointer interface{}) {
	UniqueValue(reflect.ValueOf(slicePointer))
}

func UniqueValue(slice reflect.Value) {
	if slice.Kind() == reflect.Ptr {
		slice = slice.Elem()
	}
	j := 0
	length := slice.Len()
	for i := 0; i < length; i++ {
		v := slice.Index(i)
		if !ContainsValue(slice.Slice(0, j), v.Interface()) {
			slice.Index(j).Set(v)
			j++
		}
	}
	slice.Set(slice.Slice(0, j))
}

func UniqueInterface(slice *[]interface{}) {
	j := 0
	for _, v := range *slice {
		if !ContainsInterface((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}

func UniqueString(slice *[]string) {
	j := 0
	for _, v := range *slice {
		if !ContainsString((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}

func UniqueInt(slice *[]int) {
	j := 0
	for _, v := range *slice {
		if !ContainsInt((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueInt8(slice *[]int8) {
	j := 0
	for _, v := range *slice {
		if !ContainsInt8((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueInt16(slice *[]int16) {
	j := 0
	for _, v := range *slice {
		if !ContainsInt16((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueInt32(slice *[]int32) {
	j := 0
	for _, v := range *slice {
		if !ContainsInt32((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueInt64(slice *[]int64) {
	j := 0
	for _, v := range *slice {
		if !ContainsInt64((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}

func UniqueUint(slice *[]uint) {
	j := 0
	for _, v := range *slice {
		if !ContainsUint((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueUint8(slice *[]uint8) {
	j := 0
	for _, v := range *slice {
		if !ContainsUint8((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueUint16(slice *[]uint16) {
	j := 0
	for _, v := range *slice {
		if !ContainsUint16((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueUint32(slice *[]uint32) {
	j := 0
	for _, v := range *slice {
		if !ContainsUint32((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
func UniqueUint64(slice *[]uint64) {
	j := 0
	for _, v := range *slice {
		if !ContainsUint64((*slice)[0:j], v) {
			(*slice)[j] = v
			j++
		}
	}
	*slice = (*slice)[:j]
}
