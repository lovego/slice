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

func ExampleUniqueBy() {
	s := []*fakeData{
		{1, `1`, 1},
		{1, `1`, 2},
		{2, `2`, 3},
		{2, `2`, 4},
	}

	UniqueBy(&s, func(i int) interface{} {
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

func ExampleUniqueBy2() {
	s := []*fakeData{
		{1, `1`, 1},
		{1, `1`, 2},
		{2, `2`, 3},
		{2, `2`, 4},
	}

	UniqueBy(&s, func(i int) interface{} {
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
