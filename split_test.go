package slice

import "fmt"

func ExampleSplitString() {
	fmt.Println(SplitString([]string{}, "--"))
	fmt.Println(SplitString([]string{"--"}, "--"))
	fmt.Println(SplitString([]string{"a"}, "--"))
	fmt.Println(SplitString([]string{"a", "b"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--", "c"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--", "c", "d"}, "--"))
	fmt.Println(SplitString([]string{"a", "b", "--", "c", "d", "--"}, "--"))

	// Output:
	// []
	// [[] []]
	// [[a]]
	// [[a b]]
	// [[a b] []]
	// [[a b] [c]]
	// [[a b] [c d]]
	// [[a b] [c d] []]
}
