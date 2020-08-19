package slice

import (
	"fmt"
	"reflect"
)

func ExampleIndexString() {
	slice := []string{`1`, `2`, `3`}
	fmt.Println(IndexString(slice, `3`))
	fmt.Println(IndexString(slice, `4`))
	// Output:
	// 2
	// -1
}

func ExampleLastIndexString() {
	slice := []string{`1`, `2`, `3`, `1`}
	fmt.Println(LastIndexString(slice, `1`))
	fmt.Println(LastIndexString(slice, `4`))
	// Output:
	// 3
	// -1
}

func ExampleIndexInt() {
	slice := []int{1, 2, 3, 1}
	fmt.Println(IndexInt(slice, 1))
	fmt.Println(IndexInt(slice, 4))
	// Output:
	// 0
	// -1
}

func ExampleLastIndexInt() {
	slice := []int{1, 2, 3, 1, 2}
	fmt.Println(LastIndexInt(slice, 1))
	fmt.Println(LastIndexInt(slice, 4))
	// Output:
	// 3
	// -1
}

func ExampleIndexInt64() {
	var slice = []int64{1, 2, 3, 1}
	fmt.Println(IndexInt64(slice, 1))
	fmt.Println(IndexInt64(slice, 4))
	// Output:
	// 0
	// -1
}

func ExampleLastIndexInt64() {
	var slice = []int64{1, 2, 3, 1, 2}
	fmt.Println(LastIndexInt64(slice, 1))
	fmt.Println(LastIndexInt64(slice, 4))
	// Output:
	// 3
	// -1
}

type T struct {
	Id   int
	Name string
}

func ExampleIndex() {
	var slice = []T{{3, "c"}, {}, {2, "b"}, {9, "f"}}
	fmt.Println(Index(slice, T{2, "b"}))
	fmt.Println(Index(slice, T{2, "c"}))
	// Output:
	// 2
	// -1
}

func ExampleLastIndex() {
	var slice = []T{{3, "c"}, {}, {2, "b"}, {}, {9, "f"}}
	fmt.Println(LastIndex(slice, T{}))
	fmt.Println(LastIndex(slice, T{2, "c"}))
	// Output:
	// 3
	// -1
}

func ExampleIndexValue() {
	var slice = reflect.ValueOf([]T{{3, "c"}, {}, {2, "b"}, {9, "f"}})
	fmt.Println(IndexValue(slice, reflect.ValueOf(T{2, "b"})))
	fmt.Println(IndexValue(slice, reflect.ValueOf(T{2, "c"})))
	// Output:
	// 2
	// -1
}

func ExampleLastIndexValue() {
	var slice = reflect.ValueOf([]T{{3, "c"}, {}, {2, "b"}, {}, {9, "f"}})
	fmt.Println(LastIndexValue(slice, reflect.ValueOf(T{})))
	fmt.Println(LastIndexValue(slice, reflect.ValueOf(T{2, "c"})))
	// Output:
	// 3
	// -1
}
