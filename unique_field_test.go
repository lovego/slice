package slice

import "fmt"

func ExampleUniqueField() {
	slice := []T{{1, "a"}, {2, "a"}, {1, "b"}, {3, "c"}}
	fmt.Println(UniqueField(slice, "Id"))
	fmt.Println(UniqueField(slice, "Name"))
	// Output:
	// [1 2 3]
	// [a b c]
}

func ExampleUniqueFieldString() {
	slice := []T{{1, "a"}, {2, "a"}, {1, "b"}, {3, "c"}}
	fmt.Println(UniqueFieldString(slice, "Name"))
	// Output:
	// [a b c]
}

func ExampleUniqueFieldBool() {
	type t struct{ Flag bool }
	slice := []t{{true}, {true}}
	fmt.Println(UniqueFieldBool(slice, "Flag"))
	// Output:
	// [true]
}

func ExampleUniqueFieldInt() {
	slice := []T{{1, "a"}, {2, "a"}, {1, "b"}, {3, "c"}}
	fmt.Println(UniqueFieldInt(slice, "Id"))
	// Output:
	// [1 2 3]
}
func ExampleUniqueFieldInt8() {
	slice := []T{{1, "a"}, {2, "a"}, {1, "b"}, {3, "c"}}
	fmt.Println(UniqueFieldInt8(slice, "Id"))
	// Output:
	// [1 2 3]
}
func ExampleUniqueFieldInt16() {
	slice := []T{{1, "a"}, {2, "a"}, {1, "b"}, {3, "c"}}
	fmt.Println(UniqueFieldInt16(slice, "Id"))
	// Output:
	// [1 2 3]
}
func ExampleUniqueFieldInt32() {
	slice := []T{{1, "a"}, {2, "a"}, {1, "b"}, {3, "c"}}
	fmt.Println(UniqueFieldInt32(slice, "Id"))
	// Output:
	// [1 2 3]
}
func ExampleUniqueFieldInt64() {
	slice := []T{{1, "a"}, {2, "a"}, {1, "b"}, {3, "c"}}
	fmt.Println(UniqueFieldInt64(slice, "Id"))
	// Output:
	// [1 2 3]
}

type inner struct{ Uint uint }

func ExampleUniqueFieldUint() {
	slice := []struct{ inner }{{inner{1}}, {inner{1}}, {inner{2}}, {inner{3}}, {inner{3}}}
	fmt.Println(UniqueFieldUint(slice, "inner", "Uint"))
	// Output:
	// [1 2 3]
}

func ExampleUniqueFieldUint8() {
	slice := []struct{ inner }{{inner{1}}, {inner{1}}, {inner{2}}, {inner{3}}, {inner{3}}}
	fmt.Println(UniqueFieldUint8(slice, "inner", "Uint"))
	// Output:
	// [1 2 3]
}

func ExampleUniqueFieldUint16() {
	slice := []struct{ inner }{{inner{1}}, {inner{1}}, {inner{2}}, {inner{3}}, {inner{3}}}
	fmt.Println(UniqueFieldUint16(slice, "inner", "Uint"))
	// Output:
	// [1 2 3]
}

func ExampleUniqueFieldUint32() {
	slice := []struct{ inner }{{inner{1}}, {inner{1}}, {inner{2}}, {inner{3}}, {inner{3}}}
	fmt.Println(UniqueFieldUint32(slice, "inner", "Uint"))
	// Output:
	// [1 2 3]
}
func ExampleUniqueFieldUint64() {
	slice := []struct{ inner }{{inner{1}}, {inner{1}}, {inner{2}}, {inner{3}}, {inner{3}}}
	fmt.Println(UniqueFieldUint64(slice, "inner", "Uint"))
	// Output:
	// [1 2 3]
}
