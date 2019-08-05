package slice

import "fmt"

func ExampleDifferenceInt64() {
	fmt.Println(DifferenceInt64([]int64{1, 2}, []int64{2}))
	fmt.Println(DifferenceInt64([]int64{1}, []int64{2}))
	fmt.Println(DifferenceInt64([]int64{2}, []int64{1}))
	fmt.Println(DifferenceInt64(nil, []int64{1}))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}

func ExampleDifferenceString() {
	fmt.Println(DifferenceString([]string{"1", "2"}, []string{"2"}))
	fmt.Println(DifferenceString([]string{"1"}, []string{"2"}))
	fmt.Println(DifferenceString([]string{"2"}, []string{"1"}))
	fmt.Println(DifferenceString(nil, []string{"1"}))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}
