package slice

import (
	"fmt"
	"reflect"
)

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
