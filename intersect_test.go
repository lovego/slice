package slice

import (
	"fmt"
)

func ExampleIntersect() {
	var left = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	fmt.Println(Intersect(left, nil))
	fmt.Printf("%#v\n", Intersect(left, []int64{}))
	fmt.Println(Intersect(left, []T{{}, {2, "b"}}))
	// Output:
	// <nil>
	// []slice.T(nil)
	// [{0 } {2 b}]
}

func ExampleIntersectInterface() {
	var left = []interface{}{1, `2`, 3, `4`}
	fmt.Println(IntersectInterface(left, []interface{}{`2`, 1, `3`}))
	fmt.Println(IntersectInterface(left, []interface{}{0, 5}))
	// Output:
	// [1 2]
	// []
}

func ExampleIntersectString() {
	var left = []string{"a", "b", "c", "d"}
	fmt.Println(IntersectString(left, []string{"b", "a", "e"}))
	fmt.Println(IntersectString(left, []string{"x", "e"}))
	// Output:
	// [a b]
	// []
}

func ExampleIntersectInt() {
	var left = []int{1, 2, 3, 4}
	fmt.Println(IntersectInt(left, []int{2, 1, 5}))
	fmt.Println(IntersectInt(left, []int{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectInt8() {
	var left = []int8{1, 2, 3, 4}
	fmt.Println(IntersectInt8(left, []int8{2, 1, 5}))
	fmt.Println(IntersectInt8(left, []int8{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectInt16() {
	var left = []int16{1, 2, 3, 4}
	fmt.Println(IntersectInt16(left, []int16{2, 1, 5}))
	fmt.Println(IntersectInt16(left, []int16{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectInt32() {
	var left = []int32{1, 2, 3, 4}
	fmt.Println(IntersectInt32(left, []int32{2, 1, 5}))
	fmt.Println(IntersectInt32(left, []int32{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectInt64() {
	var left = []int64{1, 2, 3, 4}
	fmt.Println(IntersectInt64(left, []int64{2, 1, 5}))
	fmt.Println(IntersectInt64(left, []int64{0, 5}))
	// Output:
	// [1 2]
	// []
}

func ExampleIntersectUint() {
	var left = []uint{1, 2, 3, 4}
	fmt.Println(IntersectUint(left, []uint{2, 1, 5}))
	fmt.Println(IntersectUint(left, []uint{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectUint8() {
	var left = []uint8{1, 2, 3, 4}
	fmt.Println(IntersectUint8(left, []uint8{2, 1, 5}))
	fmt.Println(IntersectUint8(left, []uint8{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectUint16() {
	var left = []uint16{1, 2, 3, 4}
	fmt.Println(IntersectUint16(left, []uint16{2, 1, 5}))
	fmt.Println(IntersectUint16(left, []uint16{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectUint32() {
	var left = []uint32{1, 2, 3, 4}
	fmt.Println(IntersectUint32(left, []uint32{2, 1, 5}))
	fmt.Println(IntersectUint32(left, []uint32{0, 5}))
	// Output:
	// [1 2]
	// []
}
func ExampleIntersectUint64() {
	var left = []uint64{1, 2, 3, 4}
	fmt.Println(IntersectUint64(left, []uint64{2, 1, 5}))
	fmt.Println(IntersectUint64(left, []uint64{0, 5}))
	// Output:
	// [1 2]
	// []
}
