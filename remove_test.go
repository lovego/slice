package slice

import (
	"fmt"
	"reflect"
)

func ExampleRemove() {
	fmt.Println(Remove([]int{1, 2, 2, 3}, func(i int) bool { return i < 2 }))
	// Output:
	// [2 3]
}

func ExampleRemoveGeneric() {
	fmt.Println(RemoveGeneric([]int{1, 2, 2, 3}))
	fmt.Println(RemoveGeneric([]int{1, 2, 2, 3}, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
}

func ExampleRemoveValue() {
	fmt.Println(RemoveValue(reflect.ValueOf([]int{1, 2, 2, 3}), 1, 3))
	// Output:
	// [2 2]
}

func ExampleRemoveInterface() {
	fmt.Println(RemoveInterface([]interface{}{1, 2, 2, "c"}))
	fmt.Println(RemoveInterface([]interface{}{1, 2, 2, "c"}, 2))
	fmt.Println(RemoveInterface([]interface{}{1, 2, 2, "c"}, 1, 2))
	// Output:
	// [1 2 2 c]
	// [1 c]
	// [c]
}

func ExampleRemoveString() {
	fmt.Println(RemoveString([]string{"a", "b", "b", "c"}))
	fmt.Println(RemoveString([]string{"a", "b", "b", "c"}, "b"))
	fmt.Println(RemoveString([]string{"a", "b", "b", "c"}, "a", "b"))
	// Output:
	// [a b b c]
	// [a c]
	// [c]
}

func ExampleRemoveBool() {
	fmt.Println(RemoveBool([]bool{}))
	fmt.Println(RemoveBool([]bool{true, false, false, true}))
	fmt.Println(RemoveBool([]bool{true, false, false, true}, false))
	fmt.Println(RemoveBool([]bool{true, false, false, true}, true, false))
	// Output:
	// []
	// [true false false true]
	// [true true]
	// []
}

func ExampleRemoveInt() {
	fmt.Println(RemoveInt([]int{1, 2, 2, 3}))
	fmt.Println(RemoveInt([]int{1, 2, 2, 3}, 2))
	fmt.Println(RemoveInt([]int{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveInt8() {
	fmt.Println(RemoveInt8([]int8{1, 2, 2, 3}))
	fmt.Println(RemoveInt8([]int8{1, 2, 2, 3}, 2))
	fmt.Println(RemoveInt8([]int8{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveInt16() {
	fmt.Println(RemoveInt16([]int16{1, 2, 2, 3}))
	fmt.Println(RemoveInt16([]int16{1, 2, 2, 3}, 2))
	fmt.Println(RemoveInt16([]int16{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveInt32() {
	fmt.Println(RemoveInt32([]int32{1, 2, 2, 3}))
	fmt.Println(RemoveInt32([]int32{1, 2, 2, 3}, 2))
	fmt.Println(RemoveInt32([]int32{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveInt64() {
	fmt.Println(RemoveInt64([]int64{1, 2, 2, 3}))
	fmt.Println(RemoveInt64([]int64{1, 2, 2, 3}, 2))
	fmt.Println(RemoveInt64([]int64{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveUint() {
	fmt.Println(RemoveUint([]uint{1, 2, 2, 3}))
	fmt.Println(RemoveUint([]uint{1, 2, 2, 3}, 2))
	fmt.Println(RemoveUint([]uint{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveUint8() {
	fmt.Println(RemoveUint8([]uint8{1, 2, 2, 3}))
	fmt.Println(RemoveUint8([]uint8{1, 2, 2, 3}, 2))
	fmt.Println(RemoveUint8([]uint8{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveUint16() {
	fmt.Println(RemoveUint16([]uint16{1, 2, 2, 3}))
	fmt.Println(RemoveUint16([]uint16{1, 2, 2, 3}, 2))
	fmt.Println(RemoveUint16([]uint16{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveUint32() {
	fmt.Println(RemoveUint32([]uint32{1, 2, 2, 3}))
	fmt.Println(RemoveUint32([]uint32{1, 2, 2, 3}, 2))
	fmt.Println(RemoveUint32([]uint32{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
func ExampleRemoveUint64() {
	fmt.Println(RemoveUint64([]uint64{1, 2, 2, 3}))
	fmt.Println(RemoveUint64([]uint64{1, 2, 2, 3}, 2))
	fmt.Println(RemoveUint64([]uint64{1, 2, 2, 3}, 1, 2))
	// Output:
	// [1 2 2 3]
	// [1 3]
	// [3]
}
