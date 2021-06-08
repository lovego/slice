package slice

import (
	"fmt"
)

func ExampleUniqueGeneric() {
	slice := []T{{1, `a`}, {2, `b`}, {}, {3, `c`}, {2, `b`}, {3, `c`}}
	UniqueGeneric(&slice)
	fmt.Println(slice)
	// Output:
	// [{1 a} {2 b} {0 } {3 c}]
}

func ExampleUniqueInterface() {
	slice := []interface{}{0, 1, 1, `2`, `2`, true, true, 4}
	UniqueInterface(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 true 4]
}

func ExampleUniqueString() {
	slice := []string{`0`, `a`, `a`, `b`, `b`, `c`, `c`, `d`}
	UniqueString(&slice)
	fmt.Println(slice)
	// Output:
	// [0 a b c d]
}

func ExampleUniqueInt() {
	slice := []int{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueInt(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt8() {
	slice := []int8{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueInt8(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt16() {
	slice := []int16{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueInt16(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt32() {
	slice := []int32{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueInt32(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueInt64() {
	slice := []int64{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueInt64(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}

func ExampleUniqueUint() {
	slice := []uint{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueUint(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint8() {
	slice := []uint8{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueUint8(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint16() {
	slice := []uint16{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueUint16(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint32() {
	slice := []uint32{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueUint32(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
func ExampleUniqueUint64() {
	slice := []uint64{0, 1, 1, 2, 2, 3, 3, 4}
	UniqueUint64(&slice)
	fmt.Println(slice)
	// Output:
	// [0 1 2 3 4]
}
