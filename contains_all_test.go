package slice

import (
	"fmt"
	"reflect"
)

func ExampleContainsAllGeneric() {
	fmt.Println(ContainsAllGeneric(nil))
	fmt.Println(ContainsAllGeneric(nil, T{}))
	fmt.Println(ContainsAllGeneric([]bool{}))
	fmt.Println(ContainsAllGeneric([]bool{}, "xx"))
	// Output:
	// true
	// false
	// true
	// false
}

func ExampleContainsAllGeneric_2() {
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	fmt.Println(ContainsAllGeneric(slice))
	fmt.Println(ContainsAllGeneric(slice, T{}))
	fmt.Println(ContainsAllGeneric(slice, T{}, T{2, "b"}))
	fmt.Println(ContainsAllGeneric(slice, T{2, "c"}))
	fmt.Println(ContainsAllGeneric(slice, 2))
	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleContainsAllValue() {
	var slice = reflect.ValueOf([]interface{}{`1`, 2, `3`})
	fmt.Println(ContainsAllValue(slice, 2, `3`))
	fmt.Println(ContainsAllValue(slice, `2`, `3`))
	fmt.Println(ContainsAllValue(reflect.ValueOf([]int{}), 2))
	// Output:
	// true
	// false
	// false
}

func ExampleContainsAllInterface() {
	var slice = []interface{}{`1`, 2, `3`}
	fmt.Println(ContainsAllInterface(slice, 2, `3`))
	fmt.Println(ContainsAllInterface(slice, `2`, `3`))
	// Output:
	// true
	// false
}
func ExampleContainsAllString() {
	fmt.Println(ContainsAllString([]string{"1", "2"}))
	fmt.Println(ContainsAllString([]string{"1", "2"}, "1"))
	fmt.Println(ContainsAllString([]string{"1", "2"}, "2", "1"))
	fmt.Println(ContainsAllString([]string{"1", "2"}, "2", "3"))
	fmt.Println(ContainsAllString([]string{"1", "2"}, "3", "4"))
	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleContainsAllInt() {
	fmt.Println(ContainsAllInt([]int{1, 2}))
	fmt.Println(ContainsAllInt([]int{1, 2}, 1))
	fmt.Println(ContainsAllInt([]int{1, 2}, 2, 1))
	fmt.Println(ContainsAllInt([]int{1, 2}, 2, 3))
	fmt.Println(ContainsAllInt([]int{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllInt8() {
	fmt.Println(ContainsAllInt8([]int8{1, 2}))
	fmt.Println(ContainsAllInt8([]int8{1, 2}, 1))
	fmt.Println(ContainsAllInt8([]int8{1, 2}, 2, 1))
	fmt.Println(ContainsAllInt8([]int8{1, 2}, 2, 3))
	fmt.Println(ContainsAllInt8([]int8{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllInt16() {
	fmt.Println(ContainsAllInt16([]int16{1, 2}))
	fmt.Println(ContainsAllInt16([]int16{1, 2}, 1))
	fmt.Println(ContainsAllInt16([]int16{1, 2}, 2, 1))
	fmt.Println(ContainsAllInt16([]int16{1, 2}, 2, 3))
	fmt.Println(ContainsAllInt16([]int16{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllInt32() {
	fmt.Println(ContainsAllInt32([]int32{1, 2}))
	fmt.Println(ContainsAllInt32([]int32{1, 2}, 1))
	fmt.Println(ContainsAllInt32([]int32{1, 2}, 2, 1))
	fmt.Println(ContainsAllInt32([]int32{1, 2}, 2, 3))
	fmt.Println(ContainsAllInt32([]int32{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllInt64() {
	fmt.Println(ContainsAllInt64([]int64{1, 2}))
	fmt.Println(ContainsAllInt64([]int64{1, 2}, 1))
	fmt.Println(ContainsAllInt64([]int64{1, 2}, 2, 1))
	fmt.Println(ContainsAllInt64([]int64{1, 2}, 2, 3))
	fmt.Println(ContainsAllInt64([]int64{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleContainsAllUint() {
	fmt.Println(ContainsAllUint([]uint{1, 2}))
	fmt.Println(ContainsAllUint([]uint{1, 2}, 1))
	fmt.Println(ContainsAllUint([]uint{1, 2}, 2, 1))
	fmt.Println(ContainsAllUint([]uint{1, 2}, 2, 3))
	fmt.Println(ContainsAllUint([]uint{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllUint8() {
	fmt.Println(ContainsAllUint8([]uint8{1, 2}))
	fmt.Println(ContainsAllUint8([]uint8{1, 2}, 1))
	fmt.Println(ContainsAllUint8([]uint8{1, 2}, 2, 1))
	fmt.Println(ContainsAllUint8([]uint8{1, 2}, 2, 3))
	fmt.Println(ContainsAllUint8([]uint8{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllUint16() {
	fmt.Println(ContainsAllUint16([]uint16{1, 2}))
	fmt.Println(ContainsAllUint16([]uint16{1, 2}, 1))
	fmt.Println(ContainsAllUint16([]uint16{1, 2}, 2, 1))
	fmt.Println(ContainsAllUint16([]uint16{1, 2}, 2, 3))
	fmt.Println(ContainsAllUint16([]uint16{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllUint32() {
	fmt.Println(ContainsAllUint32([]uint32{1, 2}))
	fmt.Println(ContainsAllUint32([]uint32{1, 2}, 1))
	fmt.Println(ContainsAllUint32([]uint32{1, 2}, 2, 1))
	fmt.Println(ContainsAllUint32([]uint32{1, 2}, 2, 3))
	fmt.Println(ContainsAllUint32([]uint32{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
func ExampleContainsAllUint64() {
	fmt.Println(ContainsAllUint64([]uint64{1, 2}))
	fmt.Println(ContainsAllUint64([]uint64{1, 2}, 1))
	fmt.Println(ContainsAllUint64([]uint64{1, 2}, 2, 1))
	fmt.Println(ContainsAllUint64([]uint64{1, 2}, 2, 3))
	fmt.Println(ContainsAllUint64([]uint64{1, 2}, 3, 4))
	// Output:
	// true
	// true
	// true
	// false
	// false
}
