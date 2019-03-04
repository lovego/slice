package slice

import (
	"fmt"
	"reflect"
	"testing"
)

func TestInersectInt64(t *testing.T) {
	var left = []int64{1, 2, 3, 4}
	var right = []int64{2, 1}
	var expect = []int64{2, 1}
	if got := IntersectInt64(left, right); !reflect.DeepEqual(got, expect) {
		t.Errorf("expect %v, got %v", expect, got)
	}
	left = []int64{}
	right = []int64{}
	if got := IntersectInt64(left, right); !reflect.DeepEqual([]int64{}, got) {
		t.Errorf("expect %v, got %v", []int64{}, got)
	}
}

func ExampleHasIntersectInt64() {
	fmt.Println(HasIntersectInt64([]int64{1, 2}, []int64{2}))
	fmt.Println(HasIntersectInt64([]int64{1}, []int64{2}))
	fmt.Println(HasIntersectInt64([]int64{1}, nil))
	// Output:
	// true
	// false
	// false
}
