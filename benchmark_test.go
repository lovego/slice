package slice

import (
	"math/rand"
	"reflect"
	"testing"
)

var benchSlice []int

func init() {
	var size = 300
	benchSlice = make([]int, size)
	for i := 0; i < size; i++ {
		benchSlice[i] = rand.Int() % 20
	}
}

func BenchmarkIndexGeneric(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexGeneric(benchSlice, int(3))
	}
}

// go test -v -run ^$ -bench BenchmarkIndex_by -benchtime 400x
// if benchSlice length and bench times:
// < about 300, Index_bySlice is faster.
// > about 300, Index_byMap   is faster.
func BenchmarkIndex_bySlice(b *testing.B) {
	for i := 0; i < b.N; i++ {
		IndexInt(benchSlice, len(benchSlice)-1)
	}
}

func BenchmarkIndex_byMap(b *testing.B) {
	m := make(map[int]int, len(benchSlice))
	for i, v := range benchSlice {
		m[v] = i
	}
	for i := 0; i < b.N; i++ {
		_, _ = m[len(benchSlice)-1]
	}
}

func BenchmarkValueOf_reflect(b *testing.B) {
	for i := 0; i < b.N; i++ {
		reflect.ValueOf(benchSlice)
	}
}

func BenchmarkIndex_direct(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = benchSlice[3]
	}
}
func BenchmarkIndex_reflect(b *testing.B) {
	slice := reflect.ValueOf(benchSlice)
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
