package slice

import "fmt"

func ExampleUniqueInt() {
	s := []int{1, 1, 2, 2, 3, 3, 4}
	res := UniqueInt(s)
	fmt.Println(res)
	// Output:
	// [1 2 3 4]
}

func ExampleUniqueInt64() {
	s := []int64{1, 1, 2, 2, 3, 3, 4}
	res := UniqueInt64(s)
	fmt.Println(res)
	// Output:
	// [1 2 3 4]
}

func ExampleUniqueString() {
	s := []string{`a`, `a`, `b`, `b`, `c`}
	res := UniqueString(s)
	fmt.Println(res)
	// Output:
	// [a b c]
}
