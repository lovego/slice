package slice

import (
	"reflect"
	"testing"
)

func TestRemoveStrings(t *testing.T) {
	var slice = []string{`1`, `2`, `3`, `4`}
	var toRemove = []string{`2`, `4`}
	var expect = []string{`1`, `3`}
	if got := RemoveStrings(slice, toRemove); !reflect.DeepEqual(got, expect) {
		t.Errorf("expect %v, got %v", expect, got)
	}
	slice = []string{}
	toRemove = []string{}
	if got := RemoveStrings(slice, toRemove); !reflect.DeepEqual(got, []string{}) {
		t.Errorf("expect %v, got %v", []int{}, got)
	}
}

func TestRemoveInts(t *testing.T) {
	var slice = []int{1, 2, 3, 10}
	var toRemove = []int{2, 10}
	var expect = []int{1, 3}
	if got := RemoveInts(slice, toRemove); !reflect.DeepEqual(got, expect) {
		t.Errorf("expect %v,got %v", expect, got)
	}
	slice = []int{}
	toRemove = []int{}
	if got := RemoveInts(slice, toRemove); !reflect.DeepEqual(got, []int{}) {
		t.Errorf("expect %v, got %v", []int{}, got)
	}
}

func TestRemoveInt64s(t *testing.T) {
	var slice = []int64{1, 2, 3, 10}
	var toRemove = []int64{2, 10}
	var expect = []int64{1, 3}
	if got := RemoveInt64s(slice, toRemove); !reflect.DeepEqual(got, expect) {
		t.Errorf("expect %v,got %v", expect, got)
	}
	slice = []int64{}
	toRemove = []int64{}
	if got := RemoveInt64s(slice, toRemove); !reflect.DeepEqual(got, []int64{}) {
		t.Errorf("expect %v, got %v", []int64{}, got)
	}
}

func TestRemove(t *testing.T) {
	var slice = []interface{}{`1`, 2, `3`, `4`}
	var toRemove = []interface{}{2, `4`}
	if got := Remove(slice, toRemove); len(got) != 2 ||
		(got[0] != `1` && got[0] != `3`) {
		t.Errorf("unexpect got %v,unexpect type %b", got, reflect.TypeOf(got))
	}
	slice = []interface{}{}
	toRemove = []interface{}{}
	if got := Remove(slice, toRemove); !reflect.DeepEqual(got, slice) {
		t.Errorf("unexpect got %v", got)
	}
}

func TestRemoveStruct(t *testing.T) {
	type st struct {
		Id   int
		Name string
	}
	var a = []st{
		{
			Id:   1,
			Name: `小明`,
		}, {
			Id:   2,
			Name: `小红`,
		},
	}
	var slice = []interface{}{}
	for _, v := range a {
		slice = append(slice, v)
	}
	slice = append(slice, `woo~~`)
	var toRemove = []interface{}{st{Id: 1, Name: `小明`}}
	var expect = []interface{}{st{Id: 2, Name: `小红`}, `woo~~`}
	if got := Remove(slice, toRemove); !reflect.DeepEqual(got, expect) {
		t.Errorf("expect %v, got %v", expect, got)
	}
}
