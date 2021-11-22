package slice

import (
	"fmt"
	"reflect"
)

func ExampleContainsAnyGeneric() {
	fmt.Println(ContainsAnyGeneric(nil))
	fmt.Println(ContainsAnyGeneric(nil, T{}))
	fmt.Println(ContainsAnyGeneric([]bool{}))
	fmt.Println(ContainsAnyGeneric([]bool{}, "xx"))
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	fmt.Println(ContainsAnyGeneric(slice))
	fmt.Println(ContainsAnyGeneric(slice, T{}))
	fmt.Println(ContainsAnyGeneric(slice, T{}, T{2, "b"}))
	fmt.Println(ContainsAnyGeneric(slice, T{2, "c"}, T{}))
	fmt.Println(ContainsAnyGeneric(slice, 2))
	// Output:
	// false
	// false
	// false
	// false
	// false
	// true
	// true
	// true
	// false
}

func ExampleContainsAnyValue() {
	fmt.Println(ContainsAnyValue(reflect.ValueOf([]int{}), 4))
	var slice = reflect.ValueOf([]interface{}{`1`, 2, `3`})
	fmt.Println(ContainsAnyValue(slice, 2, `3`))
	fmt.Println(ContainsAnyValue(slice, `x`, `3`))
	fmt.Println(ContainsAnyValue(slice, `x`, 4))
	// Output:
	// false
	// true
	// true
	// false
}

func ExampleContainsAnyInterface() {
	fmt.Println(ContainsAnyInterface(nil))
	var slice = []interface{}{`1`, 2, `3`}
	fmt.Println(ContainsAnyInterface(slice, 2, `3`))
	fmt.Println(ContainsAnyInterface(slice, `2`, `3`))
	fmt.Println(ContainsAnyInterface(slice, `x`))
	// Output:
	// false
	// true
	// true
	// false
}

func ExampleContainsAnyString() {
	fmt.Println(ContainsAnyString(nil))
	fmt.Println(ContainsAnyString(nil, `x`))
	var slice = []string{`1`, `2`, `3`}
	fmt.Println(ContainsAnyString(slice))
	fmt.Println(ContainsAnyString(slice, `x`, `0`))
	fmt.Println(ContainsAnyString(slice, `x`, `2`))
	// Output:
	// false
	// false
	// false
	// false
	// true
}

func ExampleContainsAnyBool() {
	fmt.Println(ContainsAnyBool(nil))
	fmt.Println(ContainsAnyBool(nil, false))
	fmt.Println(ContainsAnyBool([]bool{true, false}))
	fmt.Println(ContainsAnyBool([]bool{true, true}, false))
	fmt.Println(ContainsAnyBool([]bool{true, false}, false))
	// Output:
	// false
	// false
	// false
	// false
	// true
}

func ExampleContainsAnyInt() {
	fmt.Println(ContainsAnyInt(nil))
	fmt.Println(ContainsAnyInt(nil, 0))
	var slice = []int{1, 2, 3}
	fmt.Println(ContainsAnyInt(slice))
	fmt.Println(ContainsAnyInt(slice, 3))
	fmt.Println(ContainsAnyInt(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyInt8() {
	fmt.Println(ContainsAnyInt8(nil))
	fmt.Println(ContainsAnyInt8(nil, 0))
	var slice = []int8{1, 2, 3}
	fmt.Println(ContainsAnyInt8(slice))
	fmt.Println(ContainsAnyInt8(slice, 3))
	fmt.Println(ContainsAnyInt8(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyInt16() {
	fmt.Println(ContainsAnyInt16(nil))
	fmt.Println(ContainsAnyInt16(nil, 0))
	var slice = []int16{1, 2, 3}
	fmt.Println(ContainsAnyInt16(slice))
	fmt.Println(ContainsAnyInt16(slice, 3))
	fmt.Println(ContainsAnyInt16(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyInt32() {
	fmt.Println(ContainsAnyInt32(nil))
	fmt.Println(ContainsAnyInt32(nil, 0))
	var slice = []int32{1, 2, 3}
	fmt.Println(ContainsAnyInt32(slice))
	fmt.Println(ContainsAnyInt32(slice, 3))
	fmt.Println(ContainsAnyInt32(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyInt64() {
	fmt.Println(ContainsAnyInt64(nil))
	fmt.Println(ContainsAnyInt64(nil, 0))
	var slice = []int64{1, 2, 3}
	fmt.Println(ContainsAnyInt64(slice))
	fmt.Println(ContainsAnyInt64(slice, 3))
	fmt.Println(ContainsAnyInt64(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}

func ExampleContainsAnyUint() {
	fmt.Println(ContainsAnyUint(nil))
	fmt.Println(ContainsAnyUint(nil, 0))
	var slice = []uint{1, 2, 3}
	fmt.Println(ContainsAnyUint(slice))
	fmt.Println(ContainsAnyUint(slice, 3))
	fmt.Println(ContainsAnyUint(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyUint8() {
	fmt.Println(ContainsAnyUint8(nil))
	fmt.Println(ContainsAnyUint8(nil, 0))
	var slice = []uint8{1, 2, 3}
	fmt.Println(ContainsAnyUint8(slice))
	fmt.Println(ContainsAnyUint8(slice, 3))
	fmt.Println(ContainsAnyUint8(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyUint16() {
	fmt.Println(ContainsAnyUint16(nil))
	fmt.Println(ContainsAnyUint16(nil, 0))
	var slice = []uint16{1, 2, 3}
	fmt.Println(ContainsAnyUint16(slice))
	fmt.Println(ContainsAnyUint16(slice, 3))
	fmt.Println(ContainsAnyUint16(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyUint32() {
	fmt.Println(ContainsAnyUint32(nil))
	fmt.Println(ContainsAnyUint32(nil, 0))
	var slice = []uint32{1, 2, 3}
	fmt.Println(ContainsAnyUint32(slice))
	fmt.Println(ContainsAnyUint32(slice, 3))
	fmt.Println(ContainsAnyUint32(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
func ExampleContainsAnyUint64() {
	fmt.Println(ContainsAnyUint64(nil))
	fmt.Println(ContainsAnyUint64(nil, 0))
	var slice = []uint64{1, 2, 3}
	fmt.Println(ContainsAnyUint64(slice))
	fmt.Println(ContainsAnyUint64(slice, 3))
	fmt.Println(ContainsAnyUint64(slice, 0, 4))
	// Output:
	// false
	// false
	// false
	// true
	// false
}
