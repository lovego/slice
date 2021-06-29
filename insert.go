package slice

import "reflect"

// Insert insert value at index i, it panics if i > len(slice).
func InsertGeneric(slicePointer interface{}, i int, value interface{}) {
	InsertValue(reflect.ValueOf(slicePointer).Elem(), i, reflect.ValueOf(value))
}

// InsertValue insert value at index i, it panics if i > len(slice). slice must be settable.
func InsertValue(slice reflect.Value, i int, value reflect.Value) {
	slice.Set(reflect.Append(slice, value))
	if i != slice.Len()-1 {
		reflect.Copy(slice.Slice(i+1, slice.Len()), slice.Slice(i, slice.Len()-1))
		slice.Index(i).Set(value)
	}
}

// InsertInterface insert value at index i, it panics if i > len(slice).
func InsertInterface(slicePointer *[]interface{}, i int, value interface{}) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertString insert value at index i, it panics if i > len(slice).
func InsertString(slicePointer *[]string, i int, value string) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertInt insert value at index i, it panics if i > len(slice).
func InsertInt(slicePointer *[]int, i int, value int) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertInt8 insert value at index i, it panics if i > len(slice).
func InsertInt8(slicePointer *[]int8, i int, value int8) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertInt16 insert value at index i, it panics if i > len(slice).
func InsertInt16(slicePointer *[]int16, i int, value int16) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertInt32 insert value at index i, it panics if i > len(slice).
func InsertInt32(slicePointer *[]int32, i int, value int32) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertInt64 insert value at index i, it panics if i > len(slice).
func InsertInt64(slicePointer *[]int64, i int, value int64) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertUint insert value at index i, it panics if i > len(slice).
func InsertUint(slicePointer *[]uint, i int, value uint) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertUint8 insert value at index i, it panics if i > len(slice).
func InsertUint8(slicePointer *[]uint8, i int, value uint8) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertUint16 insert value at index i, it panics if i > len(slice).
func InsertUint16(slicePointer *[]uint16, i int, value uint16) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertUint32 insert value at index i, it panics if i > len(slice).
func InsertUint32(slicePointer *[]uint32, i int, value uint32) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}

// InsertUint64 insert value at index i, it panics if i > len(slice).
func InsertUint64(slicePointer *[]uint64, i int, value uint64) {
	slice := append(*slicePointer, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	*slicePointer = slice
}
