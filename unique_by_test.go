package slice

import (
	"fmt"
)

func ExampleUniqueBy() {
	slice := []T{}
	UniqueBy(&slice, func(i int) interface{} { return slice[i].Id }, false)
	fmt.Println(slice)

	slice = []T{{1, `a`}, {1, `b`}, {2, `a`}, {2, `b`}}
	UniqueBy(&slice, func(i int) interface{} { return slice[i].Id }, false)
	fmt.Println(slice)
	// Output:
	// []
	// [{1 a} {2 a}]
}

func ExampleUniqueBy_keepLast() {
	slice := []T{}
	UniqueBy(&slice, func(i int) interface{} { return slice[i].Id }, true)
	fmt.Println(slice)

	slice = []T{{1, `a`}, {1, `b`}, {2, `a`}, {2, `b`}}
	UniqueBy(&slice, func(i int) interface{} { return slice[i].Id }, true)
	fmt.Println(slice)
	// Output:
	// []
	// [{1 b} {2 b}]
}
