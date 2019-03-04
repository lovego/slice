package slice

import "fmt"

func ExampleAppendIfNoString() {
	fmt.Println(AppendIfNoString(nil, "a"))
	fmt.Println(AppendIfNoString([]string{"a", "b"}, "a"))
	fmt.Println(AppendIfNoString([]string{"a", "b"}, "c"))
	// Output:
	// [a]
	// [a b]
	// [a b c]
}

func ExampleAppendIfNoInt() {
	fmt.Println(AppendIfNoInt(nil, 1))
	fmt.Println(AppendIfNoInt([]int{1, 2}, 1))
	fmt.Println(AppendIfNoInt([]int{1, 2}, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

func ExampleAppendIfNoInt64() {
	fmt.Println(AppendIfNoInt64(nil, 1))
	fmt.Println(AppendIfNoInt64([]int64{1, 2}, 1))
	fmt.Println(AppendIfNoInt64([]int64{1, 2}, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}
