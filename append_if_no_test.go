package slice

import "fmt"

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

func ExampleAppendIfNoInt64() {
	fmt.Println(AppendIfNoInt64(nil, 1))
	fmt.Println(AppendIfNoInt64([]int64{1, 2}, 1))
	fmt.Println(AppendIfNoInt64([]int64{1, 2}, 2, 3))
	// Output:
	// [1]
	// [1 2]
	// [1 2 3]
}

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
