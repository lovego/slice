package slice

import (
	"fmt"
	"reflect"
)

func ExampleSplitGeneric() {
	fmt.Println(SplitGeneric([]T{}, T{}))
	fmt.Println(SplitGeneric([]T{T{}}, T{}))
	fmt.Println(SplitGeneric([]T{T{1, "a"}}, T{}))
	fmt.Println(SplitGeneric([]T{T{1, "a"}, T{2, "b"}}, T{}))
	fmt.Println(SplitGeneric([]T{T{1, "a"}, T{2, "b"}, T{}}, T{}))
	fmt.Println(SplitGeneric([]T{T{1, "a"}, T{2, "b"}, T{}, T{3, "c"}}, T{}))
	fmt.Println(SplitGeneric([]T{T{1, "a"}, T{2, "b"}, T{}, T{3, "c"}, T{4, "d"}}, T{}))
	fmt.Println(SplitGeneric([]T{T{1, "a"}, T{2, "b"}, T{}, T{3, "c"}, T{4, "d"}, T{}}, T{}))

	// Output:
	// []
	// [[] []]
	// [[{1 a}]]
	// [[{1 a} {2 b}]]
	// [[{1 a} {2 b}] []]
	// [[{1 a} {2 b}] [{3 c}]]
	// [[{1 a} {2 b}] [{3 c} {4 d}]]
	// [[{1 a} {2 b}] [{3 c} {4 d}] []]
}

func ExampleSplitValue() {
	fmt.Println(SplitValue(reflect.ValueOf([]T{}), T{}))
	fmt.Println(SplitValue(reflect.ValueOf([]T{T{}}), T{}))
	fmt.Println(SplitValue(reflect.ValueOf([]T{T{1, "a"}}), T{}))
	fmt.Println(SplitValue(reflect.ValueOf([]T{T{1, "a"}, T{2, "b"}}), T{}))
	fmt.Println(SplitValue(reflect.ValueOf([]T{T{1, "a"}, T{2, "b"}, T{}}), T{}))
	fmt.Println(SplitValue(reflect.ValueOf([]T{T{1, "a"}, T{2, "b"}, T{}, T{3, "c"}}), T{}))
	fmt.Println(SplitValue(reflect.ValueOf([]T{T{1, "a"}, T{2, "b"}, T{}, T{3, "c"}, T{4, "d"}}), T{}))
	fmt.Println(SplitValue(reflect.ValueOf([]T{T{1, "a"}, T{2, "b"}, T{}, T{3, "c"}, T{4, "d"}, T{}}), T{}))

	// Output:
	// []
	// [[] []]
	// [[{1 a}]]
	// [[{1 a} {2 b}]]
	// [[{1 a} {2 b}] []]
	// [[{1 a} {2 b}] [{3 c}]]
	// [[{1 a} {2 b}] [{3 c} {4 d}]]
	// [[{1 a} {2 b}] [{3 c} {4 d}] []]

}

func ExampleSplitInterface() {
	fmt.Println(SplitInterface([]interface{}{}, "--"))
	fmt.Println(SplitInterface([]interface{}{"--"}, "--"))
	fmt.Println(SplitInterface([]interface{}{"a"}, "--"))
	fmt.Println(SplitInterface([]interface{}{"a", 1}, "--"))
	fmt.Println(SplitInterface([]interface{}{"a", 1, "--"}, "--"))
	fmt.Println(SplitInterface([]interface{}{"a", 1, "--", "c"}, "--"))
	fmt.Println(SplitInterface([]interface{}{"a", 1, "--", "c", 3}, "--"))
	fmt.Println(SplitInterface([]interface{}{"a", 1, "--", "c", 3, "--"}, "--"))

	// Output:
	// []
	// [[] []]
	// [[a]]
	// [[a 1]]
	// [[a 1] []]
	// [[a 1] [c]]
	// [[a 1] [c 3]]
	// [[a 1] [c 3] []]
}

