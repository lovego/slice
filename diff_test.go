package slice

import "fmt"

func ExampleDiffInt64() {
	fmt.Println(DiffInt64([]int64{1, 2}, []int64{2}))
	fmt.Println(DiffInt64([]int64{1}, []int64{2}))
	fmt.Println(DiffInt64([]int64{2}, []int64{1}))
	fmt.Println(DiffInt64(nil, []int64{1}))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}

func ExampleDiffString() {
	fmt.Println(DiffString([]string{"1", "2"}, []string{"2"}))
	fmt.Println(DiffString([]string{"1"}, []string{"2"}))
	fmt.Println(DiffString([]string{"2"}, []string{"1"}))
	fmt.Println(DiffString(nil, []string{"1"}))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}
