package slice

import (
	"reflect"
	"testing"
)

func BenchmarkIndexGeneric(b *testing.B) {
	slice := []int{1, 2, 3, 1}
	for i := 0; i < b.N; i++ {
		IndexGeneric(slice, int(3))
	}
}

func BenchmarkIndexInt(b *testing.B) {
	slice := []int{1, 2, 3, 1}
	for i := 0; i < b.N; i++ {
		IndexInt(slice, 3)
	}
}

func BenchmarkValueOf_reflect(b *testing.B) {
	slice := []int{1, 2, 3, 1}
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(slice)
	}
}

func BenchmarkIndex_slice(b *testing.B) {
	slice := []int{1, 2, 3, 1}
	for i := 0; i < b.N; i++ {
		_ = slice[3]
	}
}
func BenchmarkIndex_reflect(b *testing.B) {
	slice := reflect.ValueOf([]int{1, 2, 3, 1})
	for i := 0; i < b.N; i++ {
		_ = slice.Index(3)
	}
}

func BenchmarkInterface_reflect(b *testing.B) {
	v := reflect.ValueOf(1)
	for i := 0; i < b.N; i++ {
		v.Interface()
	}
}

func BenchmarkEqual_int(b *testing.B) {
	var x, y int = 1, 2
	for i := 0; i < b.N; i++ {
		if x == y {
		}
	}
}
func BenchmarkEqual_interface(b *testing.B) {
	var x, y interface{} = 1, 2
	for i := 0; i < b.N; i++ {
		if x == y {
		}
	}
}
