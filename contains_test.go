package slice

import (
	"fmt"
)

func ExampleContainsString() {
	var slice = []string{`1`, `2`, `3`}
	fmt.Println(ContainsString(slice, `2`))
	fmt.Println(ContainsString(slice, `4`))
	// Output:
	// true
	// false
}

func ExampleContainsInt() {
	var slice = []int{1, 2, 3}
	fmt.Println(ContainsInt(slice, 2))
	fmt.Println(ContainsInt(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsInt64() {
	var slice = []int64{1, 2, 3}
	fmt.Println(ContainsInt64(slice, 2))
	fmt.Println(ContainsInt64(slice, 4))
	// Output:
	// true
	// false
}

func ExampleContainsAll() {

	fmt.Println(ContainsAllInt64([]int64{1, 2}, []int64{2, 1}))
	fmt.Println(ContainsAllString([]string{"1", "2"}, []string{"2", "1"}))

	fmt.Println(ContainsAllInt64([]int64{1, 2}, nil))
	fmt.Println(ContainsAllInt64([]int64{1, 2}, []int64{1}))
	fmt.Println(ContainsAllInt64([]int64{1, 2}, []int64{2, 3}))
	fmt.Println(ContainsAllInt64([]int64{1, 2}, []int64{3, 4}))

	fmt.Println(ContainsAllString([]string{"1", "2"}, nil))
	fmt.Println(ContainsAllString([]string{"1", "2"}, []string{"1"}))
	fmt.Println(ContainsAllString([]string{"1", "2"}, []string{"2", "3"}))
	fmt.Println(ContainsAllString([]string{"1", "2"}, []string{"3", "4"}))

	// Output:
	// true
	// true
	// false
	// false
	// false
	// false
	// false
	// false
	// false
	// false
}
