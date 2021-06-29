package slice

import "fmt"

func ExampleMoveBackward() {
	s := make([]int, 0, 4)
	MoveBackward(&s, 0, 3)
	fmt.Println(s)

	s = []int{0, 4, 5}
	MoveBackward(&s, 0, 3)
	fmt.Println(s)

	s = []int{0, 4, 5}
	MoveBackward(&s, 1, 3)
	fmt.Println(s)

	s = []int{0, 4, 5}
	MoveBackward(&s, 2, 3)
	fmt.Println(s)

	s = []int{0, 4, 5}
	MoveBackward(&s, 3, 3)
	fmt.Println(s)

	// Output:
	// [0 0 0]
	// [0 4 5 0 4 5]
	// [0 4 5 0 4 5]
	// [0 4 5 0 0 5]
	// [0 4 5 0 0 0]
}
