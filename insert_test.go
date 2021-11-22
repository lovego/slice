package slice

import (
	"fmt"
	"reflect"
)

func ExampleInsertGeneric() {
	fmt.Println(InsertGeneric([]string(nil), 0, "a"))
	fmt.Println(InsertGeneric([]string{}, 0, "a"))
	fmt.Println(InsertGeneric([]string{"b", "c"}, 0, "a"))
	fmt.Println(InsertGeneric([]string{"b", "c"}, 1, "a"))
	fmt.Println(InsertGeneric([]string{"b", "c"}, 2, "a"))
	// Output:
	// [a]
	// [a]
	// [a b c]
	// [b a c]
	// [b c a]
}

func ExampleInsertValue() {
	fmt.Println(InsertValue(reflect.ValueOf([]string(nil)), 0, reflect.ValueOf("a")))
	fmt.Println(InsertValue(reflect.ValueOf([]string{}), 0, reflect.ValueOf("a")))
	fmt.Println(InsertValue(reflect.ValueOf([]string{"b", "c"}), 0, reflect.ValueOf("a")))
	fmt.Println(InsertValue(reflect.ValueOf([]string{"b", "c"}), 1, reflect.ValueOf("a")))
	fmt.Println(InsertValue(reflect.ValueOf([]string{"b", "c"}), 2, reflect.ValueOf("a")))
	// Output:
	// [a]
	// [a]
	// [a b c]
	// [b a c]
	// [b c a]
}

func ExampleInsertInterface() {
	fmt.Println(InsertInterface(nil, 0, "a"))
	fmt.Println(InsertInterface([]interface{}{}, 0, "a"))
	fmt.Println(InsertInterface([]interface{}{"b", 1}, 0, "a"))
	fmt.Println(InsertInterface([]interface{}{"b", 1}, 1, "a"))
	fmt.Println(InsertInterface([]interface{}{"b", 1}, 2, "a"))
	// Output:
	// [a]
	// [a]
	// [a b 1]
	// [b a 1]
	// [b 1 a]
}

func ExampleInsertString() {
	fmt.Println(InsertString(nil, 0, "a"))
	fmt.Println(InsertString([]string{}, 0, "a"))
	fmt.Println(InsertString([]string{"b", "c"}, 0, "a"))
	fmt.Println(InsertString([]string{"b", "c"}, 1, "a"))
	fmt.Println(InsertString([]string{"b", "c"}, 2, "a"))
	// Output:
	// [a]
	// [a]
	// [a b c]
	// [b a c]
	// [b c a]
}

func ExampleInsertBool() {
	fmt.Println(InsertBool(nil, 0, true))
	fmt.Println(InsertBool([]bool{}, 0, true))
	fmt.Println(InsertBool([]bool{false, true}, 0, true))
	fmt.Println(InsertBool([]bool{false, true}, 1, true))
	fmt.Println(InsertBool([]bool{false, true}, 2, true))
	// Output:
	// [true]
	// [true]
	// [true false true]
	// [false true true]
	// [false true true]
}

func ExampleInsertInt() {
	fmt.Println(InsertInt(nil, 0, 3))
	fmt.Println(InsertInt([]int{}, 0, 3))
	fmt.Println(InsertInt([]int{2, 3}, 0, 1))
	fmt.Println(InsertInt([]int{2, 3}, 1, 1))
	fmt.Println(InsertInt([]int{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertInt8() {
	fmt.Println(InsertInt8(nil, 0, 3))
	fmt.Println(InsertInt8([]int8{}, 0, 3))
	fmt.Println(InsertInt8([]int8{2, 3}, 0, 1))
	fmt.Println(InsertInt8([]int8{2, 3}, 1, 1))
	fmt.Println(InsertInt8([]int8{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertInt16() {
	fmt.Println(InsertInt16(nil, 0, 3))
	fmt.Println(InsertInt16([]int16{}, 0, 3))
	fmt.Println(InsertInt16([]int16{2, 3}, 0, 1))
	fmt.Println(InsertInt16([]int16{2, 3}, 1, 1))
	fmt.Println(InsertInt16([]int16{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertInt32() {
	fmt.Println(InsertInt32(nil, 0, 3))
	fmt.Println(InsertInt32([]int32{}, 0, 3))
	fmt.Println(InsertInt32([]int32{2, 3}, 0, 1))
	fmt.Println(InsertInt32([]int32{2, 3}, 1, 1))
	fmt.Println(InsertInt32([]int32{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertInt64() {
	fmt.Println(InsertInt64(nil, 0, 3))
	fmt.Println(InsertInt64([]int64{}, 0, 3))
	fmt.Println(InsertInt64([]int64{2, 3}, 0, 1))
	fmt.Println(InsertInt64([]int64{2, 3}, 1, 1))
	fmt.Println(InsertInt64([]int64{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}

func ExampleInsertUint() {
	fmt.Println(InsertUint(nil, 0, 3))
	fmt.Println(InsertUint([]uint{}, 0, 3))
	fmt.Println(InsertUint([]uint{2, 3}, 0, 1))
	fmt.Println(InsertUint([]uint{2, 3}, 1, 1))
	fmt.Println(InsertUint([]uint{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint8() {
	fmt.Println(InsertUint8(nil, 0, 3))
	fmt.Println(InsertUint8([]uint8{}, 0, 3))
	fmt.Println(InsertUint8([]uint8{2, 3}, 0, 1))
	fmt.Println(InsertUint8([]uint8{2, 3}, 1, 1))
	fmt.Println(InsertUint8([]uint8{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint16() {
	fmt.Println(InsertUint16(nil, 0, 3))
	fmt.Println(InsertUint16([]uint16{}, 0, 3))
	fmt.Println(InsertUint16([]uint16{2, 3}, 0, 1))
	fmt.Println(InsertUint16([]uint16{2, 3}, 1, 1))
	fmt.Println(InsertUint16([]uint16{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint32() {
	fmt.Println(InsertUint32(nil, 0, 3))
	fmt.Println(InsertUint32([]uint32{}, 0, 3))
	fmt.Println(InsertUint32([]uint32{2, 3}, 0, 1))
	fmt.Println(InsertUint32([]uint32{2, 3}, 1, 1))
	fmt.Println(InsertUint32([]uint32{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint64() {
	fmt.Println(InsertUint64(nil, 0, 3))
	fmt.Println(InsertUint64([]uint64{}, 0, 3))
	fmt.Println(InsertUint64([]uint64{2, 3}, 0, 1))
	fmt.Println(InsertUint64([]uint64{2, 3}, 1, 1))
	fmt.Println(InsertUint64([]uint64{2, 3}, 2, 1))
	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
