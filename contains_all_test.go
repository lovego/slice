package slice

import (
	"fmt"
	"reflect"
)

func ExampleContainsAll() {
	fmt.Println(ContainsAll(nil))
	fmt.Println(ContainsAll(nil, T{}))
	fmt.Println(ContainsAll([]bool{}))
	fmt.Println(ContainsAll([]bool{}, "xx"))
	// Output:
	// true
	// false
	// true
	// false
}

func ExampleContainsAll_2() {
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	fmt.Println(ContainsAll(slice))
	fmt.Println(ContainsAll(slice, T{}))
	fmt.Println(ContainsAll(slice, T{}, T{2, "b"}))
	fmt.Println(ContainsAll(slice, T{2, "c"}))
	fmt.Println(ContainsAll(slice, 2))
	// Output:
	// true
	// true
	// true
	// false
	// false
}

func ExampleContainsAllValue() {
	var slice = reflect.ValueOf([]interface{}{`1`, 2, `3`})
	fmt.Println(ContainsAllValue(slice, reflect.ValueOf(2), reflect.ValueOf(`3`)))
	fmt.Println(ContainsAllValue(slice, reflect.ValueOf(`2`), reflect.ValueOf(`3`)))
	fmt.Println(ContainsAllValue(reflect.ValueOf([]int{}), reflect.ValueOf(2)))
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
