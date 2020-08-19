package slice

import (
	"fmt"
)

func ExampleInersectInt64() {
	var left = []int64{1, 2, 3, 4}
	var right = []int64{2, 1}

	fmt.Println(IntersectInt64(left, right))
	fmt.Println(IntersectInt64(left, []int64{0, 5}))

	// Output:
	// [1 2]
	// []
}

func ExampleHasIntersectInt64() {
	fmt.Println(HasIntersectInt64([]int64{1, 2}, []int64{2}))
	fmt.Println(HasIntersectInt64([]int64{1}, []int64{2}))
	fmt.Println(HasIntersectInt64([]int64{1}, nil))
	// Output:
	// true
	// false
	// false
}
