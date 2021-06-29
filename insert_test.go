package slice

import (
	"fmt"
)

func ExampleInsertGeneric() {
	s := []string(nil)
	InsertGeneric(&s, 0, "a")
	fmt.Println(s)

	s = []string{}
	InsertGeneric(&s, 0, "a")
	fmt.Println(s)

	s = []string{"b", "c"}
	InsertGeneric(&s, 0, "a")
	fmt.Println(s)

	s = []string{"b", "c"}
	InsertGeneric(&s, 1, "a")
	fmt.Println(s)

	s = []string{"b", "c"}
	InsertGeneric(&s, 2, "a")
	fmt.Println(s)
	// Output:
	// [a]
	// [a]
	// [a b c]
	// [b a c]
	// [b c a]
}

func ExampleInsertInterface() {
	var s []interface{}
	InsertInterface(&s, 0, "a")
	fmt.Println(s)

	s = []interface{}{}
	InsertInterface(&s, 0, "a")
	fmt.Println(s)

	s = []interface{}{"b", 1}
	InsertInterface(&s, 0, "a")
	fmt.Println(s)

	s = []interface{}{"b", 1}
	InsertInterface(&s, 1, "a")
	fmt.Println(s)

	s = []interface{}{"b", 1}
	InsertInterface(&s, 2, "a")
	fmt.Println(s)
	// Output:
	// [a]
	// [a]
	// [a b 1]
	// [b a 1]
	// [b 1 a]
}

func ExampleInsertString() {
	var s []string
	InsertString(&s, 0, "a")
	fmt.Println(s)

	s = []string{}
	InsertString(&s, 0, "a")
	fmt.Println(s)

	s = []string{"b", "c"}
	InsertString(&s, 0, "a")
	fmt.Println(s)

	s = []string{"b", "c"}
	InsertString(&s, 1, "a")
	fmt.Println(s)

	s = []string{"b", "c"}
	InsertString(&s, 2, "a")
	fmt.Println(s)

	// Output:
	// [a]
	// [a]
	// [a b c]
	// [b a c]
	// [b c a]
}

func ExampleInsertInt() {
	var s []int
	InsertInt(&s, 0, 3)
	fmt.Println(s)

	s = []int{}
	InsertInt(&s, 0, 3)
	fmt.Println(s)

	s = []int{2, 3}
	InsertInt(&s, 0, 1)
	fmt.Println(s)

	s = []int{2, 3}
	InsertInt(&s, 1, 1)
	fmt.Println(s)

	s = []int{2, 3}
	InsertInt(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}

func ExampleInsertInt8() {
	var s []int8
	InsertInt8(&s, 0, 3)
	fmt.Println(s)

	s = []int8{}
	InsertInt8(&s, 0, 3)
	fmt.Println(s)

	s = []int8{2, 3}
	InsertInt8(&s, 0, 1)
	fmt.Println(s)

	s = []int8{2, 3}
	InsertInt8(&s, 1, 1)
	fmt.Println(s)

	s = []int8{2, 3}
	InsertInt8(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertInt16() {
	var s []int16
	InsertInt16(&s, 0, 3)
	fmt.Println(s)

	s = []int16{}
	InsertInt16(&s, 0, 3)
	fmt.Println(s)

	s = []int16{2, 3}
	InsertInt16(&s, 0, 1)
	fmt.Println(s)

	s = []int16{2, 3}
	InsertInt16(&s, 1, 1)
	fmt.Println(s)

	s = []int16{2, 3}
	InsertInt16(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertInt32() {
	var s []int32
	InsertInt32(&s, 0, 3)
	fmt.Println(s)

	s = []int32{}
	InsertInt32(&s, 0, 3)
	fmt.Println(s)

	s = []int32{2, 3}
	InsertInt32(&s, 0, 1)
	fmt.Println(s)

	s = []int32{2, 3}
	InsertInt32(&s, 1, 1)
	fmt.Println(s)

	s = []int32{2, 3}
	InsertInt32(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertInt64() {
	var s []int64
	InsertInt64(&s, 0, 3)
	fmt.Println(s)

	s = []int64{}
	InsertInt64(&s, 0, 3)
	fmt.Println(s)

	s = []int64{2, 3}
	InsertInt64(&s, 0, 1)
	fmt.Println(s)

	s = []int64{2, 3}
	InsertInt64(&s, 1, 1)
	fmt.Println(s)

	s = []int64{2, 3}
	InsertInt64(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}

func ExampleInsertUint() {
	var s []uint
	InsertUint(&s, 0, 3)
	fmt.Println(s)

	s = []uint{}
	InsertUint(&s, 0, 3)
	fmt.Println(s)

	s = []uint{2, 3}
	InsertUint(&s, 0, 1)
	fmt.Println(s)

	s = []uint{2, 3}
	InsertUint(&s, 1, 1)
	fmt.Println(s)

	s = []uint{2, 3}
	InsertUint(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint8() {
	var s []uint8
	InsertUint8(&s, 0, 3)
	fmt.Println(s)

	s = []uint8{}
	InsertUint8(&s, 0, 3)
	fmt.Println(s)

	s = []uint8{2, 3}
	InsertUint8(&s, 0, 1)
	fmt.Println(s)

	s = []uint8{2, 3}
	InsertUint8(&s, 1, 1)
	fmt.Println(s)

	s = []uint8{2, 3}
	InsertUint8(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint16() {
	var s []uint16
	InsertUint16(&s, 0, 3)
	fmt.Println(s)

	s = []uint16{}
	InsertUint16(&s, 0, 3)
	fmt.Println(s)

	s = []uint16{2, 3}
	InsertUint16(&s, 0, 1)
	fmt.Println(s)

	s = []uint16{2, 3}
	InsertUint16(&s, 1, 1)
	fmt.Println(s)

	s = []uint16{2, 3}
	InsertUint16(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint32() {
	var s []uint32
	InsertUint32(&s, 0, 3)
	fmt.Println(s)

	s = []uint32{}
	InsertUint32(&s, 0, 3)
	fmt.Println(s)

	s = []uint32{2, 3}
	InsertUint32(&s, 0, 1)
	fmt.Println(s)

	s = []uint32{2, 3}
	InsertUint32(&s, 1, 1)
	fmt.Println(s)

	s = []uint32{2, 3}
	InsertUint32(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
func ExampleInsertUint64() {
	var s []uint64
	InsertUint64(&s, 0, 3)
	fmt.Println(s)

	s = []uint64{}
	InsertUint64(&s, 0, 3)
	fmt.Println(s)

	s = []uint64{2, 3}
	InsertUint64(&s, 0, 1)
	fmt.Println(s)

	s = []uint64{2, 3}
	InsertUint64(&s, 1, 1)
	fmt.Println(s)

	s = []uint64{2, 3}
	InsertUint64(&s, 2, 1)
	fmt.Println(s)

	// Output:
	// [3]
	// [3]
	// [1 2 3]
	// [2 1 3]
	// [2 3 1]
}
