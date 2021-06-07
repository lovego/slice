package slice

import (
	"fmt"
)

func ExampleSubstract() {
	fmt.Println(Substract(nil, []T{}))
	fmt.Println(Substract([]T{}, nil))
	var left = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}, {3, "c"}}
	fmt.Println(Substract(left, nil))
	fmt.Println(Substract(left, []T{{}, {2, "b"}, {2, "c"}}))

	// Output:
	// <nil>
	// []
	// [{3 c} {0 } {2 b} {9 f} {3 c}]
	// [{3 c} {9 f} {3 c}]
}

func ExampleSubstractInterface() {
	fmt.Println(SubstractInterface(nil, []interface{}{"a"}))
	fmt.Println(SubstractInterface([]interface{}{"a", "b", 3, "a"}, []interface{}{"b"}))
	fmt.Println(SubstractInterface([]interface{}{"a", "b", 3, "a"}, []interface{}{"a", 3, "d"}))
	fmt.Println(SubstractInterface([]interface{}{"a", "b", 3, "a"}, nil))
	fmt.Println(SubstractInterface([]interface{}{"a", "b", 3, "a"}, []interface{}{"d"}))
	// Output:
	// []
	// [a 3 a]
	// [b]
	// [a b 3 a]
	// [a b 3 a]
}
func ExampleSubstractString() {
	fmt.Println(SubstractString(nil, []string{"a"}))
	fmt.Println(SubstractString([]string{"a", "b", "c", "a"}, []string{"b"}))
	fmt.Println(SubstractString([]string{"a", "b", "c", "a"}, []string{"a", "c", "d"}))
	fmt.Println(SubstractString([]string{"a", "b", "c", "a"}, nil))
	fmt.Println(SubstractString([]string{"a", "b", "c", "a"}, []string{"d"}))
	// Output:
	// []
	// [a c a]
	// [b]
	// [a b c a]
	// [a b c a]
}

func ExampleSubstractInt() {
	fmt.Println(SubstractInt([]int{}, []int{2}))
	fmt.Println(SubstractInt([]int{1, 2, 3, 1}, []int{2}))
	fmt.Println(SubstractInt([]int{1, 2, 3, 1}, []int{1, 3, 4}))
	fmt.Println(SubstractInt([]int{1, 2, 3, 1}, nil))
	fmt.Println(SubstractInt([]int{1, 2, 3, 1}, []int{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractInt8() {
	fmt.Println(SubstractInt8([]int8{}, []int8{2}))
	fmt.Println(SubstractInt8([]int8{1, 2, 3, 1}, []int8{2}))
	fmt.Println(SubstractInt8([]int8{1, 2, 3, 1}, []int8{1, 3, 4}))
	fmt.Println(SubstractInt8([]int8{1, 2, 3, 1}, nil))
	fmt.Println(SubstractInt8([]int8{1, 2, 3, 1}, []int8{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractInt16() {
	fmt.Println(SubstractInt16([]int16{}, []int16{2}))
	fmt.Println(SubstractInt16([]int16{1, 2, 3, 1}, []int16{2}))
	fmt.Println(SubstractInt16([]int16{1, 2, 3, 1}, []int16{1, 3, 4}))
	fmt.Println(SubstractInt16([]int16{1, 2, 3, 1}, nil))
	fmt.Println(SubstractInt16([]int16{1, 2, 3, 1}, []int16{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractInt32() {
	fmt.Println(SubstractInt32([]int32{}, []int32{2}))
	fmt.Println(SubstractInt32([]int32{1, 2, 3, 1}, []int32{2}))
	fmt.Println(SubstractInt32([]int32{1, 2, 3, 1}, []int32{1, 3, 4}))
	fmt.Println(SubstractInt32([]int32{1, 2, 3, 1}, nil))
	fmt.Println(SubstractInt32([]int32{1, 2, 3, 1}, []int32{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractInt64() {
	fmt.Println(SubstractInt64([]int64{}, []int64{2}))
	fmt.Println(SubstractInt64([]int64{1, 2, 3, 1}, []int64{2}))
	fmt.Println(SubstractInt64([]int64{1, 2, 3, 1}, []int64{1, 3, 4}))
	fmt.Println(SubstractInt64([]int64{1, 2, 3, 1}, nil))
	fmt.Println(SubstractInt64([]int64{1, 2, 3, 1}, []int64{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}

func ExampleSubstractUint() {
	fmt.Println(SubstractUint([]uint{}, []uint{2}))
	fmt.Println(SubstractUint([]uint{1, 2, 3, 1}, []uint{2}))
	fmt.Println(SubstractUint([]uint{1, 2, 3, 1}, []uint{1, 3, 4}))
	fmt.Println(SubstractUint([]uint{1, 2, 3, 1}, nil))
	fmt.Println(SubstractUint([]uint{1, 2, 3, 1}, []uint{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractUint8() {
	fmt.Println(SubstractUint8([]uint8{}, []uint8{2}))
	fmt.Println(SubstractUint8([]uint8{1, 2, 3, 1}, []uint8{2}))
	fmt.Println(SubstractUint8([]uint8{1, 2, 3, 1}, []uint8{1, 3, 4}))
	fmt.Println(SubstractUint8([]uint8{1, 2, 3, 1}, nil))
	fmt.Println(SubstractUint8([]uint8{1, 2, 3, 1}, []uint8{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractUint16() {
	fmt.Println(SubstractUint16([]uint16{}, []uint16{2}))
	fmt.Println(SubstractUint16([]uint16{1, 2, 3, 1}, []uint16{2}))
	fmt.Println(SubstractUint16([]uint16{1, 2, 3, 1}, []uint16{1, 3, 4}))
	fmt.Println(SubstractUint16([]uint16{1, 2, 3, 1}, nil))
	fmt.Println(SubstractUint16([]uint16{1, 2, 3, 1}, []uint16{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractUint32() {
	fmt.Println(SubstractUint32([]uint32{}, []uint32{2}))
	fmt.Println(SubstractUint32([]uint32{1, 2, 3, 1}, []uint32{2}))
	fmt.Println(SubstractUint32([]uint32{1, 2, 3, 1}, []uint32{1, 3, 4}))
	fmt.Println(SubstractUint32([]uint32{1, 2, 3, 1}, nil))
	fmt.Println(SubstractUint32([]uint32{1, 2, 3, 1}, []uint32{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
func ExampleSubstractUint64() {
	fmt.Println(SubstractUint64([]uint64{}, []uint64{2}))
	fmt.Println(SubstractUint64([]uint64{1, 2, 3, 1}, []uint64{2}))
	fmt.Println(SubstractUint64([]uint64{1, 2, 3, 1}, []uint64{1, 3, 4}))
	fmt.Println(SubstractUint64([]uint64{1, 2, 3, 1}, nil))
	fmt.Println(SubstractUint64([]uint64{1, 2, 3, 1}, []uint64{4}))
	// Output:
	// []
	// [1 3 1]
	// [2]
	// [1 2 3 1]
	// [1 2 3 1]
}
