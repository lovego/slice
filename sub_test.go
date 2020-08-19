package slice

import "fmt"

func ExampleSubInt64() {
	fmt.Println(SubInt64([]int64{1, 2}, []int64{2}))
	fmt.Println(SubInt64([]int64{1}, []int64{2}))
	fmt.Println(SubInt64([]int64{2}, []int64{1}))
	fmt.Println(SubInt64(nil, []int64{1}))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}

func ExampleSubString() {
	fmt.Println(SubString([]string{"1", "2"}, []string{"2"}))
	fmt.Println(SubString([]string{"1"}, []string{"2"}))
	fmt.Println(SubString([]string{"2"}, []string{"1"}))
	fmt.Println(SubString(nil, []string{"1"}))
	// Output:
	// [1]
	// [1]
	// [2]
	// []
}
