package slice

import (
	"fmt"
	"reflect"
)

type T struct {
	Id   int
	Name string
}

func ExampleIndex() {
	fmt.Println(Index(0, func(i int) bool { return true }))
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	var length = len(slice)
	fmt.Println(Index(length, func(i int) bool { return slice[i] == T{2, "c"} }))
	fmt.Println(Index(length, func(i int) bool { return slice[i] == T{2, "b"} }))
	// Output:
	// -1
	// -1
	// 2
}

func ExampleIndexGeneric() {
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	fmt.Println(IndexGeneric(nil, 3))
	fmt.Println(IndexGeneric([]int(nil), 3))
	fmt.Println(IndexGeneric(slice, 3))
	fmt.Println(IndexGeneric(slice, T{2, "c"}))
	fmt.Println(IndexGeneric(slice, T{2, "b"}))
	// Output:
	// -1
	// -1
	// -1
	// -1
	// 2
}

func ExampleIndexValue() {
	fmt.Println(IndexValue(reflect.ValueOf([]int(nil)), 3))
	var slice = reflect.ValueOf([]T{{3, "c"}, {}, {2, "b"}, {9, "f"}})
	fmt.Println(IndexValue(slice, 3))
	fmt.Println(IndexValue(slice, T{2, "c"}))
	fmt.Println(IndexValue(slice, T{2, "b"}))

	// Output:
	// -1
	// -1
	// -1
	// 2
}

func ExampleIndexInterface() {
	slice := []interface{}{true, `1`, 2, `3`}
	fmt.Println(IndexInterface(nil, `3`))
	fmt.Println(IndexInterface(slice, 0))
	fmt.Println(IndexInterface(slice, `2`))
	fmt.Println(IndexInterface(slice, true))
	fmt.Println(IndexInterface(slice, 2))
	// Output:
	// -1
	// -1
	// -1
	// 0
	// 2
}

func ExampleIndexString() {
	slice := []string{`1`, `2`, `3`}
	fmt.Println(IndexString(nil, `3`))
	fmt.Println(IndexString(slice, `0`))
	fmt.Println(IndexString(slice, `1`))
	fmt.Println(IndexString(slice, `3`))
	// Output:
	// -1
	// -1
	// 0
	// 2
}

func ExampleIndexInt() {
	slice := []int{1, 2, 3}
	fmt.Println(IndexInt(nil, 3))
	fmt.Println(IndexInt(slice, 0))
	fmt.Println(IndexInt(slice, 1))
	fmt.Println(IndexInt(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexInt8() {
	slice := []int8{1, 2, 3}
	fmt.Println(IndexInt8(nil, 3))
	fmt.Println(IndexInt8(slice, 0))
	fmt.Println(IndexInt8(slice, 1))
	fmt.Println(IndexInt8(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexInt16() {
	slice := []int16{1, 2, 3}
	fmt.Println(IndexInt16(nil, 3))
	fmt.Println(IndexInt16(slice, 0))
	fmt.Println(IndexInt16(slice, 1))
	fmt.Println(IndexInt16(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexInt32() {
	slice := []int32{1, 2, 3}
	fmt.Println(IndexInt32(nil, 3))
	fmt.Println(IndexInt32(slice, 0))
	fmt.Println(IndexInt32(slice, 1))
	fmt.Println(IndexInt32(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexInt64() {
	slice := []int64{1, 2, 3}
	fmt.Println(IndexInt64(nil, 3))
	fmt.Println(IndexInt64(slice, 0))
	fmt.Println(IndexInt64(slice, 1))
	fmt.Println(IndexInt64(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexUint() {
	slice := []uint{1, 2, 3}
	fmt.Println(IndexUint(nil, 3))
	fmt.Println(IndexUint(slice, 0))
	fmt.Println(IndexUint(slice, 1))
	fmt.Println(IndexUint(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexUint8() {
	slice := []uint8{1, 2, 3}
	fmt.Println(IndexUint8(nil, 3))
	fmt.Println(IndexUint8(slice, 0))
	fmt.Println(IndexUint8(slice, 1))
	fmt.Println(IndexUint8(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexUint16() {
	slice := []uint16{1, 2, 3}
	fmt.Println(IndexUint16(nil, 3))
	fmt.Println(IndexUint16(slice, 0))
	fmt.Println(IndexUint16(slice, 1))
	fmt.Println(IndexUint16(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexUint32() {
	slice := []uint32{1, 2, 3}
	fmt.Println(IndexUint32(nil, 3))
	fmt.Println(IndexUint32(slice, 0))
	fmt.Println(IndexUint32(slice, 1))
	fmt.Println(IndexUint32(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
func ExampleIndexUint64() {
	slice := []uint64{1, 2, 3}
	fmt.Println(IndexUint64(nil, 3))
	fmt.Println(IndexUint64(slice, 0))
	fmt.Println(IndexUint64(slice, 1))
	fmt.Println(IndexUint64(slice, 3))
	// Output:
	// -1
	// -1
	// 0
	// 2
}
