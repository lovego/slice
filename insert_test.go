package slice

import "fmt"

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
