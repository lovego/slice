package slice

import (
	"fmt"
)

func ExampleUnoinString() {
	left := []string{"1", "2", "2", "3"}
	right := []string{"3", "4", "3"}
	fmt.Println(UnionString(left, right))
	// Output:
	// [1 2 3 4]
}
