package slice

import (
	"testing"
)

func TestContainsString(t *testing.T) {
	var slice = []string{`1`, `2`, `3`}
	var target = `2`
	if got := ContainsString(slice, target); !got {
		t.Errorf("unexpect got %t", got)
	}
	if got := ContainsString(slice, `4`); got {
		t.Errorf("unexpect got %t", got)
	}
}

func TestContainsInt(t *testing.T) {
	var slice = []int{1, 2, 3}
	var target = 2
	if got := ContainsInt(slice, target); !got {
		t.Errorf("unexpect got %t", got)
	}
	if got := ContainsInt(slice, 4); got {
		t.Errorf("unexpect got %t", got)
	}
}

func TestContainsInt64(t *testing.T) {
	var slice = []int64{1, 2, 3}
	var target = int64(2)
	if got := ContainsInt64(slice, target); !got {
		t.Errorf("unexpect got %t", got)
	}
	if got := ContainsInt64(slice, 4); got {
		t.Errorf("unexpect got %t", got)
	}
}
