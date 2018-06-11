package slice

import (
	"reflect"
	"testing"
)

// func TestRemoveStrings(t *testing.T) {
//  var slice = []string{`1`, `2`, `3`, `4`}
//  var toRemove = []string{`2`, `4`}

//  if got {

//  }
// }

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
