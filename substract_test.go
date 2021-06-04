package slice

import (
	"fmt"
)

func ExampleSubstractString() {
	fmt.Println(SubstractString(nil, []string{"a"}))
	fmt.Println(SubstractString([]string{"a", "b", "c"}, []string{"b"}))
	fmt.Println(SubstractString([]string{"a", "b", "c"}, []string{"a", "c", "d"}))
	fmt.Println(SubstractString([]string{"a", "b", "c"}, nil))
	fmt.Println(SubstractString([]string{"a", "b", "c"}, []string{"d"}))
	// Output:
	// []
	// [a c]
	// [b]
	// [a b c]
	// [a b c]
}

func ExampleSubstractInt64() {
	fmt.Println(SubstractInt64([]int64{}, []int64{2}))
	fmt.Println(SubstractInt64([]int64{1, 2, 3}, []int64{2}))
	fmt.Println(SubstractInt64([]int64{1, 2, 3}, []int64{1, 3, 4}))
	fmt.Println(SubstractInt64([]int64{1, 2, 3}, nil))
	fmt.Println(SubstractInt64([]int64{1, 2, 3}, []int64{4}))
	// Output:
	// []
	// [1 3]
	// [2]
	// [1 2 3]
	// [1 2 3]
}