func ExampleSplitString() {
	fmt.Println(SplitString([]string{}, "--"))
	fmt.Println(SplitString([]string{"--"}, "--"))
	fmt.Println(SplitString([]string{"a"}, "--"))
	fmt.Println(SplitString([]string{"a", "b"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--", "c"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--", "c", "d"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--", "c", "d", "--"}, "--"))

	// Output:
	// []
	// [[] []]
	// [[a]]
	// [[a b]]
	// [[a b] []]
	// [[a b] [c]]
	// [[a b] [c d]]
	// [[a b] [c d] []]
}

func ExampleSplitBool() {
	fmt.Println(SplitBool([]bool{}, false))
	fmt.Println(SplitBool([]bool{false}, false))
	fmt.Println(SplitBool([]bool{true}, false))
	fmt.Println(SplitBool([]bool{true, true}, false))
	fmt.Println(SplitBool([]bool{true, true, false}, false))
	fmt.Println(SplitBool([]bool{true, true, false, true}, false))
	fmt.Println(SplitBool([]bool{true, true, false, true, true}, false))
	fmt.Println(SplitBool([]bool{true, true, false, true, true, false}, false))

	// Output:
	// []
	// [[] []]
	// [[true]]
	// [[true true]]
	// [[true true] []]
	// [[true true] [true]]
	// [[true true] [true true]]
	// [[true true] [true true] []]
}

func ExampleSplitInt() {
	fmt.Println(SplitInt([]int{}, 0))
	fmt.Println(SplitInt([]int{0}, 0))
	fmt.Println(SplitInt([]int{1}, 0))
	fmt.Println(SplitInt([]int{1, 2}, 0))
	fmt.Println(SplitInt([]int{1, 2, 0}, 0))
	fmt.Println(SplitInt([]int{1, 2, 0, 3}, 0))
	fmt.Println(SplitInt([]int{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitInt([]int{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitInt8() {
	fmt.Println(SplitInt8([]int8{}, 0))
	fmt.Println(SplitInt8([]int8{0}, 0))
	fmt.Println(SplitInt8([]int8{1}, 0))
	fmt.Println(SplitInt8([]int8{1, 2}, 0))
	fmt.Println(SplitInt8([]int8{1, 2, 0}, 0))
	fmt.Println(SplitInt8([]int8{1, 2, 0, 3}, 0))
	fmt.Println(SplitInt8([]int8{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitInt8([]int8{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitInt16() {
	fmt.Println(SplitInt16([]int16{}, 0))
	fmt.Println(SplitInt16([]int16{0}, 0))
	fmt.Println(SplitInt16([]int16{1}, 0))
	fmt.Println(SplitInt16([]int16{1, 2}, 0))
	fmt.Println(SplitInt16([]int16{1, 2, 0}, 0))
	fmt.Println(SplitInt16([]int16{1, 2, 0, 3}, 0))
	fmt.Println(SplitInt16([]int16{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitInt16([]int16{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitInt32() {
	fmt.Println(SplitInt32([]int32{}, 0))
	fmt.Println(SplitInt32([]int32{0}, 0))
	fmt.Println(SplitInt32([]int32{1}, 0))
	fmt.Println(SplitInt32([]int32{1, 2}, 0))
	fmt.Println(SplitInt32([]int32{1, 2, 0}, 0))
	fmt.Println(SplitInt32([]int32{1, 2, 0, 3}, 0))
	fmt.Println(SplitInt32([]int32{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitInt32([]int32{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitInt64() {
	fmt.Println(SplitInt64([]int64{}, 0))
	fmt.Println(SplitInt64([]int64{0}, 0))
	fmt.Println(SplitInt64([]int64{1}, 0))
	fmt.Println(SplitInt64([]int64{1, 2}, 0))
	fmt.Println(SplitInt64([]int64{1, 2, 0}, 0))
	fmt.Println(SplitInt64([]int64{1, 2, 0, 3}, 0))
	fmt.Println(SplitInt64([]int64{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitInt64([]int64{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitUint() {
	fmt.Println(SplitUint([]uint{}, 0))
	fmt.Println(SplitUint([]uint{0}, 0))
	fmt.Println(SplitUint([]uint{1}, 0))
	fmt.Println(SplitUint([]uint{1, 2}, 0))
	fmt.Println(SplitUint([]uint{1, 2, 0}, 0))
	fmt.Println(SplitUint([]uint{1, 2, 0, 3}, 0))
	fmt.Println(SplitUint([]uint{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitUint([]uint{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitUint8() {
	fmt.Println(SplitUint8([]uint8{}, 0))
	fmt.Println(SplitUint8([]uint8{0}, 0))
	fmt.Println(SplitUint8([]uint8{1}, 0))
	fmt.Println(SplitUint8([]uint8{1, 2}, 0))
	fmt.Println(SplitUint8([]uint8{1, 2, 0}, 0))
	fmt.Println(SplitUint8([]uint8{1, 2, 0, 3}, 0))
	fmt.Println(SplitUint8([]uint8{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitUint8([]uint8{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitUint16() {
	fmt.Println(SplitUint16([]uint16{}, 0))
	fmt.Println(SplitUint16([]uint16{0}, 0))
	fmt.Println(SplitUint16([]uint16{1}, 0))
	fmt.Println(SplitUint16([]uint16{1, 2}, 0))
	fmt.Println(SplitUint16([]uint16{1, 2, 0}, 0))
	fmt.Println(SplitUint16([]uint16{1, 2, 0, 3}, 0))
	fmt.Println(SplitUint16([]uint16{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitUint16([]uint16{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitUint32() {
	fmt.Println(SplitUint32([]uint32{}, 0))
	fmt.Println(SplitUint32([]uint32{0}, 0))
	fmt.Println(SplitUint32([]uint32{1}, 0))
	fmt.Println(SplitUint32([]uint32{1, 2}, 0))
	fmt.Println(SplitUint32([]uint32{1, 2, 0}, 0))
	fmt.Println(SplitUint32([]uint32{1, 2, 0, 3}, 0))
	fmt.Println(SplitUint32([]uint32{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitUint32([]uint32{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}

func ExampleSplitUint64() {
	fmt.Println(SplitUint64([]uint64{}, 0))
	fmt.Println(SplitUint64([]uint64{0}, 0))
	fmt.Println(SplitUint64([]uint64{1}, 0))
	fmt.Println(SplitUint64([]uint64{1, 2}, 0))
	fmt.Println(SplitUint64([]uint64{1, 2, 0}, 0))
	fmt.Println(SplitUint64([]uint64{1, 2, 0, 3}, 0))
	fmt.Println(SplitUint64([]uint64{1, 2, 0, 3, 4}, 0))
	fmt.Println(SplitUint64([]uint64{1, 2, 0, 3, 4, 0}, 0))

	// Output:
	// []
	// [[] []]
	// [[1]]
	// [[1 2]]
	// [[1 2] []]
	// [[1 2] [3]]
	// [[1 2] [3 4]]
	// [[1 2] [3 4] []]
}
