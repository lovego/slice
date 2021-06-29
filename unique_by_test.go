package slice

import (
	"fmt"
)

func ExampleUniqueBy() {
	slice := []T{}
	fmt.Println(UniqueBy(slice, func(i int) interface{} { return slice[i].Id }, false))

	slice = []T{{1, `a`}, {1, `b`}, {2, `a`}, {2, `b`}}
	fmt.Println(UniqueBy(slice, func(i int) interface{} { return slice[i].Id }, false))
	// Output:
	// []
	// [{1 a} {2 a}]
}

func ExampleUniqueBy_keepLast() {
	slice := []T{}
	fmt.Println(UniqueBy(slice, func(i int) interface{} { return slice[i].Id }, true))

	slice = []T{{1, `a`}, {1, `b`}, {2, `a`}, {2, `b`}}
	fmt.Println(UniqueBy(slice, func(i int) interface{} { return slice[i].Id }, true))
	// Output:
	// []
	// [{1 b} {2 b}]
}
