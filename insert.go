package slice

import "reflect"

// Insert insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertGeneric(slice interface{}, i int, value interface{}) interface{} {
	return InsertValue(reflect.ValueOf(slice), i, reflect.ValueOf(value))
}

// InsertValue insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertValue(slice reflect.Value, i int, value reflect.Value) reflect.Value {
	slice = reflect.Append(slice, value)
	if i != slice.Len()-1 {
		reflect.Copy(slice.Slice(i+1, slice.Len()), slice.Slice(i, slice.Len()-1))
		slice.Index(i).Set(value)
	}
	return slice
}

// InsertInterface insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertInterface(slice []interface{}, i int, value interface{}) []interface{} {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertString insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertString(slice []string, i int, value string) []string {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertInt insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertInt(slice []int, i int, value int) []int {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertInt8 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertInt8(slice []int8, i int, value int8) []int8 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertInt16 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertInt16(slice []int16, i int, value int16) []int16 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertInt32 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertInt32(slice []int32, i int, value int32) []int32 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertInt64 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertInt64(slice []int64, i int, value int64) []int64 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertUint insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertUint(slice []uint, i int, value uint) []uint {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertUint8 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertUint8(slice []uint8, i int, value uint8) []uint8 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertUint16 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertUint16(slice []uint16, i int, value uint16) []uint16 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertUint32 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertUint32(slice []uint32, i int, value uint32) []uint32 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}

// InsertUint64 insert value at index i, it panics if i > len(slice).
// The input slice will be modified.
func InsertUint64(slice []uint64, i int, value uint64) []uint64 {
	slice = append(slice, value)
	if i != len(slice)-1 {
		copy(slice[i+1:], slice[i:len(slice)-1])
		slice[i] = value
	}
	return slice
}
