package slice

import (
	"fmt"
)

func ExampleUniqueGeneric() {
	fmt.Println(UniqueGeneric([]T{{1, `a`}, {2, `b`}, {}, {3, `c`}, {2, `b`}, {3, `c`}}))
	// Output:
	// [{1 a} {2 b} {0 } {3 c}]
}

func ExampleUniqueInterface() {
	fmt.Println(UniqueInterface([]interface{}{0, 1, 1, `2`, `2`, true, true, 4}))
	// Output:
	// [0 1 2 true 4]
}

func ExampleUniqueString() {
	fmt.Println(UniqueString([]string{`0`, `a`, `a`, `b`, `b`, `c`, `c`, `d`}))
	// Output:
	// [0 a b c d]
}

func ExampleUniqueInt() {
	fmt.Println(UniqueInt([]int{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt8() {
	fmt.Println(UniqueInt8([]int8{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt16() {
	fmt.Println(UniqueInt16([]int16{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt32() {
	fmt.Println(UniqueInt32([]int32{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt64() {
	fmt.Println(UniqueInt64([]int64{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}

func ExampleUniqueUint() {
	fmt.Println(UniqueUint([]uint{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint8() {
	fmt.Println(UniqueUint8([]uint8{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint16() {
	fmt.Println(UniqueUint16([]uint16{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint32() {
	fmt.Println(UniqueUint32([]uint32{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint64() {
	fmt.Println(UniqueUint64([]uint64{0, 1, 1, 2, 2, 3, 3, 4}))
	// Output:
	// [0 1 2 3 4]
}
