package slice

import (
	"testing"
)

func TestIndexString(t *testing.T) {
	var slice = []string{`1`, `2`, `3`}
	var target = `3`
	if got := IndexString(slice, target); got != 2 {
		t.Errorf("expect %d, got %d", 2, got)
	}
	if got := IndexString(slice, `4`); got != -1 {
		t.Errorf("expect -1,got %d", got)
	}
}

func TestLastIndexString(t *testing.T) {
	var slice = []string{`1`, `2`, `3`, `1`}
	var target = `1`
	if got := LastIndexString(slice, target); got != 3 {
		t.Errorf("expect %d, got %d", 3, got)
	}
	if got := LastIndexString(slice, `4`); got != -1 {
		t.Errorf("expect -1,got %d", got)
	}
}

func TestIndexInt(t *testing.T) {
	var slice = []int{1, 2, 3, 1}
	var target = 1
	if got := IndexInt(slice, target); got != 0 {
		t.Errorf("expect %d, got %d", 0, got)
	}
	if got := IndexInt(slice, 4); got != -1 {
		t.Errorf("expect -1,got %d", got)
	}
}

func TestLastIndexInt(t *testing.T) {
	var slice = []int{1, 2, 3, 1, 2}
	var target = 1
	if got := LastIndexInt(slice, target); got != 3 {
		t.Errorf("expect %d, got %d", 3, got)
	}
	if got := LastIndexInt(slice, 4); got != -1 {
		t.Errorf("expect -1,got %d", got)
	}
}

func TestIndexInt64(t *testing.T) {
	var slice = []int64{1, 2, 3, 1}
	var target = int64(1)
	if got := IndexInt64(slice, target); got != 0 {
		t.Errorf("expect %d, got %d", 0, got)
	}
	if got := IndexInt64(slice, 4); got != -1 {
		t.Errorf("expect -1,got %d", got)
	}
}

func TestLastIndexInt64(t *testing.T) {
	var slice = []int64{1, 2, 3, 1, 2}
	var target = int64(1)
	if got := LastIndexInt64(slice, target); got != 3 {
		t.Errorf("expect %d, got %d", 3, got)
	}
	if got := LastIndexInt64(slice, 4); got != -1 {
		t.Errorf("expect -1,got %d", got)
	}
}
