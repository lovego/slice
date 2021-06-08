package slice

import (
	"fmt"
	"reflect"
)

func ExampleLastIndex() {
	fmt.Println(LastIndex(0, func(i int) bool { return true }))
	var slice = []T{{3, "c"}, {}, {2, "b"}, {}, {9, "f"}}
	var length = len(slice)
	fmt.Println(LastIndex(length, func(i int) bool { return slice[i] == T{} }))
	fmt.Println(LastIndex(length, func(i int) bool { return slice[i] == T{2, "c"} }))
	// Output:
	// -1
	// 3
	// -1
}

func ExampleLastIndexGeneric() {
	fmt.Println(LastIndexGeneric(nil, T{}))
	var slice = []T{{3, "c"}, {}, {2, "b"}, {}, {9, "f"}}
	fmt.Println(LastIndexGeneric(slice, T{}))
	fmt.Println(LastIndexGeneric(slice, T{2, "c"}))
	// Output:
	// -1
	// 3
	// -1
}

func ExampleLastIndexValue() {
	fmt.Println(LastIndexValue(reflect.ValueOf([]T{}), T{}))
	var slice = reflect.ValueOf([]T{{3, "c"}, {}, {2, "b"}, {}, {9, "f"}})
	fmt.Println(LastIndexValue(slice, T{}))
	fmt.Println(LastIndexValue(slice, T{2, "c"}))
	// Output:
	// -1
	// 3
	// -1
}

func ExampleLastIndexInterface() {
	slice := []interface{}{`1`, `2`, `3`, `1`}
	fmt.Println(LastIndexInterface(slice, `1`))
	fmt.Println(LastIndexInterface(slice, 1))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexString() {
	slice := []string{`1`, `2`, `3`, `1`}
	fmt.Println(LastIndexString(slice, `1`))
	fmt.Println(LastIndexString(slice, `4`))
	// Output:
	// 3
	// -1
}

func ExampleLastIndexInt() {
	slice := []int{1, 2, 3, 1, 2}
	fmt.Println(LastIndexInt(slice, 1))
	fmt.Println(LastIndexInt(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexInt8() {
	slice := []int8{1, 2, 3, 1, 2}
	fmt.Println(LastIndexInt8(slice, 1))
	fmt.Println(LastIndexInt8(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexInt16() {
	slice := []int16{1, 2, 3, 1, 2}
	fmt.Println(LastIndexInt16(slice, 1))
	fmt.Println(LastIndexInt16(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexInt32() {
	slice := []int32{1, 2, 3, 1, 2}
	fmt.Println(LastIndexInt32(slice, 1))
	fmt.Println(LastIndexInt32(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexInt64() {
	var slice = []int64{1, 2, 3, 1, 2}
	fmt.Println(LastIndexInt64(slice, 1))
	fmt.Println(LastIndexInt64(slice, 4))
	// Output:
	// 3
	// -1
}

func ExampleLastIndexUint() {
	slice := []uint{1, 2, 3, 1, 2}
	fmt.Println(LastIndexUint(slice, 1))
	fmt.Println(LastIndexUint(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexUint8() {
	slice := []uint8{1, 2, 3, 1, 2}
	fmt.Println(LastIndexUint8(slice, 1))
	fmt.Println(LastIndexUint8(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexUint16() {
	slice := []uint16{1, 2, 3, 1, 2}
	fmt.Println(LastIndexUint16(slice, 1))
	fmt.Println(LastIndexUint16(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexUint32() {
	slice := []uint32{1, 2, 3, 1, 2}
	fmt.Println(LastIndexUint32(slice, 1))
	fmt.Println(LastIndexUint32(slice, 4))
	// Output:
	// 3
	// -1
}
func ExampleLastIndexUint64() {
	var slice = []uint64{1, 2, 3, 1, 2}
	fmt.Println(LastIndexUint64(slice, 1))
	fmt.Println(LastIndexUint64(slice, 4))
	// Output:
	// 3
	// -1
}
