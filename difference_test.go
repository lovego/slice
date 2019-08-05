package slice

import "fmt"

func ExampleDiffrenceInt64() {
	fmt.Println(differenceInt64([]int64{1, 2}, []int64{2}))
	fmt.Println(differenceInt64([]int64{1}, []int64{2}))
	fmt.Println(differenceInt64([]int64{2}, []int64{1}))
	fmt.Println(differenceInt64([]int64{1}, nil))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}

func ExampleDiffrenceString() {
	fmt.Println(differenceString([]string{"1", "2"}, []string{"2"}))
	fmt.Println(differenceString([]string{"1"}, []string{"2"}))
	fmt.Println(differenceString([]string{"2"}, []string{"1"}))
	fmt.Println(differenceString([]string{"1"}, nil))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}
