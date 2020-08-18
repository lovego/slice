package slice

import "fmt"

func ExampleMoveBackward() {
	var s []int
	fmt.Println(MoveBackward(s, 0, 3))
	fmt.Println(MoveBackward([]int{0, 4, 5}, 0, 3))
	fmt.Println(MoveBackward([]int{0, 4, 5}, 1, 3))
	fmt.Println(MoveBackward([]int{0, 4, 5}, 2, 3))
	fmt.Println(MoveBackward([]int{0, 4, 5}, 3, 3))
	// Output:
	// [0 0 0]
	// [0 4 5 0 4 5]
	// [0 4 5 0 4 5]
	// [0 4 5 0 0 5]
	// [0 4 5 0 0 0]
}
