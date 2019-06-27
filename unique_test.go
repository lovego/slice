package slice

import (
	"encoding/json"
	"fmt"
	"log"
)

type fakeData struct {
	A int
	B string
	C int64
}
type keys struct {
	A int
	B string
}

func ExampleUnique() {
	s := []*fakeData{
		{1, `1`, 1},
		{1, `1`, 2},
		{2, `2`, 3},
		{1, `1`, 4},
		{2, `2`, 5},
	}

	Unique(&s, func(i int) interface{} {
		return keys{s[i].A, s[i].B}
	}, false)

	data, err := json.Marshal(s)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(data))

	// Output:
	// [{"A":1,"B":"1","C":1},{"A":2,"B":"2","C":3}]
}

func ExampleUnique2() {
	s := []*fakeData{
		{1, `1`, 1},
		{1, `1`, 2},
		{2, `2`, 3},
		{2, `2`, 4},
	}

	Unique(&s, func(i int) interface{} {
		return keys{s[i].A, s[i].B}
	}, true)

	data, err := json.Marshal(s)
	if err != nil {
		log.Panic(err)
	}
	fmt.Println(string(data))

	// Output:
	// [{"A":1,"B":"1","C":2},{"A":2,"B":"2","C":4}]
}

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
