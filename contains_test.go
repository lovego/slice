package slice

import (
	"fmt"
	"reflect"
)

func ExampleContains() {
	fmt.Println(Contains(0, func(i int) bool { return true }))
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	var length = len(slice)
	fmt.Println(Contains(length, func(i int) bool { return slice[i] == T{} }))
	fmt.Println(Contains(length, func(i int) bool { return slice[i] == T{2, "b"} }))
	fmt.Println(Contains(length, func(i int) bool { return slice[i] == T{2, "c"} }))
	// Output:
	// false
	// true
	// true
	// false
}

func ExampleContainsGeneric() {
	fmt.Println(ContainsGeneric(nil, T{}))
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	fmt.Println(ContainsGeneric(slice, T{}))
	fmt.Println(ContainsGeneric(slice, T{2, "b"}))
	fmt.Println(ContainsGeneric(slice, T{2, "c"}))
	fmt.Println(ContainsGeneric(slice, 2))
	// Output:
	// false
	// true
	// true
	// false
	// false
}

func ExampleContainsValue() {
	var slice = reflect.ValueOf([]interface{}{`1`, 2, `3`})
	fmt.Println(ContainsValue(slice, 2))
	fmt.Println(ContainsValue(slice, `2`))
	fmt.Println(ContainsValue(reflect.ValueOf([]int{}), 2))
	// Output:
	// true
	// false
	// false
}

func ExampleContainsInterface() {
	var slice = []interface{}{`1`, 2, `3`}
	fmt.Println(ContainsInterface(slice, 2))
	fmt.Println(ContainsInterface(slice, `2`))
	// Output:
	// true
	// false
}

func ExampleContainsString() {
	var slice = []string{`1`, `2`, `3`}
	fmt.Println(ContainsString(slice, `2`))
	fmt.Println(ContainsString(slice, `4`))
	// Output:
	// true
	// false
}

func ExampleContainsInt() {
	var slice = []int{1, 2, 3}
	fmt.Println(ContainsInt(slice, 2))
	fmt.Println(ContainsInt(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsInt8() {
	var slice = []int8{1, 2, 3}
	fmt.Println(ContainsInt8(slice, 2))
	fmt.Println(ContainsInt8(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsInt16() {
	var slice = []int16{1, 2, 3}
	fmt.Println(ContainsInt16(slice, 2))
	fmt.Println(ContainsInt16(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsInt32() {
	var slice = []int32{1, 2, 3}
	fmt.Println(ContainsInt32(slice, 2))
	fmt.Println(ContainsInt32(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsInt64() {
	var slice = []int64{1, 2, 3}
	fmt.Println(ContainsInt64(slice, 2))
	fmt.Println(ContainsInt64(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsUint() {
	var slice = []uint{1, 2, 3}
	fmt.Println(ContainsUint(slice, 2))
	fmt.Println(ContainsUint(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsUint8() {
	var slice = []uint8{1, 2, 3}
	fmt.Println(ContainsUint8(slice, 2))
	fmt.Println(ContainsUint8(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsUint16() {
	var slice = []uint16{1, 2, 3}
	fmt.Println(ContainsUint16(slice, 2))
	fmt.Println(ContainsUint16(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsUint32() {
	var slice = []uint32{1, 2, 3}
	fmt.Println(ContainsUint32(slice, 2))
	fmt.Println(ContainsUint32(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsUint64() {
	var slice = []uint64{1, 2, 3}
	fmt.Println(ContainsUint64(slice, 2))
	fmt.Println(ContainsUint64(slice, 4))
	// Output:
	// true
	// false
}
