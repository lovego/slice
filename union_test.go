package slice

import (
	"fmt"
	"reflect"
)

func ExampleUnionGeneric() {
	fmt.Println(UnionGeneric(nil, nil))

	var left = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}, {2, "b"}}
	fmt.Println(UnionGeneric(left, nil))
	fmt.Println(UnionGeneric(nil, []T{{3, "c"}, {3, "c"}}))
	fmt.Println(UnionGeneric(left, []int64{}))
	fmt.Println(UnionGeneric(left, []T{{}, {2, "b"}, {4, "d"}}))
	// Output:
	// <nil>
	// [{3 c} {0 } {2 b} {9 f} {2 b}]
	// [{3 c}]
	// [{3 c} {0 } {2 b} {9 f} {2 b}]
	// [{3 c} {0 } {2 b} {9 f} {2 b} {4 d}]
}

func ExampleUnionValue() {
	fmt.Println(UnionValue(reflect.ValueOf(nil), reflect.ValueOf(nil)).IsValid())
	// Output:
	// false
}

func ExampleUnionInterface() {
	fmt.Println(UnionInterface([]interface{}{1, 2, 2, "3"}, []interface{}{"3", "4", "3", 4}))
	// Output:
	// [1 2 2 3 4 4]
}

func ExampleUnionString() {
	fmt.Println(UnionString([]string{"1", "2", "2", "3"}, []string{"3", "4", "3", "4"}))
	// Output:
	// [1 2 2 3 4]
}

func ExampleUnionInt() {
	fmt.Println(UnionInt([]int{1, 2, 2, 3}, []int{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionInt8() {
	fmt.Println(UnionInt8([]int8{1, 2, 2, 3}, []int8{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionInt16() {
	fmt.Println(UnionInt16([]int16{1, 2, 2, 3}, []int16{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionInt32() {
	fmt.Println(UnionInt32([]int32{1, 2, 2, 3}, []int32{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionInt64() {
	fmt.Println(UnionInt64([]int64{1, 2, 2, 3}, []int64{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}

func ExampleUnionUint() {
	fmt.Println(UnionUint([]uint{1, 2, 2, 3}, []uint{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionUint8() {
	fmt.Println(UnionUint8([]uint8{1, 2, 2, 3}, []uint8{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionUint16() {
	fmt.Println(UnionUint16([]uint16{1, 2, 2, 3}, []uint16{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionUint32() {
	fmt.Println(UnionUint32([]uint32{1, 2, 2, 3}, []uint32{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
func ExampleUnionUint64() {
	fmt.Println(UnionUint64([]uint64{1, 2, 2, 3}, []uint64{3, 4, 3, 4}))
	// Output:
	// [1 2 2 3 4]
}
