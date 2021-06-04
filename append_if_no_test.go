package slice

import (
	"fmt"
	"reflect"
)

func ExampleAppendIfNo() {
	var slice = []T{{3, "c"}, {}, {2, "b"}, {}, {9, "f"}}
	fmt.Println(AppendIfNo(slice, T{2, "b"}))
	fmt.Println(AppendIfNo(slice, T{}))
	fmt.Println(AppendIfNo(slice, T{}, T{4, "d"}))

	// Output:
	// [{3 c} {0 } {2 b} {0 } {9 f}]
	// [{3 c} {0 } {2 b} {0 } {9 f}]
	// [{3 c} {0 } {2 b} {0 } {9 f} {4 d}]
}

func ExampleAppendIfNoValue() {
	var slice = reflect.ValueOf([]T{{3, "c"}, {}, {2, "b"}, {}, {9, "f"}})
	fmt.Println(AppendIfNoValue(slice, reflect.ValueOf(T{2, "b"})))
	fmt.Println(AppendIfNoValue(slice, reflect.ValueOf(T{})))
	fmt.Println(AppendIfNoValue(slice, reflect.ValueOf(T{}), reflect.ValueOf(T{4, "d"})))

	// Output:
	// [{3 c} {0 } {2 b} {0 } {9 f}]
	// [{3 c} {0 } {2 b} {0 } {9 f}]
	// [{3 c} {0 } {2 b} {0 } {9 f} {4 d}]
}

func ExampleAppendIfNoInterface() {
	var slice = []interface{}{3, "c", false, 2, "b"}
	fmt.Println(AppendIfNoInterface(slice, 3, "b"))
	fmt.Println(AppendIfNoInterface(slice, false, 2))
	fmt.Println(AppendIfNoInterface(slice, "c", 4, "d"))

	// Output:
	// [3 c false 2 b]
	// [3 c false 2 b]
	// [3 c false 2 b 4 d]
}

func ExampleAppendIfNoString() {
	fmt.Println(AppendIfNoString(nil, "a"))
	fmt.Println(AppendIfNoString([]string{"a", "b"}, "a"))
	fmt.Println(AppendIfNoString([]string{"a", "b"}, "b", "c"))
	// Output:
	// [a]
	// [a b]
	// [a b c]
}

func ExampleAppendIfNoInt() {
	fmt.Println(AppendIfNoInt(nil, 1))
	fmt.Println(AppendIfNoInt([]int{1, 2}, 1))
	fmt.Println(AppendIfNoInt([]int{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoInt8() {
	fmt.Println(AppendIfNoInt8(nil, 1))
	fmt.Println(AppendIfNoInt8([]int8{1, 2}, 1))
	fmt.Println(AppendIfNoInt8([]int8{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoInt16() {
	fmt.Println(AppendIfNoInt16(nil, 1))
	fmt.Println(AppendIfNoInt16([]int16{1, 2}, 1))
	fmt.Println(AppendIfNoInt16([]int16{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoInt32() {
	fmt.Println(AppendIfNoInt32(nil, 1))
	fmt.Println(AppendIfNoInt32([]int32{1, 2}, 1))
	fmt.Println(AppendIfNoInt32([]int32{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoInt64() {
	fmt.Println(AppendIfNoInt64(nil, 1))
	fmt.Println(AppendIfNoInt64([]int64{1, 2}, 1))
	fmt.Println(AppendIfNoInt64([]int64{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoUint() {
	fmt.Println(AppendIfNoUint(nil, 1))
	fmt.Println(AppendIfNoUint([]uint{1, 2}, 1))
	fmt.Println(AppendIfNoUint([]uint{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoUint8() {
	fmt.Println(AppendIfNoUint8(nil, 1))
	fmt.Println(AppendIfNoUint8([]uint8{1, 2}, 1))
	fmt.Println(AppendIfNoUint8([]uint8{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoUint16() {
	fmt.Println(AppendIfNoUint16(nil, 1))
	fmt.Println(AppendIfNoUint16([]uint16{1, 2}, 1))
	fmt.Println(AppendIfNoUint16([]uint16{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoUint32() {
	fmt.Println(AppendIfNoUint32(nil, 1))
	fmt.Println(AppendIfNoUint32([]uint32{1, 2}, 1))
	fmt.Println(AppendIfNoUint32([]uint32{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoUint64() {
	fmt.Println(AppendIfNoUint64(nil, 1))
	fmt.Println(AppendIfNoUint64([]uint64{1, 2}, 1))
	fmt.Println(AppendIfNoUint64([]uint64{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}
